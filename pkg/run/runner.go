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

func (r *runner) Run(ctx context.Context, code []byte, output io.Writer) error {
	fmt.Printf("Setting timeout: %v\n", r.timeout)
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	dir := filepath.Join(r.workingDir, uuid.New().String())
	err := os.MkdirAll(dir, os.ModePerm)
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
	// TODO: Limit the total amount of output that can be received
	out, err := buildCmd.CombinedOutput()

	// Never return output if the program timed out
	// Attempting to print this can result in locking up
	if ctx.Err() != nil {
		return ctx.Err()
	}
	fmt.Fprint(output, string(out))
	if err != nil {
		return err
	}
	fmt.Println("Built code")

	// Run
	runCmd, err := r.buildExec.Run(ctx, sourcePath)
	if err != nil {
		return err
	}
	// TODO: Limit the total amount of output that can be received
	out, err = runCmd.CombinedOutput()

	// Never return output if the program timed out
	// Attempting to print this can result in locking up
	if ctx.Err() != nil {
		return ctx.Err()
	}
	fmt.Fprint(output, string(out))
	fmt.Println("Program exited")
	return err
}
