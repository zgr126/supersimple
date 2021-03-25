package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
)

func uploadView(ctx iris.Context) {
	now := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(now, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))

	ctx.View("upload.html", token)
}

func upload(ctx iris.Context) {
	ctx.SetMaxRequestBodySize(int64(config.MaxUploadFileSize_kb))

	_, _, err := ctx.UploadFormFiles(config.UploadDir, beforeSave)
	if err != nil {
		ctx.StopWithError(iris.StatusPayloadTooRage, err)
		return
	}

	ctx.Redirect("/files")
}

func beforeSave(ctx iris.Context, file *multipart.FileHeader) bool {
	ip := ctx.RemoteAddr()
	ip = strings.ReplaceAll(ip, ".", "_")
	ip = strings.ReplaceAll(ip, ":", "_")

	file.Filename = ip + "-" + file.Filename
	return true
}

func deleteFile(ctx iris.Context) {
	// It does not contain the system path,
	// as we are not exposing it to the user.
	fileName := ctx.Params().Get("file")

	filePath := path.Join(config.UploadDir, fileName)

	if err := os.RemoveAll(filePath); err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}

	ctx.Redirect("/files")
}
