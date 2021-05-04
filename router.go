package main

import (
	"errors"
	"strings"

	"github.com/kataras/iris/v12"
)

type ResponseBean struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var dirOptions = iris.DirOptions{
	IndexName: "index.html",
	// The `Compress` field is ignored
	// when the file is cached (when Cache.Enable is true),
	// because the cache file has a map of pre-compressed contents for each encoding
	// that is served based on client's accept-encoding.
	Compress: true, // true or false does not matter here.
	Cache: iris.DirCacheOptions{
		Enable:         true,
		CompressIgnore: iris.MatchImagesAssets,
		// Here, define the encodings that the cached files should be pre-compressed
		// and served based on client's needs.
		Encodings:       []string{"gzip", "deflate", "br", "snappy"},
		CompressMinSize: 50, // files smaller than this size will NOT be compressed.
		Verbose:         1,
	},
}

func setRouter(app *iris.Application) {
	// app.UseRouter(crs)

	app.Use(xcors)
	app.OnErrorCode(403, errorHandle403)
	app.OnErrorCode(404, errorHandle404)
	app.Done(routeDone)
	app.Get("/", index)

	app.Get("/upload", uploadView)
	app.Post("/upload", upload)
	app.PartyFunc("/app", func(_app iris.Party) {
		_app.Get("/{name}", appGet)
		_app.Post("/{name}/page", appPostPage)
		_app.Post("/{name}", appPost)
		_app.Post("/{name}/batch", appPost)
		_app.Put("/{name}", appPut)
		_app.Put("/{name}/batch", appPut)
		_app.Delete("/{name}", appDelete)
		_app.Delete("/{name}/batch", appPut)
	})
	app.HandleDir("/admin", GetFileSystem(false, "admin/dist", staticFiles), dirOptions)
	app.PartyFunc("/adminRest", func(adminRouter iris.Party) {

		adminRouter.Get("/status", getAdminStatus)
		adminRouter.Post("/setPassword", setAdminPassword)
		adminRouter.Get("/app", getApp)
		adminRouter.Post("/login", login)
		adminRouter.Post("/logout", logout)
		//beans
		adminRouter.Post("/addBean", addBean)
		adminRouter.Post("/setBean", setBean)
		//setting
		adminRouter.Post("/setting", appSetting)
	})

	app.PartyFunc("/test", func(test iris.Party) {
		test.Get("/db", testGetAll)
	})

	// {
	// 	view := iris.HTML("./views", ".html")
	// 	view.AddFunc("formatBytes", func(b int64) string {
	// 		const unit = 1000
	// 		if b < unit {
	// 			return fmt.Sprintf("%d B", b)
	// 		}
	// 		div, exp := int64(unit), 0
	// 		for n := b / unit; n >= unit; n /= unit {
	// 			div *= unit
	// 			exp++
	// 		}
	// 		return fmt.Sprintf("%.1f %cB",
	// 			float64(b)/float64(div), "kMGTPE"[exp])
	// 	})
	// 	app.RegisterView(view)
	// 	filesRouter := app.Party("/files")
	// 	filesRouter.HandleDir("/", iris.Dir(config.UploadDir), iris.DirOptions{
	// 		Compress: true,
	// 		ShowList: true,

	// 		// Optionally, force-send files to the client inside of showing to the browser.
	// 		Attachments: iris.Attachments{
	// 			Enable: true,
	// 			// Optionally, control data sent per second:
	// 			Limit: 50.0 * iris.MB,
	// 			Burst: 100 * iris.MB,
	// 			// Change the destination name through:
	// 			// NameFunc: func(systemName string) string {...}
	// 		},

	// 		DirList: iris.DirListRich(iris.DirListRichOptions{
	// 			// Optionally, use a custom template for listing:
	// 			// Tmpl: dirListRichTemplate,
	// 			TmplName: "dirlist.html",
	// 		}),
	// 	})

	// 	auth := basicauth.Default(map[string]string{
	// 		"myusername": "mypassword",
	// 	})

	// 	filesRouter.Delete("/{file:path}", auth, deleteFile)
	// }
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

func routeDone(ctx iris.Context) {

}

// common result
func commonResult(ctx iris.Context, data interface{}) {
	ctx.ContentType("application/json")
	var _json = &ResponseBean{
		Code: 100,
		Data: data,
	}
	// _byte, err := json.Marshal(_json)
	_, err := ctx.JSON(_json)
	if err != nil {
		errorHandleJSON(ctx, errors.New("some error, please retry"), jsonParseErr)
	} else {
		// ctx.Binary(_byte)

	}
}

// error result
func errorHandleJSON(ctx iris.Context, err error, code errorCode) {
	ctx.ContentType("application/json")
	var _json = &ResponseBean{
		Code: int(code),
		Msg:  err.Error(),
	}
	// var _byte, _ = json.Marshal(_json)
	// ctx.Binary(_byte)
	ctx.JSON(_json)
}
