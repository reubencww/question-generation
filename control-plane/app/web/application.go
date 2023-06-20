package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"senkawa.moe/haa-chan/app/config"
	"senkawa.moe/haa-chan/app/config/queue"
	"senkawa.moe/haa-chan/app/db"
	"senkawa.moe/haa-chan/app/log"
	"senkawa.moe/haa-chan/app/storage"
	"senkawa.moe/haa-chan/app/web/nichika/ws"
)

type Application struct {
	Server  *fiber.App
	DB      *gorm.DB
	Log     *zap.SugaredLogger
	Hub     *ws.Hub
	Storage storage.Storage
	Queue   queue.Queue

	Session *session.Store

	MiddlewareCsrf fiber.Handler
}

func newCsrfMiddleware() fiber.Handler {
	if config.IsDevelopment() {
		return csrf.New(csrf.Config{
			Next: func(c *fiber.Ctx) bool {
				return true
			},
			KeyLookup:  "header:X-CSRF-Token",
			ContextKey: csrfKey,
		})
	}

	return csrf.New(csrf.Config{
		KeyLookup:  "header:X-CSRF-Token",
		ContextKey: csrfKey,
	})
}

func NewApplication() *Application {
	logger := log.NewZapLogger(config.IsDevelopment()).Sugar()
	application := &Application{
		Server:         Build(),
		DB:             db.NewDB(),
		Log:            logger.Named("web"),
		Session:        session.New(),
		MiddlewareCsrf: newCsrfMiddleware(),
	}

	s3, err := storage.NewS3(storage.S3Options{
		Endpoint:  viper.GetString("storage.endpoint"),
		AccessKey: viper.GetString("storage.access_key"),
		SecretKey: viper.GetString("storage.secret_key"),
		Bucket:    viper.GetString("storage.bucket"),
		UseSSL:    viper.GetBool("storage.use_ssl"),
	})
	if err != nil {
		panic(err)
	}
	application.Storage = s3

	publisher, err := queue.New(viper.GetString("queue.url"))
	if err != nil {
		panic(err)
	}
	application.Queue = publisher

	application.RegisterRoutes()

	return application
}
