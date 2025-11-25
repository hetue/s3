package main

import (
	"github.com/hetue/boot"
	"github.com/hetue/s3/internal"
)

func main() {
	bootstrap := boot.New()
	bootstrap.Build().Boot(internal.New)
}
