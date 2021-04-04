package main

import (
	"crypto/sha256"
	"encoding/hex"
	"reflect"
)

func CopyStruct(src, dst interface{}) {
	sval := reflect.ValueOf(src).Elem()
	dval := reflect.ValueOf(dst).Elem()

	for i := 0; i < sval.NumField(); i++ {
		value := sval.Field(i)
		name := sval.Type().Field(i).Name

		dvalue := dval.FieldByName(name)
		if dvalue.IsValid() == false {
			continue
		}
		dvalue.Set(value)
	}
}

func cryptoByte(b []byte) string {
	h := sha256.New()
	_b := h.Sum(b)

	return hex.EncodeToString(_b)
}
