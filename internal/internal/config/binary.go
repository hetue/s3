package config

type Binary struct {
	// 控制程序
	Git string `default:"/usr/bin/git" json:"git,omitempty"`
}

func newBinary(git *Git) *Binary {
	return git.Binary
}
