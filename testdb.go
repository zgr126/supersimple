package main

import (
	"log"

	"github.com/kataras/iris/v12"
	bolt "go.etcd.io/bbolt"
)

func testGetAll(ctx iris.Context) {
	list := beans{}

	db.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(name []byte, b *bolt.Bucket) error {
			_b := &bean{
				Name: string(name),
			}
			log.Print(b.Stats())
			list = append(list, _b)
			return nil
		})
	})

	_app := &appStruct{
		Beans: list,
	}
	commonResult(ctx, _app)
}
