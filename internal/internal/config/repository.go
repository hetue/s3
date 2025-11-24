package config

import (
	"net/url"
)

type Repository struct {
	// 远程仓库地址
	Url string `default:"${GIT_HTTP_URL}" validate:"required,url" json:"url,omitempty"`
	// 签出代码
	Checkout string `default:"${COMMIT}" json:"checkout,omitempty"`
}

func newRepository(git *Git) *Repository {
	return git.Repository
}

func (r *Repository) Host() (host string, err error) {
	if parsed, pe := url.Parse(r.Url); nil != pe {
		err = pe
	} else {
		host = parsed.Host
	}

	return
}
