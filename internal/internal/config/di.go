package config

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Instance().Put(
		newS3,

		newServer,
		newSecret,
		newSource,
	).Build().Apply()
}
