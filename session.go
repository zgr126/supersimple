package main

import "github.com/kataras/iris/v12/sessions"

const (
	cookieNameForSessionID string = "supersimple"
	adminAuthStr           string = "systemAdmin"
)

var (
	sess = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
)
