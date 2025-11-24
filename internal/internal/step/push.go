package step

import (
	"context"
	"path/filepath"

	"github.com/goexl/args"
	"github.com/goexl/gfx"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/gox/rand"
	"github.com/goexl/log"
	"github.com/hetue/boot"
	"github.com/hetue/git/internal/internal/config"
	"github.com/hetue/git/internal/internal/step/internal/command"
	"github.com/hetue/git/internal/internal/step/internal/constant"
	"github.com/hetue/git/internal/internal/step/internal/get"
)

var _ boot.Step = (*Push)(nil)

type Push struct {
	repository *config.Repository
	project    *config.Project
	credential *config.Credential
	push       *config.Push

	git    *command.Git
	logger log.Logger
}

func newPush(push get.Push) *Push {
	return &Push{
		repository: push.Repository,
		project:    push.Project,
		credential: push.Credential,
		push:       push.Push,

		git:    push.Git,
		logger: push.Logger,
	}
}

func (*Push) Name() string {
	return "推送"
}

func (p *Push) Runnable() bool {
	return p.project.Pushable()
}

func (p *Push) Retryable() bool { // 重试
	return true
}

func (p *Push) Asyncable() bool { // 不异步
	return false
}

func (p *Push) Run(ctx *context.Context) (err error) {
	if ie := p.init(ctx); nil != ie { // 初始化
		err = ie
	} else if che := p.checkout(ctx); nil != che { // 签出新代码
		err = che
	} else if name, coe := p.commit(ctx); nil != coe { // 提交代码
		err = coe
	} else if re := p.remote(ctx, name); nil != re { // 添加远程仓库地址
		err = re
	} else if te := p.tag(ctx); nil != te { // 如果有标签，推送标签
		err = te
	} else { // 推送
		err = p.do(ctx, name)
	}

	return
}

func (p *Push) init(ctx *context.Context) (err error) {
	if _, exists := gfx.Exists().Dir(filepath.Join(p.project.Directory, constant.GitHome)).Build().Check(); exists {
		// 不需要初始化仓库
	} else if ie := p.exec(ctx, "init"); nil != ie { // 初始化目录
		err = ie
	} else if dbe := p.exec(ctx, "config", "init.defaultBranch", "master"); nil != dbe { // 设置默认分支
		err = dbe
	} else if cue := p.exec(ctx, "config", "user.name", p.push.Author); nil != cue { // 设置用户名
		err = cue
	} else if cee := p.exec(ctx, "config", "user.email", p.push.Email); nil != cee { // 设置邮箱
		err = cee
	} else if cae := p.exec(ctx, "config", "boot.autocrlf", "false"); nil != cae { // 设置换行符
		err = cae
	}

	return
}

func (p *Push) checkout(ctx *context.Context) (err error) {
	dir := field.New("dir", p.project.Directory)
	p.logger.Debug("是完整的Git仓库，无需初始化和配置", dir)
	p.logger.Debug("签出目标分支开始", dir)
	// 签出目标分支
	err = p.git.Exec(ctx, args.New().Build().Subcommand("checkout").Flag("B").Add(p.repository.Checkout).Build())
	p.logger.Debug("签出目标分支完成", dir)

	return
}

func (p *Push) commit(ctx *context.Context) (name string, err error) {
	dir := field.New("dir", p.project.Directory)
	p.logger.Debug("提交代码开始", dir)
	if ae := p.exec(ctx, "add", "."); nil != ae { // 只添加改变的文件
		err = ae
	} else if me := p.message(ctx, dir); nil != me { // 设置提交消息
		err = me
	} else { // 生成随机远端名字
		name = rand.New().String().Build().Generate()
	}

	return
}

func (p *Push) remote(ctx *context.Context, name string) (err error) {
	arguments := args.New().Build().Subcommand("remote", "add").Add(name, p.repository.Url)
	err = p.git.Exec(ctx, arguments.Build())

	return
}

func (p *Push) tag(ctx *context.Context) (err error) {
	if "" == p.push.Tag { // nolint:staticcheck
		return
	}

	argument := args.New().Build().Subcommand("tag").Flag("annotate").Add(p.push.Tag).Flag("message").Add(p.push.Message)
	err = p.git.Exec(ctx, argument.Build())

	return
}

func (p *Push) message(ctx *context.Context, fields ...gox.Field[any]) (err error) {
	arguments := args.New().Build().Subcommand("commit", ".").Flag("message").Add(p.push.Message)
	if err = p.git.Exec(ctx, arguments.Build()); nil == err {
		p.logger.Debug("提交代码完成", fields...)
	}

	return
}

func (p *Push) do(ctx *context.Context, name string) (err error) {
	argument := args.New().Build().Subcommand("push").Flag("set-upstream").Add(name, p.repository.Checkout).Flag("tags")
	if nil != p.push.Force && *p.push.Force {
		argument.Flag("force")
	}
	err = p.git.Exec(ctx, argument.Build())

	return
}

func (p *Push) exec(ctx *context.Context, subcommand string, subcommands ...string) error {
	return p.git.Exec(ctx, args.New().Build().Subcommand(subcommand, subcommands...).Build())
}
