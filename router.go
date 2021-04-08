package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/basicauth"
)

type ResponseBean struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func setRouter(app *iris.Application) {
	// app.UseRouter(crs)

	app.Use(xcors)
	app.OnErrorCode(403, errorHandle403)
	app.OnErrorCode(404, errorHandle404)

	app.Get("/", index)
	app.Any("/{name}", testRouter)
	app.Get("/upload", uploadView)
	app.Post("/upload", upload)
	app.PartyFunc("/admin", func(adminRouter iris.Party) {
		adminRouter.Get("/status", getAdminStatus)
		adminRouter.Post("/setPassword", setAdminPassword)
		adminRouter.Get("/app", authContinue, getApp)
		adminRouter.Post("/login", login)
		adminRouter.Post("/logout", logout)
		//beans
		adminRouter.Post("/addBean", addBean)
	})

	{
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
}
func index(ctx iris.Context) {
	// ctx.Redirect("/upload")
}

func authContinue(ctx iris.Context) {
	auth(ctx)
	i := ctx.GetStatusCode()
	if i != 200 {
		return
	}
	ctx.Next()
}
func auth(ctx iris.Context) {
	s := ctx.Clone().RouteName()
	log.Print(s)
	session := sess.Start(ctx)
	auth, _ := session.GetBoolean(adminAuthStr)
	if !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
}

func setCors(ctx iris.Context) {
	var Origin = ctx.Request().Header["Origin"]
	ctx.Header("Access-Control-Allow-Origin", "")
	ctx.Header("Access-Control-Allow-Origin", strings.Join(Origin, ""))
	ctx.Header("Access-Control-Allow-Credentials", "")
	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.Header("Access-Control-Allow-Methods", "")
	ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
}

func xcors(ctx iris.Context) {
	setCors(ctx)
	ctx.Next()
}

func errorHandle404(ctx iris.Context) {
	setCors(ctx)
	if ctx.Request().Method == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "")
		ctx.Header("Access-Control-Allow-Headers", "")
		ctx.Header("Access-Control-Max-Age", "")
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		ctx.Header("Access-Control-Max-Age", "2592000")
		ctx.StatusCode(204)
		return
	}
}
func errorHandle403(ctx iris.Context) {

}

// common result
func commonResult(ctx iris.Context, data interface{}) {
	var _json = &ResponseBean{
		Code: 100,
		Data: data,
	}
	_byte, err := json.Marshal(_json)
	if err != nil {
		errorHandleJSON(ctx, errors.New("some error, please retry"), jsonParseErr)
	} else {
		ctx.Binary(_byte)
	}
}

// error result
func errorHandleJSON(ctx iris.Context, err error, code errorCode) {
	var _json = &ResponseBean{
		Code: int(code),
		Msg:  err.Error(),
	}
	var _byte, _ = json.Marshal(_json)
	ctx.Binary(_byte)
}
