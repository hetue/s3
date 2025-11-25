package step

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Instance().Put(
		newUpload,
	).Build().Apply()
}
