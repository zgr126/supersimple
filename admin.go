package main

import (
	"fmt"
	"time"

	"github.com/kataras/iris/v12"
)

type adminStruct struct {
	Password string
	Time     time.Time
	BoxLst   []string
}

func GetAdminStatus(ctx iris.Context) {
	// Start a writable transaction.
	tx, err := db.Begin(true)
	defer tx.Commit()
	if err != nil {
		errorHandleJSON(ctx, err, dbErr)
	}
	b, err := tx.CreateBucketIfNotExists([]byte("system"))
	if err != nil {
		errorHandleJSON(ctx, err, dbErr)
	}
	v := b.Get([]byte("adminPassword"))

	fmt.Println(v)
	// c := b.Cursor()
	// _, v := c.Seek([]byte("adminPassword"))
	if len(v) == 0 {
		basicJSON(ctx)
	} else {
		commonResponseJSON(ctx, &adminStruct{
			Password: string(v),
		})
	}
}
