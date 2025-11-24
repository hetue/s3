package config

type Credential struct {
	// 用户名
	Username string `default:"${NETRC_USERNAME=${GIT_USERNAME}}" json:"username,omitempty"`
	// 密码
	Password string `default:"${NETRC_PASSWORD=${GIT_PASSWORD}}" json:"password,omitempty"`
	// 密钥
	Key string `json:"key,omitempty"`
}

func newCredential(git *Git) *Credential {
	return git.Credential
}
