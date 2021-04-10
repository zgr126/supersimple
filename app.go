package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/kataras/iris/v12"
	bolt "go.etcd.io/bbolt"
)

var (
	app_system []byte = []byte("__system")
	app_beans  []byte = []byte("__beans")
)

type appStruct struct {
	Beans      beans             `json:"beans"`
	HttpHeader map[string]string `json:"httpHeader"`
}

var app *appStruct

type bean struct {
	FilePath  string     `json:"path"`
	Name      string     `json:"name"`
	Num       uint64     `json:"num"`
	Doc       string     `json:"doc"`
	CreatTime time.Time  `json:"time"`
	Status    beanStatue `json:"status"`
}

type beanStatue uint8

const (
	beanStatusNotExist beanStatue = iota
	beanStatusActive
	beanStatusDelete
)

type beans []*bean

func newApp() {
	app = &appStruct{}
	app.getBeans()
}

func (app *appStruct) addBean() {

}

func (app *appStruct) getBeans() {
	tx, _ := db.Begin(true)
	b, _ := tx.CreateBucketIfNotExists(app_beans)
	app.Beans = []*bean{}
	_ = b.ForEach(func(k, v []byte) error {
		_b := &bean{}
		_ = json.Unmarshal(v, _b)
		app.Beans = append(app.Beans, _b)
		return nil
	})
	tx.Commit()
}
func newBean() *bean {
	b := &bean{
		CreatTime: time.Now(),
		Status:    beanStatusActive,
	}
	return b
}
func getApp(ctx iris.Context) {
	app.getBeans()
	commonResult(ctx, app)
}

func addBean(ctx iris.Context) {
	type input struct {
		Name     string `json:"name"`
		Doc      string `json:"doc"`
		FilePath string `json:"filePath"`
	}
	_i := &input{}
	ctx.ReadJSON(_i)

	tx, _ := db.Begin(true)
	_, err := tx.CreateBucketIfNotExists([]byte(_i.Name))
	b := newBean()
	b.Name = _i.Name
	b.Doc = _i.Doc
	b.FilePath = _i.FilePath
	beanValue, _ := json.Marshal(b)
	bucket := tx.Bucket(app_beans)

	err = bucket.Put([]byte(_i.Name), beanValue)
	if err = tx.Commit(); err != nil {
		app.getBeans()
		errorHandleJSON(ctx, err, dbErr)
	} else {
		commonResult(ctx, nil)
	}
}

func disableBean(ctx iris.Contest) {

}

func appauth(ctx iris.Context) {

}

func appRouter(ctx iris.Context) {
	s := ctx.Clone().RouteName()
	log.Print(s)
	name := ctx.Params().GetString("name")
	log.Print(name)

}

func testGetAll(ctx iris.Context) {
	list := beans{}

	db.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(name []byte, _ *bolt.Bucket) error {
			_b := &bean{
				Name: string(name),
			}
			list = append(list, _b)
			return nil
		})
	})

	_app := &appStruct{
		Beans: list,
	}
	commonResult(ctx, _app)
}
