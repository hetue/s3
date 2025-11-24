package internal

import (
	"github.com/harluo/di"
	"github.com/hetue/git/internal/internal/step"
)

type Steps struct {
	di.Get

	Credential *step.Credential
	Pull       *step.Pull
	Push       *step.Push
	SSH        *step.SSH
}
