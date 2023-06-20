package repositories

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"

	. "senkawa.moe/haa-chan/app/db"
)

type ChallengeRepository interface {
	CreateChallenge(challenge *Challenge) (*Challenge, error)
	GetChallenge(id uint) (*Challenge, error)
	UpdateChallenge(challenge *Challenge) error
	CompleteChallenge(challenge *Challenge) error
	GetAllChallenge() ([]Challenge, error)
	GetAllPaginatedChallenge(ctx *fiber.Ctx) (challenges []Challenge, total int64, err error)

	GetQuestion(id uint) (*Question, error)
	DeleteQuestion(question *Question) error
	UpdateQuestion(question *Question) error
}

type challengeRepository struct {
	db *gorm.DB
}

func (r *challengeRepository) GetQuestion(id uint) (*Question, error) {
	var question Question
	tx := r.db.First(&question, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &question, nil
}

func (r *challengeRepository) UpdateQuestion(question *Question) error {
	tx := r.db.Save(question)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *challengeRepository) DeleteQuestion(question *Question) error {
	tx := r.db.Delete(question)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func NewChallengeRepository(ctx context.Context, db *gorm.DB) ChallengeRepository {
	return &challengeRepository{db: db.WithContext(ctx)}
}

func NewChallengeRepositoryWithC(c *fiber.Ctx, db *gorm.DB) ChallengeRepository {
	return NewChallengeRepository(c.Context(), db)
}

// GetAllChallenge returns all challenges
func (r *challengeRepository) GetAllChallenge() ([]Challenge, error) {
	var challenges []Challenge
	tx := r.db.Preload("Questions").Order("created_at desc").Find(&challenges)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return challenges, nil
}

func (r *challengeRepository) GetAllPaginatedChallenge(c *fiber.Ctx) (challenges []Challenge, total int64, err error) {
	tx := r.db.Preload("Questions").Scopes(Paginate(c)).Order("created_at desc").Find(&challenges)
	if tx.Error != nil {
		return nil, total, tx.Error
	}

	tx = r.db.Model(&Challenge{}).Count(&total)

	return challenges, total, nil
}

// CreateChallenge inserts a new Challenge
func (r *challengeRepository) CreateChallenge(challenge *Challenge) (*Challenge, error) {
	tx := r.db.Create(challenge)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return challenge, nil
}

// GetChallenge returns a Challenge by ID
func (r *challengeRepository) GetChallenge(id uint) (*Challenge, error) {
	var challenge Challenge
	tx := r.db.Preload("Questions").First(&challenge, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &challenge, nil
}

// UpdateChallenge updates a challenge with the generated questions
func (r *challengeRepository) UpdateChallenge(challenge *Challenge) error {
	tx := r.db.Save(challenge)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// CompleteChallenge updates a challenge with the generated questions and sets the status to complete
func (r *challengeRepository) CompleteChallenge(challenge *Challenge) error {
	challenge.CompletedAt = null.TimeFrom(time.Now())
	tx := r.db.Save(challenge)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
