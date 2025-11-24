package step

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/hetue/boot"
	"github.com/hetue/git/internal/internal/config"
	"github.com/hetue/git/internal/internal/step/internal/constant"
)

type SSH struct {
	runtime    *boot.Runtime
	credential *config.Credential
	logger     log.Logger
}

func newSSH(runtime *boot.Runtime, credential *config.Credential, logger log.Logger) *SSH {
	return &SSH{
		runtime:    runtime,
		credential: credential,
		logger:     logger,
	}
}

func (*SSH) Name() string {
	return "密钥"
}

func (s *SSH) Runnable() bool {
	return "" != s.credential.Key // nolint:staticcheck
}

func (s *SSH) Retryable() bool { // 不重试
	return false
}

func (s *SSH) Asyncable() bool { // 不异步
	return false
}

func (s *SSH) Run(_ *context.Context) (err error) {
	home := s.runtime.Path(constant.SSHHome)
	keyfile := filepath.Join(home, constant.SSHKeyFilename)
	configFile := filepath.Join(home, constant.SSHConfigDir)
	if mhe := s.makeHome(home); nil != mhe {
		err = mhe
	} else if wke := s.writeKey(keyfile); nil != wke {
		err = wke
	} else {
		err = s.writeConfig(configFile, keyfile)
	}

	return
}

func (s *SSH) makeHome(home string) (err error) {
	homeField := field.New("home", home)
	if err = os.MkdirAll(home, 0o750); nil != err {
		s.logger.Error("创建SSH目录出错", homeField, field.Error(err))
	}

	return
}

func (s *SSH) writeKey(keyfile string) (err error) {
	key := s.credential.Key
	keyfileField := field.New("keyfile", keyfile)
	// 必须以换行符结束
	if !strings.HasSuffix(key, "\n") {
		key = fmt.Sprintf("%s\n", key)
	}

	if err = os.WriteFile(keyfile, []byte(key), constant.DefaultFilePerm); nil != err {
		s.logger.Error("写入密钥文件出错", keyfileField, field.Error(err))
	}

	return
}

func (s *SSH) writeConfig(configFile string, keyfile string) (err error) {
	configFileField := field.New("file", configFile)
	content := fmt.Sprintf(constant.SSHConfigFormatter, keyfile)
	if err = os.WriteFile(configFile, []byte(content), constant.DefaultFilePerm); nil != err {
		s.logger.Error("写入SSH配置文件出错", configFileField, field.Error(err))
	}

	return
}
