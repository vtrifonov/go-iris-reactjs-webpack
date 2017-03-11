package routes

import (
	"gopkg.in/kataras/iris.v6"
)

type index struct {
	Title   string
	Message string
}

func Index(ctx *iris.Context) {
	// MustRender, same as Render but sends status 500 internal server error if rendering failed
	ctx.MustRender("index.html", nil)
}

func (i *index) Serve(ctx *iris.Context) {
	if err := ctx.Render("index.html", i); err != nil {
		// ctx.EmitError(iris.StatusInternalServerError) =>
		ctx.Panic()
	}
}
