package config

type Source struct {
	// 本地上传目录
	Folder string `default:"${FOLDER=.}" json:"folder,omitempty"`
	// 是否清空存储桶
	Clear *bool `default:"${CLEAR=true}" json:"clear,omitempty"`
	// 路径前缀，所有文件上传都会在这上面加上前缀
	Prefix string `default:"${PREFIX}" json:"prefix,omitempty"`
	// 路径后缀，所有文件上传都会在这上面加上后缀
	Suffix string `default:"${SUFFIX}" json:"suffix,omitempty"`
}

func newSource(s3 *S3) *Source {
	return s3.Source
}
