package main

import "github.com/kataras/iris/v12"

func GetAdminStatus(ctx iris.Context) {
	ctx.View("upload.html")
}
