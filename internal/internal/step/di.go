package step

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Instance().Put(
		newCredential,
		newSSH,
		newPull,
		newPush,
	).Build().Apply()
}
