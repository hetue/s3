package config

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Instance().Put(
		newGit,

		newBinary,
		newCredential,
		newProject,
		newPull,
		newPush,
		newRepository,
	).Build().Apply()
}
