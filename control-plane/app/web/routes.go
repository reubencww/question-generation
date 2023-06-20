package web

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/websocket/v2"

	"senkawa.moe/haa-chan/app/web/nichika/ws"
	"senkawa.moe/haa-chan/frontend"
)

const outBoundBufferSize = 256

func (a *Application) RegisterRoutes() {
	app := a.Server
	a.RegisterWebsocketRoutes(app.Group("/ws"))

	app.Use("/assets", filesystem.New(filesystem.Config{
		Root:       http.FS(frontend.DistAssets),
		PathPrefix: "dist/assets",
	}))

	a.RegisterApiRoutes(app.Group("/api/v1"))

	app.Use(func(c *fiber.Ctx) error {
		if strings.HasPrefix(c.OriginalURL(), "/api") {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "route not found"})
		}

		return c.Render("main", fiber.Map{
			"csrfToken": c.Locals(csrfKey),
		})
	})
}

func (a *Application) RegisterApiRoutes(api fiber.Router) {
	api.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	challenge := api.Group("/challenge")
	{
		challengeController := NewChallengeController(a.Log.Named("challenge"), a.DB, a.Storage, a.Queue, a.Hub)
		challenge.Get("/", challengeController.Index)
		challenge.Post("/", challengeController.Store)
		challenge.Get("/:id<int>", challengeController.Show)
		challenge.Patch("/:id<int>", challengeController.Update)

		challenge.Patch("/question/:id<int>", challengeController.UpdateQuestion)
		challenge.Delete("/question/:id<int>", challengeController.DeleteQuestion)
		challenge.Post("/new-questions/:id<int>", challengeController.NewQuestionsFromCorpus)
	}
}

func (a *Application) RegisterWebsocketRoutes(app fiber.Router) {
	hub := ws.NewHub(a.Log.Named("ws"))
	go hub.Run()
	a.Hub = hub

	app.Get("", ws.Upgrader, websocket.New(func(c *websocket.Conn) {
		client := &ws.Client{Hub: hub, Conn: c.Conn, Send: make(chan []byte, outBoundBufferSize), Log: a.Log.Named("ws-client")}
		client.Hub.Register <- client

		go client.WritePump()
		client.ReadPump()
	}))
}
