package config

import (
	"github.com/harluo/config"
)

type Git struct {
	*Credential `default:"{}" json:"credential,omitempty"`
	*Project    `default:"{}" json:"project,omitempty"`
	*Pull       `default:"{}" json:"pull,omitempty"`
	*Push       `default:"{}" json:"push,omitempty"`
	*Repository `default:"{}" json:"repository,omitempty"`

	Binary *Binary `default:"{}" json:"binary,omitempty"`
}

func newGit(getter config.Getter) (git *Git, err error) {
	git = new(Git)
	err = getter.Get(git)

	return
}
