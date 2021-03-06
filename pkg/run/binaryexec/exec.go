package binaryexec

import (
	"context"
	"os/exec"
	"path/filepath"

	"github.com/theothertomelliott/emojicode-playground/pkg/run"
)

var _ run.BuildExec = &buildExec{}

func New() run.BuildExec {
	return &buildExec{}
}

type buildExec struct {
}

func (e *buildExec) Build(ctx context.Context, sourcePath string) (*exec.Cmd, error) {
	return exec.CommandContext(ctx,
		"nice",
		"emojicodec",
		sourcePath,
	), nil
}

func (e *buildExec) binaryForSource(sourcePath string) (string, error) {
	extension := filepath.Ext(sourcePath)
	path := sourcePath[:len(sourcePath)-len(extension)]
	return path, nil
}

func (e *buildExec) Run(ctx context.Context, sourcePath string) (*exec.Cmd, error) {
	binaryPath, err := e.binaryForSource(sourcePath)
	if err != nil {
		return nil, err
	}

	return exec.CommandContext(ctx,
		"nice",
		binaryPath,
	), nil
}
