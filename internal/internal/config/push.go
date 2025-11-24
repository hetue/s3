package config

type Push struct {
	// 标签
	Tag string `json:"tag,omitempty"`
	// 作者
	Author string `default:"${COMMIT_AUTHOR_NAME}" json:"author,omitempty"`
	// 邮箱
	Email string `default:"${COMMIT_AUTHOR_EMAIL}" json:"email,omitempty"`
	// 提交消息
	Message string `default:"${COMMIT_MESSAGE=drone}" json:"message,omitempty"`
	// 是否强制提交
	Force *bool `json:"force,omitempty"`
}

func newPush(git *Git) *Push {
	return git.Push
}
