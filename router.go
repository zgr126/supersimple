package main

import (
	"encoding/json"
	"errors"
	"fmt"
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

	app.Use(cors)
	app.OnAnyErrorCode(errorHandle)
	// Serve assets (e.g. javascript, css).
	// app.HandleDir("/public", iris.Dir("./public"))
	app.Get("/", index)
	app.Any("/{name}", testRouter)
	app.Get("/upload", uploadView)
	app.Post("/upload", upload)
	app.PartyFunc("/admin", func(basic iris.Party) {
		basic.Get("/status", authContinue, getAdminStatus)
		basic.Post("/password", auth, setAdminPassword)
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
	ctx.Next()
}
func auth(ctx iris.Context) {

	session := sess.Start(ctx)
	auth, _ := session.GetBoolean(adminAuthStr)

	if !auth {
		fmt.Print("not auth \n")
		ctx.StatusCode(iris.StatusForbidden)
	}
	fmt.Print("auth")

}

func setCors(ctx iris.Context) {
	var Origin = ctx.Request().Header["Origin"]
	ctx.Header("Access-Control-Allow-Origin", "")
	ctx.Header("Access-Control-Allow-Origin", strings.Join(Origin, ""))
	ctx.Header("Access-Control-Allow-Credentials", "")
	ctx.Header("Access-Control-Allow-Credentials", "true")
}
func cors(ctx iris.Context) {
	setCors(ctx)
	ctx.Next()
}

func errorHandle(ctx iris.Context) {
	fmt.Print("errorcode")
	setCors(ctx)
	if ctx.Request().Method == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		ctx.Header("Access-Control-Max-Age", "2592000")
		ctx.StatusCode(204)
		return
	}
}

// common error response
func errorHandleJSON(ctx iris.Context, err error, code errorCode) {
	var _json = &ResponseBean{
		Code: int(code),
		Msg:  err.Error(),
	}
	var _byte, _ = json.Marshal(_json)
	ctx.Binary(_byte)
}

// basic response null data
func basicJSON(ctx iris.Context) {
	var _json = &ResponseBean{
		Code: 100,
	}
	var _byte, _ = json.Marshal(_json)
	ctx.Binary(_byte)
}

// basic response
func commonResponseJSON(ctx iris.Context, i interface{}) {
	var _json = &ResponseBean{
		Code: 100,
		Data: i,
		Msg:  "",
	}
	_byte, err := json.Marshal(_json)
	if err != nil {
		errorHandleJSON(ctx, errors.New("some error, please retry"), jsonParseErr)
	} else {
		ctx.Binary(_byte)
	}
}
