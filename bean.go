package main

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/kataras/iris/v12"
	bolt "go.etcd.io/bbolt"
)

var app *appStruct

type bean struct {
	Id           uint64     `json:"id"`
	IsFileServer bool       `json:"isFileServer"`
	Name         string     `json:"name"`
	Des          string     `json:"des"` // bean's description
	CreatTime    string     `json:"time"`
	Status       beanStatue `json:"status"`

	KVsize   uint64 `json:"kvSize"`
	ByteSize uint64 `json:"byteSize"`
}

type beanStatue uint8

const (
	beanStatusNotExist beanStatue = iota
	beanStatusActive
	beanStatusDisable
	beanStatusDelete
)

type beans []*bean

type fileItem struct {
	Name      string `json:"name"`
	IP        string `json:"ip"`
	Uuid      string `json:uuid`
	CreatTime string `json:time`
}

func (b beans) hasBean(s string) bool {

	for _, v := range b {
		if v.Name == s {
			return true
		}
	}
	return false
}

func addBean(ctx iris.Context) {
	newBean := &bean{}
	ctx.ReadJSON(newBean)

	tx, _ := db.Begin(true)
	_b := tx.Bucket([]byte(newBean.Name))
	if _b != nil {
		tx.Commit()
		errorHandleJSON(ctx, errors.New("document exsit"), dbErr)
		return
	}
	_, err := tx.CreateBucketIfNotExists([]byte(newBean.Name))
	bucket := tx.Bucket(app_beans)
	id, _ := bucket.NextSequence()
	_time := time.Now()
	timeFormat := _time.Format("2006-01-02 15:04:05")
	newBean.CreatTime = timeFormat
	newBean.Id = id
	newBean.Status = beanStatusActive
	beanValue, _ := json.Marshal(newBean)

	err = bucket.Put(itob(int(newBean.Id)), beanValue)
	if err = tx.Commit(); err != nil {
		app.getBeans()
		errorHandleJSON(ctx, err, dbErr)
	} else {
		commonResult(ctx, nil)
	}
}

func setBean(ctx iris.Context) {
	newBean := &bean{}
	ctx.ReadJSON(newBean)
	name := newBean.Name
	// if delete bean
	if newBean.Status == beanStatusDelete {
		deleteBean(ctx, newBean)
		return
	}
	tx, _ := db.Begin(true)
	bucket := tx.Bucket(app_beans)
	target := bucket.Get(itob(int(newBean.Id)))
	if len(target) != 0 {
		jsonStr, _ := json.Marshal(newBean)
		bucket.Put(itob(int(newBean.Id)), jsonStr)
		setBeanStatus(tx, []byte(name), newBean)
	}
	tx.Commit()
	commonResult(ctx, newBean)
}

func deleteBean(ctx iris.Context, _bean *bean) {
	tx, _ := db.Begin(true)
	byteName := []byte(_bean.Name)
	b := tx.Bucket(byteName)
	if b != nil {
		tx.DeleteBucket(byteName)
	}
	_b := tx.Bucket(app_beans)
	_b.Delete(itob(int(_bean.Id)))
	tx.Commit()
	commonResult(ctx, nil)
}

func setBeanStatus(tx *bolt.Tx, name []byte, _bean *bean) {
	b := tx.Bucket(name)
	s := b.Stats()
	_bean.ByteSize = uint64(s.InlineBucketInuse)
	_bean.KVsize = uint64(s.KeyN)
}
