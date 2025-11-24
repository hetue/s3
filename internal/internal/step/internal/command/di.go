package command

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Instance().Put(
		NewGit,
	).Build().Apply()
}
