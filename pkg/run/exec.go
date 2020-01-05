package run

import "os/exec"

type BuildExec interface {
	Build(sourcePath string) (*exec.Cmd, error)
	Run(sourcePath string) (*exec.Cmd, error)
}
