package step

import (
	"context"
	"io"
	"mime"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/goexl/gfx"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/hetue/s3/internal/internal/config"
	"github.com/hetue/s3/internal/internal/step/internal"
)

type Upload struct {
	source *config.Source
	server *config.Server

	paths  []string
	client *s3.Client
	logger log.Logger
}

func newUpload(source *config.Source, server *config.Server, client *internal.Client, logger log.Logger) *Upload {
	return &Upload{
		source: source,
		server: server,

		client: client,
		logger: logger,
	}
}

func (u *Upload) Runnable() (runnable bool) {
	if paths, ae := gfx.All(u.source.Folder); nil == ae || 0 != len(paths) {
		runnable = true
		u.paths = paths
	}

	return
}

func (u *Upload) Run(ctx *context.Context) (err error) {
	for _, path := range u.paths {
		if err = u.run(ctx, path); nil != err {
			return
		}
	}

	return
}

func (u *Upload) Name() string {
	return "上传"
}

func (u *Upload) Retryable() bool {
	return true
}

func (u *Upload) Asyncable() bool {
	return false
}

func (u *Upload) run(ctx *context.Context, path string) (err error) {
	if really, re := filepath.Rel(u.source.Folder, path); nil != re {
		err = re
		u.logger.Error("获取文件相对路径出错", field.New("path", path), field.Error(err))
	} else if body, oe := os.Open(path); nil != oe {
		err = oe
	} else {
		err = u.upload(ctx, really, body)
	}

	return
}

func (u *Upload) upload(ctx *context.Context, path string, body io.Reader) (err error) {
	poi := new(s3.PutObjectInput)
	poi.Bucket = aws.String(u.server.Bucket)
	poi.Body = body
	poi.ContentType = aws.String(mime.TypeByExtension(filepath.Ext(path)))

	paths := strings.Split(path, string(filepath.Separator))
	if "" != u.source.Prefix {
		paths = append([]string{u.source.Prefix}, paths...)
	}
	if "" != u.source.Suffix {
		paths = append(paths, u.source.Suffix)
	}
	poi.Key = aws.String(strings.Join(paths, u.server.Separator))

	fields := gox.Fields[any]{
		field.New("path", path),
	}
	if out, poe := u.client.PutObject(*ctx, poi); nil != poe {
		err = poe
		u.logger.Error("上传文件出错", field.Error(err), fields...)
	} else if nil == out {
		u.logger.Warn("上传文件失败", fields[0], fields[1:]...)
	} else {
		u.logger.Debug("文件上传成功", fields[0], fields[1:]...)
	}

	return
}
