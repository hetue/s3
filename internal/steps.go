package internal

import (
	"github.com/hetue/boot"
	"github.com/hetue/s3/internal/internal/step"
)

func New(
	upload *step.Upload,
) []boot.Step {
	return []boot.Step{
		upload,
	}
}
