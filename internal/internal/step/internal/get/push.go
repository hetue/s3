package get

import (
	"github.com/goexl/log"
	"github.com/harluo/di"
	"github.com/hetue/git/internal/internal/config"
	"github.com/hetue/git/internal/internal/step/internal/command"
)

type Push struct {
	di.Get

	Repository *config.Repository
	Project    *config.Project
	Credential *config.Credential
	Push       *config.Push

	Git    *command.Git
	Logger log.Logger
}
