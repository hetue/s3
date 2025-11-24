package step

import (
	"context"
	"fmt"
	"os"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/hetue/boot"
	"github.com/hetue/git/internal/internal/config"
	"github.com/hetue/git/internal/internal/step/internal/constant"
)

var _ boot.Step = (*Credential)(nil)

type Credential struct {
	runtime    *boot.Runtime
	credential *config.Credential
	repository *config.Repository
	logger     log.Logger
}

func newCredential(runtime *boot.Runtime, credential *config.Credential, repository *config.Repository, logger log.Logger) *Credential {
	return &Credential{
		runtime:    runtime,
		credential: credential,
		repository: repository,
		logger:     logger,
	}
}

func (c *Credential) Name() string {
	return "授权"
}

func (c *Credential) Runnable() bool {
	return "" != c.credential.Username && "" != c.credential.Password // nolint:staticcheck
}

func (c *Credential) Retryable() bool { // 不重试
	return false
}

func (c *Credential) Asyncable() bool { // 不异步
	return false
}

func (c *Credential) Run(_ *context.Context) (err error) {
	filepath := c.runtime.Path(constant.NetrcFilename)
	if _, se := os.Stat(filepath); nil == se || os.IsExist(se) {
		_ = os.Remove(filepath)
	}

	fields := gox.Fields[any]{
		field.New("filepath", filepath),
		field.New("username", c.credential.Username),
		field.New("password", c.credential.Password),
	}

	content := fmt.Sprintf(constant.NetrcConfigFormatter, c.credential.Username, c.credential.Password)
	if err = os.WriteFile(filepath, []byte(content), constant.DefaultFilePerm); nil != err {
		c.logger.Error("写入授权文件出错", fields.Add(field.Error(err))...)
	} else {
		c.logger.Info("写入授权文件成功", fields...)
	}

	return
}
