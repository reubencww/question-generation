package web

import (
	"context"
	"encoding/json"
	"log"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"senkawa.moe/haa-chan/app/config/queue"
	"senkawa.moe/haa-chan/app/db"
	"senkawa.moe/haa-chan/app/db/repositories"
	"senkawa.moe/haa-chan/app/storage"
	"senkawa.moe/haa-chan/app/web/nichika/ws"
	pb "senkawa.moe/haa-chan/hskw"
)

type server struct {
	pb.UnimplementedSnappyServer

	DB      *gorm.DB
	Log     *zap.SugaredLogger
	Hub     *ws.Hub
	Storage storage.Storage
	Queue   queue.Queue
}

func (s *server) CreatedCaption(ctx context.Context, request *pb.CreatedCaptionRequest) (*pb.CreatedCaptionResponse, error) {
	challengeRepository := repositories.NewChallengeRepository(ctx, s.DB)
	challenge, err := challengeRepository.GetChallenge(uint(request.GetChallengeId()))
	if err != nil {
		return nil, err
	}

	challenge.Caption = request.GetCaption()
	if err := challengeRepository.UpdateChallenge(challenge); err != nil {
		s.Log.Errorw("failed to update challenge", "error", err)
		return nil, status.Error(codes.Internal, "failed to update challenge")
	}

	if err := s.Queue.DispatchQNA(int(request.GetChallengeId()), challenge.Caption); err != nil {
		s.Log.Errorw("failed to dispatch qna", "error", err)
		return nil, status.Error(codes.Internal, "failed to dispatch qna")
	}

	return &pb.CreatedCaptionResponse{}, nil
}

func (s *server) CreatedQuestion(ctx context.Context, request *pb.CreatedQuestionRequest) (*pb.CreatedQuestionResponse, error) {
	challengeRepository := repositories.NewChallengeRepository(ctx, s.DB)
	challenge, err := challengeRepository.GetChallenge(uint(request.GetChallengeId()))
	if err != nil {
		return nil, err
	}

	// prealloc
	questions := make([]db.Question, 0, len(request.GetQnas()))
	for _, generated := range request.GetQnas() {
		questions = append(questions, db.Question{
			Question: generated.GetQuestion(),
			Answer:   generated.GetAnswer(),
		})
	}

	// update challenge
	challenge.Questions = questions
	if err := challengeRepository.CompleteChallenge(challenge); err != nil {
		s.Log.Errorw("failed to update challenge", "error", err)
		return nil, status.Error(codes.Internal, "failed to update challenge")
	}

	body, err := json.Marshal(challenge)
	if err != nil {
		return nil, err
	}

	s.Hub.Broadcast <- body

	return &pb.CreatedQuestionResponse{}, nil
}

func buildFromWebApplication(application *Application) *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterSnappyServer(s, &server{
		DB:      application.DB,
		Log:     application.Log,
		Hub:     application.Hub,
		Storage: application.Storage,
		Queue:   application.Queue,
	})

	return s
}

func RunGRPCServer(application *Application, listenAddr string) error {
	bindAddr, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}

	grpcServer := buildFromWebApplication(application)

	go func() {
		if err := grpcServer.Serve(bindAddr); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	return nil
}
