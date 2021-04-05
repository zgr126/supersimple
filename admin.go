package main

import (
	"errors"
	"time"

	"github.com/kataras/iris/v12"
	bolt "go.etcd.io/bbolt"
)

type adminStruct struct {
	password   string `json:"" `
	CreateTime string `json:"createTime"`
	db         *bolt.DB
}

type upload_Pass struct {
	Password string `json:"password" validate:"required,min=8,max=300"`
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
	return len(ad.password) != 0
}
func (ad *adminStruct) loadAdminDetail() {
	tx, _ := db.Begin(true)
	defer tx.Commit()
	b, _ := tx.CreateBucketIfNotExists([]byte(bean_system))
	v := b.Get(adminPassword_s)
	t := b.Get(adminCreatTime_s)

	admin.password = string(v)
	if len(t) != 0 {
		admin.CreateTime = string(t)

	}
}
func (ad *adminStruct) setPassword(s string) {
	// sum := sha256.Sum256([]byte(s))
	// tx, err := ad.db.Begin(true)
}

func (ad *adminStruct) getConfig() *adminStruct {
	_c := *ad
	_c.password = ""
	return &_c
}

func getAdminStatus(ctx iris.Context) {
	commonResult(ctx, admin.getConfig())
}

func setAdminPassword(ctx iris.Context) {
	if admin.password != "" {
		errorHandleJSON(ctx, errors.New("Password exists"), authErr)
		return
	}
	c := &upload_Pass{}
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
				errorHandleJSON(ctx, err, dbErr)
				return
			}
			b := tx.Bucket([]byte("system"))
			p := cryptoByte([]byte(c.Password))
			b.Put(adminPassword_s, []byte(p))
			t := time.Now()
			b.Put(adminCreatTime_s, []byte(t.String()))
			err = tx.Commit()
			if err != nil {
				errorHandleJSON(ctx, err, dbErr)
				return
			}
			admin.password = p
			admin.CreateTime = t.String()
			setAuth(ctx)
			commonResult(ctx, nil)
		}
	}
}

func removeAuth(ctx iris.Context) {

}

func login(ctx iris.Context) {

	c := &upload_Pass{}
	if err := ctx.ReadJSON(c); err != nil {
		errorHandleJSON(ctx, err, userUploadErr)
		return
	} else {
		p := cryptoByte([]byte(c.Password))
		if p == admin.password {
			setAuth(ctx)
			commonResult(ctx, nil)

		} else {
			ctx.StatusCode(iris.StatusForbidden)
			errorHandleJSON(ctx, errors.New("password not pass"), userUploadErr)
		}
	}
}

func logout(ctx iris.Context) {
	sess.Start(ctx).Clear()
	sess.Destroy(ctx)
	commonResult(ctx, nil)
}
