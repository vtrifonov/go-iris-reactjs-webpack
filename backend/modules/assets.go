// +build !bindata

package modules

import (
	"github.com/vtrifonov/go-iris-reactjs-webpack/backend/routes"
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/view"
)

var (
	Public = func(app *iris.Framework) {
		// set the template engine
		app.Adapt(view.HTML("../frontend/templates", ".html").Layout("layout.html"))

		//app.Get("/-/:rand(.*).hot-update.:ext(.*)", iris.ToHandler(routes.ReloadProxy))
		// serve bundle file from the nodejs server looking for changes
		app.Get("/bundle/*path", iris.ToHandler(routes.ReloadProxy))

		// set static folder(s)
		app.StaticWeb("/public", "../frontend/public")
	}
)
