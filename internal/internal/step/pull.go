package step

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/goexl/args"
	"github.com/goexl/exception"
	"github.com/goexl/gfx"
	"github.com/goexl/gox/field"
	"github.com/hetue/boot"
	"github.com/hetue/git/internal/internal/config"
	"github.com/hetue/git/internal/internal/step/internal/command"
	"github.com/hetue/git/internal/internal/step/internal/constant"
	"github.com/hetue/git/internal/internal/step/internal/get"
)

var _ boot.Step = (*Pull)(nil)

type Pull struct {
	git        *command.Git
	repository *config.Repository
	project    *config.Project
	credential *config.Credential
	pull       *config.Pull
}

func newPull(get get.Pull) *Pull {
	return &Pull{
		git:        get.Git,
		repository: get.Repository,
		project:    get.Project,
		credential: get.Credential,
		pull:       get.Pull,
	}
}

func (*Pull) Name() string {
	return "拉取"
}

func (p *Pull) Runnable() bool {
	return !p.project.Pushable()
}

func (p *Pull) Retryable() bool { // 不重试
	return false
}

func (p *Pull) Asyncable() bool { // 不异步
	return false
}

func (p *Pull) Run(ctx *context.Context) (err error) {
	check := gfx.Exists().Reset().Directory(p.project.Directory)
	if _, exists := check.Build().Check(); !exists { // 检查工作目录是否存在
		err = exception.New().Message("目录不存在").Field(field.New("filepath", p.project.Directory)).Build()
	} else if cle := p.clone(ctx); nil != cle { // 克隆项目
		err = cle
	} else if che := p.checkout(ctx); nil != che { // 检出提交的代码
		err = che
	} else { // 处理子模块因为各种原因无法下载的情况
		err = p.update(ctx)
	}

	return
}

func (p *Pull) clone(ctx *context.Context) (err error) {
	arguments := args.New().Build().Subcommand("clone", p.url())
	if p.pull.Submodules {
		arguments.Flag("remote-submodules").Flag("recurse-submodules")
	}
	if 0 != p.pull.Depth { // nolint:staticcheck
		arguments.Argument("depth", p.pull.Depth)
	}
	// 防止证书错误
	arguments.Flag("config").Add("http.sslVerify=false")
	arguments.Add(p.project.Directory)
	if ee := p.git.Exec(ctx, arguments.Build()); nil != ee {
		// err = p.again(ctx, arguments.Build())
		err = ee
	}

	return
}

func (p *Pull) checkout(ctx *context.Context) (err error) {
	checkout := strings.TrimSpace(p.repository.Checkout)
	if "" == checkout { // nolint:staticcheck
		return
	}

	arguments := args.New().Build().Subcommand("checkout").Add(checkout)
	err = p.git.Exec(ctx, arguments.Build())

	return
}

func (p *Pull) update(ctx *context.Context) (err error) {
	check := gfx.Exists().Dir(filepath.Join(p.project.Directory, constant.GitSubmodulesFilename))
	if _, exists := check.Build().Check(); !exists && p.pull.Submodules { // 是否有子模块配置文件
		return
	}

	arguments := args.New().Build().Subcommand("submodule", "update").Flag("init", "recursive", "remote")
	err = p.git.Exec(ctx, arguments.Build())

	return
}

func (p *Pull) url() (url string) {
	if constant.Pull == p.project.Mode && "" != p.credential.Key { // nolint:staticcheck
		url = os.Getenv(constant.DroneSSHUrl)
	} else {
		url = p.repository.Url
	}

	return
}
