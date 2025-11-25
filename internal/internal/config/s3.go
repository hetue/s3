package config

import (
	"github.com/harluo/config"
)

type S3 struct {
	*Server `default:"{}" json:"server,omitempty"`
	*Secret `default:"{}" json:"secret,omitempty"`
	*Source `default:"{}" json:"source,omitempty"`
}

func newS3(getter config.Getter) (s3 *S3, err error) {
	s3 = new(S3)
	err = getter.Get(s3)

	return
}
