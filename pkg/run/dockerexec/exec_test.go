package dockerexec

import (
	"context"
	"strings"
	"testing"
	"time"
)

func TestBuild(t *testing.T) {
	e := &buildExec{}
	buildCmd, err := e.Build(context.Background(), "testdata/hello/hello.emojic")
	if err != nil {
		t.Fatal(err)
	}
	out, err := buildCmd.CombinedOutput()
	if err != nil {
		t.Error(err)
	}
	t.Log(string(out))
}

func TestBuildFailure(t *testing.T) {
	e := &buildExec{}
	buildCmd, err := e.Build(context.Background(), "testdata/hellobadsyntax/hello.emojic")
	if err != nil {
		t.Fatal(err)
	}
	out, err := buildCmd.CombinedOutput()
	if err == nil {
		t.Error("expected an error")
	}
	// Expect that the output includes the bad line
	if !strings.Contains(string(out), "testdata/hellobadsyntax/hello.emojic:3:1") {
		t.Errorf("Output not as expected: %q", string(out))
	}
}

func TestExecTimeout(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	e := &buildExec{}
	buildCmd, err := e.Build(ctx, "testdata/loop/main.emojic")
	if err != nil {
		t.Fatal(err)
	}
	out, err := buildCmd.CombinedOutput()
	if err != nil {
		t.Errorf("building failed: %v", err)
	}
	t.Log(string(out))
	runCmd, err := e.Run(ctx, "testdata/loop/main.emojic")
	if err != nil {
		t.Fatal(err)
	}
	got, err := runCmd.CombinedOutput()
	if err == nil {
		t.Error("expected an error")
	}
	if string(got) != "" {
		t.Errorf("Expected no output, got: %q", string(got))
	}
}

func TestBuildAndRun(t *testing.T) {
	e := &buildExec{}
	buildCmd, err := e.Build(context.Background(), "testdata/hello/hello.emojic")
	if err != nil {
		t.Fatal(err)
	}
	out, err := buildCmd.CombinedOutput()
	if err != nil {
		t.Error(err)
	}
	t.Log(string(out))
	runCmd, err := e.Run(context.Background(), "testdata/hello/hello.emojic")
	if err != nil {
		t.Fatal(err)
	}
	got, err := runCmd.CombinedOutput()
	if err != nil {
		t.Error(err)
	}
	if string(got) != "Hello world!\n" {
		t.Errorf("Output not as expected: %q", string(got))
	}
}

func TestGetBinaryPath(t *testing.T) {
	e := &buildExec{}
	got, err := e.binaryForSource("testdata/hello/hello.emojic")
	if err != nil {
		t.Fatal(err)
	}
	if got != "testdata/hello/hello" {
		t.Errorf("filepath not as expected: %v", got)
	}
}
