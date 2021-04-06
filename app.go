package main

import (
	"github.com/kataras/iris/v12"
	bolt "go.etcd.io/bbolt"
)

const (
	bean_system string = "__system"
)

type app struct {
	Beans      beans             `json:"beans"`
	HttpHeader map[string]string `json:"httpHeader"`
}

type bean struct {
	IsFileServer bool
	Name         string
	Num          uint64
	Status       beanStatue
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
	_app := &app{
		Beans: list,
	}
	db.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(name []byte, _ *bolt.Bucket) error {
			_b := &bean{
				Name: string(name),
			}
			list = append(list, _b)
			return nil
		})
	})
	commonResult(ctx, _app)
}
