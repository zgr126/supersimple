package main

import (
	"encoding/json"
	"time"

	"github.com/kataras/iris/v12"
	bolt "go.etcd.io/bbolt"
)

var app *appStruct

type bean struct {
	Id        uint64     `json:"id"`
	FilePath  string     `json:"path"`
	Name      string     `json:"name"`
	Des       string     `json:"des"` // bean's description
	CreatTime string     `json:"time"`
	Status    beanStatue `json:"status"`

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
	b, err := tx.CreateBucketIfNotExists([]byte(newBean.Name))
	id, _ := b.NextSequence()
	_time := time.Now()
	timeFormat := _time.Format("2006-01-02 15:04:05")
	newBean.CreatTime = timeFormat
	newBean.Id = id
	newBean.Status = beanStatusActive
	beanValue, _ := json.Marshal(newBean)
	bucket := tx.Bucket(app_beans)

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
	tx, _ := db.Begin(true)
	bucket := tx.Bucket(app_beans)
	target := bucket.Get(itob(int(newBean.Id)))
	_target := &bean{}
	if len(target) != 0 {
		json.Unmarshal(target, _target)
		_target.Status = beanStatusDisable
		jsonStr, _ := json.Marshal(_target)
		bucket.Put(itob(int(newBean.Id)), jsonStr)
		setBeanStatus(tx, []byte(name), _target)
	}
	tx.Commit()
	commonResult(ctx, _target)
}

func setBeanStatus(tx *bolt.Tx, name []byte, _bean *bean) {
	b := tx.Bucket(name)
	s := b.Stats()
	_bean.ByteSize = uint64(s.InlineBucketInuse)
	_bean.KVsize = uint64(s.KeyN)
}
