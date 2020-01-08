package dockerexec

import (
	"context"
	"fmt"
	"os"
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
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return exec.CommandContext(ctx,
		"docker",
		"run",
		"--rm",
		"--volume", fmt.Sprintf("%s:/working", wd),
		"-w", "/working",
		"theothertomelliott/emojicode:0.8.4",
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
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	binaryPath, err := e.binaryForSource(sourcePath)
	if err != nil {
		return nil, err
	}

	return exec.CommandContext(ctx,
		"nice",
		"docker",
		"run",
		"--rm",
		"--volume", fmt.Sprintf("%s:/working", wd),
		"-w", "/working",
		"theothertomelliott/emojicode:0.8.4",
		binaryPath,
	), nil
}
