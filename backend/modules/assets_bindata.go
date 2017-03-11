// +build bindata

package modules

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/view"
)

var (
	// serve from bindata
	Public = func(app *iris.Framework) {
		// set the template engine
		app.Adapt(view.HTML("./templates", ".html").Layout("layout.html").Binary(Asset, AssetNames))

		// bundle file
		app.StaticEmbedded("/bundle", "./bundle", Asset, AssetNames)

		// set static folder(s)
		app.StaticEmbedded("/public", "./public", Asset, AssetNames)
	}
)
