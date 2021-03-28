package main

type errorCode int

const (
	jsonParseErr errorCode = 101 + iota
	dbErr
	userUploadErr
	authErr
)
