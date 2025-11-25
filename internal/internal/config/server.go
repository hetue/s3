package config

type Server struct {
	// 存储桶地址
	Endpoint string `default:"${PLUGIN_ENDPOINT}" validate:"required,url" json:"endpoint,omitempty"`
	// 区域
	Region string `default:"${PLUGIN_REGION}" json:"region,omitempty"`
	// 桶
	Bucket string `default:"${PLUGIN_BUCKET}" json:"bucket,omitempty"`
	// 分隔符
	Separator string `default:"${PLUGIN_SEPARATOR}" json:"separator,omitempty"`
}

func newServer(s3 *S3) *Server {
	return s3.Server
}
