package command

import (
	"context"

	"github.com/goexl/args"
	"github.com/hetue/boot"
	"github.com/hetue/git/internal/internal/config"
	"github.com/hetue/git/internal/internal/step/internal/constant"
)

type Git struct {
	command *boot.Command
	binary  *config.Binary
	project *config.Project
}

func NewGit(base *boot.Command, binary *config.Binary, project *config.Project) *Git {
	return &Git{
		command: base,
		binary:  binary,
		project: project,
	}
}

func (g *Git) Exec(ctx *context.Context, arguments *args.Arguments) (err error) {
	command := g.command.New(g.binary.Git).Arguments(arguments).Dir(g.project.Directory)
	environment := command.Environment()
	environment.String(constant.SpeedLimit)
	environment.String(constant.SpeedTime)
	command = environment.Build()
	_, err = command.Context(*ctx).Build().Exec()

	return
}
