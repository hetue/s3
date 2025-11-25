package config

type Server struct {
	// 存储桶地址
	Endpoint string `default:"${ENDPOINT}" validate:"required,url" json:"endpoint,omitempty"`
	// 区域
	Region string `default:"${REGIN=ap-chengdu}" json:"region,omitempty"`
	// 桶
	Bucket string `default:"${BUCKET}" json:"bucket,omitempty"`
	// 分隔符
	Separator string `default:"${SEPARATOR=/}" json:"separator,omitempty"`
}

func newServer(s3 *S3) *Server {
	return s3.Server
}
