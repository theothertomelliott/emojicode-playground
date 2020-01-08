package run

import (
	"context"
	"os/exec"
)

type BuildExec interface {
	Build(ctx context.Context, sourcePath string) (*exec.Cmd, error)
	Run(ctx context.Context, sourcePath string) (*exec.Cmd, error)
}
