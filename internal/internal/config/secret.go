package config

type Secret struct {
	// 授权，类型于用户名
	Ak string `validate:"required" json:"ak,omitempty"`
	// 授权，类型于密码
	Sk string `validate:"required" json:"sk,omitempty"`
	// 会话
	Session string `json:"session,omitempty"`
}

func newSecret(s3 *S3) *Secret {
	return s3.Secret
}
