package main

import (
	"encoding/json"
	"errors"
	"path/filepath"

	"github.com/go-basic/uuid"
	"github.com/kataras/iris/v12"
)

var (
	app_system  []byte = []byte("__system")
	app_beans   []byte = []byte("__beans")
	app_setting []byte = []byte("setting")
)

type appStruct struct {
	Beans   beans   `json:"beans"`
	Setting setting `json:"setting"`
}

type beanResult map[uint64][]byte
type beanResults []beanResult

type setting struct {
	Headers []*settingHttpHeader `json:"headers"`
}

type settingHttpHeader struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func newApp() {
	app = &appStruct{}
	app.getBeans()
	app.getSetting()
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

func (app *appStruct) getSetting() {
	tx, _ := db.Begin(true)
	b, _ := tx.CreateBucketIfNotExists(app_system)

	v := b.Get(app_setting)
	_setting := &setting{}
	_ = json.Unmarshal(v, _setting)
	app.Setting = *_setting
	tx.Commit()
}

func appSetting(ctx iris.Context) {
	newSetting := &setting{}
	ctx.ReadJSON(newSetting)

	tx, _ := db.Begin(true)
	b, _ := tx.CreateBucketIfNotExists(app_system)

	settingByte, _ := json.Marshal(newSetting)
	_ = b.Put(app_setting, settingByte)
	app.Setting = *newSetting
	tx.Commit()
	commonResult(ctx, nil)
}

func getApp(ctx iris.Context) {
	app.getBeans()
	app.getSetting()
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
	list := beanResults{}
	if b != nil {
		b.ForEach(func(k, v []byte) error {
			item := make(beanResult)
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
	name := ctx.Params().GetString("name")

	_bean := app.Beans.getBean(name)
	if _bean == nil {
		errorHandleJSON(ctx, errors.New("can't find router"), routerErr)
		return
	}
	if _bean.IsFileServer {
		postFile(ctx)
	} else {

	}

	tx, _ := db.Begin(true)
	defer tx.Commit()
	b := tx.Bucket([]byte(name))
	result := make(beanResult)
	if b != nil {
		_id, _ := b.NextSequence()
		_body, _ := ctx.GetBody()
		if len(_body) == 0 {
			errorHandleJSON(ctx, errors.New("body length is 0"), userUploadErr)
			return
		}

		b.Put(itob(int(_id)), _body)
		result[_id] = b.Get(itob(int(_id)))
	}

	commonResult(ctx, result)
}

func appPut(ctx iris.Context) {

}

func appDelete(ctx iris.Context) {

}

func appAny(ctx iris.Context) {

}

type postFileResult struct {
	FileName string `json:"filename"`
	RawFile  string `json:"rawFile"`
	Size     int64  `json:"size"`
	Id       uint64 `json:"id"`
}

func postFile(ctx iris.Context) {
	f, fh, _ := ctx.FormFile("uploadfile")
	_uuid := uuid.New()
	// if err != nil {
	// 	ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
	// 	return
	// }
	defer f.Close()
	name := ctx.Params().GetString("name")
	_, _ = ctx.SaveFormFile(fh, filepath.Join("./"+name, _uuid))
	// if err != nil {
	// 	ctx.StatusCode(iris.StatusInternalServerError)
	// 	ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
	// 	return
	// }

	tx, _ := db.Begin(true)
	defer tx.Commit()
	b := tx.Bucket([]byte(name))
	if b != nil {
		_id, _ := b.NextSequence()
		_body := postFileResult{
			FileName: _uuid,
			Size:     fh.Size,
			RawFile:  filepath.Join("./"+name, _uuid),
			Id:       _id,
		}
		__b, _ := json.Marshal(_body)
		b.Put(itob(int(_id)), __b)
		commonResult(ctx, _body)
		return
	}
	errorHandleJSON(ctx, errors.New("server error"), dbErr)

}
