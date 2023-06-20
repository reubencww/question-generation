package web

import (
	"html/template"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	recoveryMiddleware "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/template/html"
	vueglue "github.com/torenware/vite-go"

	"senkawa.moe/haa-chan/app/config"
	"senkawa.moe/haa-chan/frontend"
	"senkawa.moe/haa-chan/templates"
)

type ViewManager struct {
	VueGlue *vueglue.VueGlue
	Engine  *html.Engine
}

const csrfKey = "csrf"

func NewViewEngine() *ViewManager {
	var viteConfig *vueglue.ViteConfig

	if config.IsDevelopment() {
		viteConfig = &vueglue.ViteConfig{
			Environment: "development",
			AssetsPath:  "frontend",
			EntryPoint:  "src/main.ts",
			Platform:    "vue",
			FS:          os.DirFS("frontend"),
		}
	} else {
		viteConfig = &vueglue.ViteConfig{
			Environment: "production",
			AssetsPath:  "dist",
			EntryPoint:  "src/main.ts",
			Platform:    "vue",
			FS:          frontend.DistAssets,
		}
	}

	glue, err := vueglue.NewVueGlue(viteConfig)
	if err != nil {
		panic(err)
	}

	engine := html.NewFileSystem(http.FS(templates.Templates), ".gohtml")
	engine.AddFunc("glue", func() template.HTML {
		tags, err := glue.RenderTags()
		if err != nil {
			panic(err)
		}

		return tags
	})

	return &ViewManager{
		VueGlue: glue,
		Engine:  engine,
	}
}

func Build() *fiber.App {
	viewEngine := NewViewEngine()
	app := fiber.New(fiber.Config{
		Views:             viewEngine.Engine,
		EnablePrintRoutes: config.IsDevelopment(),
		BodyLimit:         20 * 1024 * 1024, // 20MiB file upload limit
		ServerHeader:      "Haa",
		AppName:           "はあ",
	})
	app.Use(recoveryMiddleware.New())
	app.Use(requestid.New())

	return app
}
