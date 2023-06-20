package db

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	PageSize = 15
	queryKey = "page"
)

func Paginate(r *fiber.Ctx) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r.Query(queryKey, "1")
		page, _ := strconv.Atoi(q)
		if page == 0 {
			page = 1
		}

		offset := (page - 1) * PageSize

		return db.Offset(offset).Limit(PageSize)
	}
}
