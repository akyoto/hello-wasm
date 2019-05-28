package main

import (
	"github.com/aerogo/aero"
)

func main() {
	app := aero.New()
	configure(app).Run()
}

func configure(app *aero.Application) *aero.Application {
	app.Get("/", func(ctx *aero.Context) string {
		return ctx.File("frontend/index.html")
	})

	app.Get("/wasm_exec.js", func(ctx *aero.Context) string {
		return ctx.File("frontend/wasm_exec.js")
	})

	app.Get("/main.wasm", func(ctx *aero.Context) string {
		return ctx.File("frontend/main.wasm")
	})

	return app
}
