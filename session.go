package main

import (
	"log"
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
	sess = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
)

func setAuth(ctx iris.Context) {
	host := strings.Split(ctx.Host(), ":")[0]
	// cookieConfig := &http.Cookie{
	// 	Domain: host,
	// }
	ctx.SetCookieKV("Domain", host)
	session := sess.Start(ctx)
	session.Set("authenticated", true)

	//
	ctx.SetCookie(&http.Cookie{
		Domain: host,
	})
	log.Print(host)
	// s := strconv.Itoa(session.Len())
	// log.Print(s)
}
