package config

type Secret struct {
	// 授权，类型于用户名
	Ak string `default:"${PLUGIN_AK}" validate:"required" json:"ak,omitempty"`
	// 授权，类型于密码
	Sk string `default:"${PLUGIN_SK}" validate:"required" json:"sk,omitempty"`
	// 会话
	Session string `default:"${PLUGIN_SESSION}" json:"session,omitempty"`
}

func newSecret(s3 *S3) *Secret {
	return s3.Secret
}
