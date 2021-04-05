package main

import (
	"net/http"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

const (
	cookieNameForSessionID string = "supersimple"
	adminAuthStr           string = "systemAdmin"
)

var (
	sess = sessions.New(sessions.Config{
		Cookie:       cookieNameForSessionID,
		AllowReclaim: true,
	})
)

func setAuth(ctx iris.Context) {
	host := strings.Split(ctx.Host(), ":")[0]
	cookieConfig := func(ctx iris.Context, c *http.Cookie, i uint8) {
		c.Domain = host
	}
	session := sess.Start(ctx, cookieConfig)
	session.Set(adminAuthStr, true)

}
