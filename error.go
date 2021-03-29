package main

type errorCode int

const (
	jsonParseErr  errorCode = 101 + iota // json parse error
	dbErr                                // bolt db error
	userUploadErr                        // user upload fields/file not invalid
	authErr                              // auth

)
