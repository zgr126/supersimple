package main

import "github.com/kataras/iris/v12/sessions"

var (
	cookieNameForSessionID = "chaojijian"
	sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
)
