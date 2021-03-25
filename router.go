package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/basicauth"
)

func setRouter(app *iris.Application) {
	view := iris.HTML("./views", ".html")
	view.AddFunc("formatBytes", func(b int64) string {
		const unit = 1000
		if b < unit {
			return fmt.Sprintf("%d B", b)
		}
		div, exp := int64(unit), 0
		for n := b / unit; n >= unit; n /= unit {
			div *= unit
			exp++
		}
		return fmt.Sprintf("%.1f %cB",
			float64(b)/float64(div), "kMGTPE"[exp])
	})
	app.RegisterView(view)

	// Serve assets (e.g. javascript, css).
	// app.HandleDir("/public", iris.Dir("./public"))

	app.Get("/", index)
	app.Get("/upload", uploadView)
	app.Post("/upload", upload)
	app.PartyFunc("/admin", func(basic iris.Party) {
		app.Get("/status", GetAdminStatus)
	})
	filesRouter := app.Party("/files")
	filesRouter.HandleDir("/", iris.Dir(config.UploadDir), iris.DirOptions{
		Compress: true,
		ShowList: true,

		// Optionally, force-send files to the client inside of showing to the browser.
		Attachments: iris.Attachments{
			Enable: true,
			// Optionally, control data sent per second:
			Limit: 50.0 * iris.MB,
			Burst: 100 * iris.MB,
			// Change the destination name through:
			// NameFunc: func(systemName string) string {...}
		},

		DirList: iris.DirListRich(iris.DirListRichOptions{
			// Optionally, use a custom template for listing:
			// Tmpl: dirListRichTemplate,
			TmplName: "dirlist.html",
		}),
	})

	auth := basicauth.Default(map[string]string{
		"myusername": "mypassword",
	})

	filesRouter.Delete("/{file:path}", auth, deleteFile)
}
func index(ctx iris.Context) {
	ctx.Redirect("/upload")
}
