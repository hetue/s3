package config

type Pull struct {
	// 子模块
	Submodules bool `default:"true" json:"submodule,omitempty"`
	// 深度
	Depth int `json:"depth,omitempty"`
}

func newPull(git *Git) *Pull {
	return git.Pull
}
