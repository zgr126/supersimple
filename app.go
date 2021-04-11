package main

import (
	"encoding/json"

	"github.com/kataras/iris/v12"
)

var (
	app_system []byte = []byte("__system")
	app_beans  []byte = []byte("__beans")
)

type appStruct struct {
	Beans      beans             `json:"beans"`
	HttpHeader map[string]string `json:"httpHeader"`
}

type resultValue map[uint64][]byte
type resultValues []resultValue

func newApp() {
	app = &appStruct{}
	app.getBeans()
}

func (app *appStruct) getBeans() {
	tx, _ := db.Begin(true)
	b, _ := tx.CreateBucketIfNotExists(app_beans)
	app.Beans = []*bean{}
	_ = b.ForEach(func(k, v []byte) error {
		_b := &bean{}
		_ = json.Unmarshal(v, _b)
		// get length
		setBeanStatus(tx, []byte(_b.Name), _b)
		app.Beans = append(app.Beans, _b)
		return nil
	})
	tx.Commit()
}

func getApp(ctx iris.Context) {
	app.getBeans()
	commonResult(ctx, app)
}

func appRouter(ctx iris.Context) {
	// s := ctx.Clone().RouteName()
	name := ctx.Params().GetString("name")
	hasBean := app.Beans.hasBean(name)
	if hasBean {

	}
}

func appGet(ctx iris.Context) {
	name := ctx.Params().GetString("name")
	tx, _ := db.Begin(true)
	b := tx.Bucket([]byte(name))
	list := resultValues{}
	if b != nil {
		b.ForEach(func(k, v []byte) error {
			item := make(resultValue)
			item[btoui64(k)] = v
			list = append(list, item)
			return nil
		})
		// log.Print(list)
	}
	tx.Commit()
	commonResult(ctx, list)

}

func appPostPage(ctx iris.Context) {

}

func appPost(ctx iris.Context) {

}

func appPut(ctx iris.Context) {

}

func appDelete(ctx iris.Context) {

}

func appAny(ctx iris.Context) {

}
