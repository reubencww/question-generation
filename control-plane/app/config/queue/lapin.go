package queue

import (
	"github.com/wagslane/go-rabbitmq"
	"google.golang.org/protobuf/proto"

	"senkawa.moe/haa-chan/hskw"
)

const (
	captionQueue = "caption"
	qnaQueue     = "qna"
)

type Queue interface {
	DispatchCaption(challenge int, path string) error
	DispatchQNA(challenge int, caption string) error
	Close() error
}

type Publisher struct {
	queue *rabbitmq.Publisher
}

func createRabbitMQPublisher(url string) (*rabbitmq.Publisher, error) {
	publisher, err := rabbitmq.NewPublisher(url, rabbitmq.Config{})
	if err != nil {
		return nil, err
	}

	return publisher, nil
}

func New(url string) (Queue, error) {
	publisher, err := createRabbitMQPublisher(url)
	if err != nil {
		return nil, err
	}

	return &Publisher{
		queue: publisher,
	}, nil
}

func (r *Publisher) DispatchCaption(challenge int, path string) error {
	body, err := proto.Marshal(&hskw.GenerateCaption{
		ChallengeId: int32(challenge),
		ImagePath:   path,
	})
	if err != nil {
		return err
	}

	if err = r.queue.Publish(body, []string{captionQueue}); err != nil {
		return err
	}

	return nil
}

func (r *Publisher) DispatchQNA(challenge int, caption string) error {
	body, err := proto.Marshal(&hskw.GenerateQuestionAnswer{
		ChallengeId: int32(challenge),
		Caption:     caption,
	})
	if err != nil {
		return err
	}

	if err = r.queue.Publish(body, []string{qnaQueue}); err != nil {
		return err
	}

	return nil
}

func (r *Publisher) Close() error {
	return r.queue.Close()
}
