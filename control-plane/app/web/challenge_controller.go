package web

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/h2non/filetype"
	"github.com/h2non/filetype/types"
	"github.com/rs/xid"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"senkawa.moe/haa-chan/app/config/queue"
	"senkawa.moe/haa-chan/app/db"
	"senkawa.moe/haa-chan/app/db/repositories"
	"senkawa.moe/haa-chan/app/storage"
	"senkawa.moe/haa-chan/app/web/nichika/ws"
)

type ChallengeController struct {
	log     *zap.SugaredLogger
	db      *gorm.DB
	storage storage.Storage
	queue   queue.Queue
	hub     *ws.Hub
}

func NewChallengeController(log *zap.SugaredLogger, db *gorm.DB, storage storage.Storage, queue queue.Queue, hub *ws.Hub) *ChallengeController {
	return &ChallengeController{
		log:     log,
		db:      db,
		storage: storage,
		queue:   queue,
		hub:     hub,
	}
}

type ChallengeUpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (cc *ChallengeController) Index(c *fiber.Ctx) error {
	challengeRepository := repositories.NewChallengeRepositoryWithC(c, cc.db)
	challenge, total, err := challengeRepository.GetAllPaginatedChallenge(c)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "error fetching all challenges",
		})
	}

	return c.JSON(fiber.Map{
		"data": challenge,
		"meta": fiber.Map{
			"total": total,
			"max":   db.PageSize,
		},
	})
}

func (cc *ChallengeController) Store(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil && len(form.File["image"]) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "no file uploaded",
		})
	}

	buf, err := form.File["image"][0].Open()
	if err != nil {
		cc.log.Errorw("failed to open file", "error", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "failed to open file for reading",
		})
	}
	defer buf.Close()

	kind, isImage := isAcceptedImageType(buf)
	if !isImage {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "file is not an accepted image type (png, jpg only)",
		})
	}

	filename := fmt.Sprintf("uploads/%s.%s", xid.New().String(), kind.Extension)
	if err = cc.storage.Upload(c.Context(), filename, kind.MIME.Value, buf); err != nil {
		cc.log.Errorw("failed to upload file", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	challengeRepository := repositories.NewChallengeRepositoryWithC(c, cc.db)
	created, err := challengeRepository.CreateChallenge(&db.Challenge{
		Name:        c.FormValue("name"),
		Description: c.FormValue("description"),
		Filename:    filename,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := cc.queue.DispatchCaption(int(created.ID), created.Filename); err != nil {
		cc.log.Errorw("failed to dispatch caption job", "error", err, "challenge_id", created.ID, "filename", created.Filename)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": created,
	})
}

func (cc *ChallengeController) Show(c *fiber.Ctx) error {
	id, err := parseId(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	challengeRepository := repositories.NewChallengeRepositoryWithC(c, cc.db)
	challenge, err := challengeRepository.GetChallenge(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "challenge not found",
		})
	}

	challenge.Filename, err = cc.storage.GetURL(c.Context(), challenge.Filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": challenge,
	})
}

func (cc *ChallengeController) Update(c *fiber.Ctx) error {
	id, err := parseId(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	challengeRepository := repositories.NewChallengeRepositoryWithC(c, cc.db)
	challenge, err := challengeRepository.GetChallenge(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "challenge not found",
		})
	}

	var req ChallengeUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	// we're not updating questions here
	challenge.Name = req.Name
	challenge.Description = req.Description
	if err := challengeRepository.UpdateChallenge(challenge); err != nil {
		cc.log.Errorw("failed to update challenge", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	challenge, err = challengeRepository.GetChallenge(id)
	if err != nil {
		cc.log.Errorw("failed to get challenge after update", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	body, err := json.Marshal(challenge)
	if err != nil {
		cc.log.Errorw("failed to marshal challenge", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	cc.hub.Broadcast <- body

	return c.JSON(fiber.Map{
		"data": challenge,
	})
}

func (cc *ChallengeController) DeleteQuestion(c *fiber.Ctx) error {
	id, err := parseId(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	challengeRepository := repositories.NewChallengeRepositoryWithC(c, cc.db)
	question, err := challengeRepository.GetQuestion(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "challenge not found",
		})
	}

	if err := challengeRepository.DeleteQuestion(question); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

type ChallengeQuestionUpdateRequest struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func (cc *ChallengeController) UpdateQuestion(c *fiber.Ctx) error {
	id, err := parseId(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	challengeRepository := repositories.NewChallengeRepositoryWithC(c, cc.db)
	question, err := challengeRepository.GetQuestion(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "challenge not found",
		})
	}

	var req ChallengeQuestionUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	question.Question = req.Question
	question.Answer = req.Answer
	if err := challengeRepository.UpdateQuestion(question); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

type ChallengeNewQuestionsRequest struct {
	Corpus string `json:"corpus"`
}

func (cc *ChallengeController) NewQuestionsFromCorpus(c *fiber.Ctx) error {
	id, err := parseId(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	challengeRepository := repositories.NewChallengeRepositoryWithC(c, cc.db)
	challenge, err := challengeRepository.GetChallenge(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "challenge not found",
		})
	}

	var req ChallengeNewQuestionsRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := cc.queue.DispatchQNA(int(challenge.ID), req.Corpus); err != nil {
		cc.log.Errorw("failed to dispatch qna job", "error", err, "challenge_id", challenge.ID, "corpus", req.Corpus)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func parseId(c *fiber.Ctx) (uint, error) {
	id := c.Params("id")
	parsed, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(parsed), nil
}

func isAcceptedImageType(file multipart.File) (types.Type, bool) {
	head := make([]byte, 261)
	file.Read(head)
	defer file.Seek(0, io.SeekStart)

	for _, t := range []string{"png", "jpg"} {
		kind, _ := filetype.Match(head)
		if kind.Extension == t {
			return kind, true
		}
	}

	return filetype.Unknown, false
}
