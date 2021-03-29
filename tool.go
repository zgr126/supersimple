package main

import (
	"crypto/sha256"
	"fmt"
	"io"
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

func cryptoByte(b string) string {
	h := sha256.New()
	io.WriteString(h, b)
	sum := fmt.Sprintf("%x", b)
	return sum
}
