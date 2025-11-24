package config

import (
	"os"
)

type Project struct {
	// 目录
	Directory string `default:"." validate:"required" json:"dir,omitempty"`
	// 模式
	Mode string `default:"push" json:"mode,omitempty"`
	// 是否清理
	Clear *bool `json:"clear,omitempty"`

	executed bool
	pushable bool
}

func newProject(git *Git) *Project {
	return git.Project
}

func (p *Project) Pushable() (pushable bool) {
	if !p.executed {
		p.check()
	}
	pushable = p.pushable

	return
}

func (p *Project) check() {
	if entries, re := os.ReadDir(p.Directory); nil == re {
		p.pushable = 0 != len(entries) // nolint:staticcheck
	} else if os.IsNotExist(re) {
		p.pushable = false
	} else {
		p.pushable = true
	}
	p.executed = true
}
