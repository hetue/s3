package internal

import (
	"github.com/hetue/boot"
	"github.com/hetue/git/internal/internal"
)

func New(params internal.Steps) []boot.Step {
	return []boot.Step{
		params.Credential,
		params.SSH,
		params.Pull,
		params.Push,
	}
}
