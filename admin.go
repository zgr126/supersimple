package main

import (
	"errors"
	"log"
	"time"

	"github.com/kataras/iris/v12"
	bolt "go.etcd.io/bbolt"
)

type adminStruct struct {
	Password   string    `json: "password"`
	CreateTime time.Time `json: "createTime"`
	db         *bolt.DB
}

var admin *adminStruct

var (
	adminPassword_s  []byte = []byte("adminPassword")
	adminCreatTime_s []byte = []byte("createTime")
)

// load Admin config
func newAdmin(db *bolt.DB) {
	admin = &adminStruct{
		db: db,
	}
	admin.loadAdminDetail()
}

func (ad *adminStruct) hasConfig() bool {
	return len(ad.Password) != 0
}
func (ad *adminStruct) loadAdminDetail() {
	tx, _ := db.Begin(true)
	defer tx.Commit()
	b, _ := tx.CreateBucketIfNotExists([]byte("system"))
	v := b.Get(adminPassword_s)
	t := b.Get(adminCreatTime_s)
	admin.Password = string(v)
	if len(t) != 0 {
		_time, _ := time.Parse("2006-01-02 15:04:05", string(t))
		admin.CreateTime = _time
	}
}
func (ad *adminStruct) setPassword(s string) {
	// sum := sha256.Sum256([]byte(s))
	// tx, err := ad.db.Begin(true)

}

func (ad *adminStruct) getConfig() *adminStruct {
	_c := *ad
	_c.Password = ""
	return &_c
}

func getAdminStatus(ctx iris.Context) {
	if !admin.hasConfig() {
		basicJSON(ctx)
	} else {
		commonResponseJSON(ctx, admin.getConfig())
	}
	ctx.Next()

}
func setAdminPassword(ctx iris.Context) {
	if admin.Password != "" {
		errorHandleJSON(ctx, errors.New("Password exists"), authErr)
		return
	}
	c := &adminStruct{}
	if err := ctx.ReadJSON(c); err != nil {
		errorHandleJSON(ctx, err, userUploadErr)
		return
	} else {
		if len(c.Password) == 0 {
			errorHandleJSON(ctx, errors.New("Upload value not format"), userUploadErr)
			return
		} else {
			tx, err := admin.db.Begin(true)
			if err != nil {
				errorHandleJSON(ctx, errors.New("Upload value not format"), userUploadErr)
				return
			}
			b := tx.Bucket([]byte("system"))
			p := cryptoByte(c.Password)
			log.Print(p)
			b.Put(adminPassword_s, []byte(p))
			err = tx.Commit()
			if err != nil {
				errorHandleJSON(ctx, err, dbErr)
				return
			}
			admin.Password = p
			session := sess.Start(ctx)
			session.Set("authenticated", true)
			basicJSON(ctx)
		}
	}
}

func login(ctx iris.Context) {

}

func logout(ctx iris.Context) {

}
