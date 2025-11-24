package get

import (
	"github.com/harluo/di"
	"github.com/hetue/git/internal/internal/config"
	"github.com/hetue/git/internal/internal/step/internal/command"
)

type Pull struct {
	di.Get

	Git        *command.Git
	Repository *config.Repository
	Project    *config.Project
	Credential *config.Credential
	Pull       *config.Pull
}
