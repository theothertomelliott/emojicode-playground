package run

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func New(e BuildExec, workingDir string) *runner {
	return &runner{
		buildExec:  e,
		workingDir: workingDir,
	}
}

type runner struct {
	buildExec  BuildExec
	workingDir string
}

func (r *runner) Run(ctx context.Context, code []byte, output io.Writer) error {
	dir := r.workingDir

	sourcePath := filepath.Join(dir, "code.emojic")

	// Save code to file
	err := ioutil.WriteFile(sourcePath, code, os.ModePerm)
	if err != nil {
		return err
	}

	// Build
	buildCmd, err := r.buildExec.Build(sourcePath)
	if err != nil {
		return err
	}
	out, err := buildCmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Fprint(output, string(out))

	// Run
	runCmd, err := r.buildExec.Run(sourcePath)
	if err != nil {
		return err
	}
	out, err = runCmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Fprint(output, string(out))
	return nil
}
