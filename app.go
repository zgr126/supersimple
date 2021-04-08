package main

import (
	"strings"

	"github.com/kataras/iris/v12"
	bolt "go.etcd.io/bbolt"
)

const (
	app_system string = "__system"
)

type app struct {
	Beans      beans             `json:"beans"`
	HttpHeader map[string]string `json:"httpHeader"`
}

type bean struct {
	IsFileServer bool       `json:"isFileServer"`
	Name         string     `json:"name"`
	Num          uint64     `json:"num"`
	Status       beanStatue `json:"status"`
}

type beanStatue uint8

const (
	beanStatusNotExist beanStatue = iota
	beanStatusActive
	beanStatusDelete
)

type beans []*bean

func getApp(ctx iris.Context) {
	list := beans{}

	db.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(name []byte, _ *bolt.Bucket) error {
			//remove system beans '__xxxxx'
			if strings.Index(string(name), "__") != 0 {
				_b := &bean{
					Name: string(name),
				}
				list = append(list, _b)
			}
			return nil
		})
	})

	_app := &app{
		Beans: list,
	}
	commonResult(ctx, _app)
}

func addBean(ctx iris.Context) {
	type input struct {
		Name string `json:"name"`
		Doc  string `json:"doc"`
	}
	_i := &input{}
	ctx.ReadJSON(_i)
	tx, _ := db.Begin(true)
	_, err := tx.CreateBucketIfNotExists([]byte(_i.Name))
	tx.Commit()
	if err != nil {
		errorHandleJSON(ctx, err, dbErr)
	} else {
		commonResult(ctx, nil)
	}
}
