package main

import "github.com/kataras/iris/v12"

func testRouter(ctx iris.Context) {
	ctx.Writef("not router: %s", ctx.Params().Get("name"))
}
