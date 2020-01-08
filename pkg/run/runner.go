package run

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

func New(e BuildExec, workingDir string, timeout time.Duration) *runner {
	return &runner{
		buildExec:  e,
		workingDir: workingDir,
		timeout:    timeout,
	}
}

type runner struct {
	buildExec  BuildExec
	workingDir string
	timeout    time.Duration
}

func (r *runner) Run(ctx context.Context, code []byte, output io.Writer) (err error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer func() {
		if ctx.Err() != nil {
			err = ctx.Err()
			return
		}
		cancel()
	}()

	dir := filepath.Join(r.workingDir, uuid.New().String())
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)

	sourcePath := filepath.Join(dir, "code.emojic")

	// Save code to file
	err = ioutil.WriteFile(sourcePath, code, os.ModePerm)
	if err != nil {
		return err
	}

	// Build
	buildCmd, err := r.buildExec.Build(ctx, sourcePath)
	if err != nil {
		return err
	}
	out, err := buildCmd.CombinedOutput()
	fmt.Fprint(output, string(out))
	if err != nil {
		return err
	}

	// Run
	runCmd, err := r.buildExec.Run(ctx, sourcePath)
	if err != nil {
		return err
	}
	out, err = runCmd.CombinedOutput()
	fmt.Fprint(output, string(out))
	return err
}
