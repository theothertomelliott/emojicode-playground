package dockerexec

import "testing"

func TestBuild(t *testing.T) {
	e := &buildExec{}
	buildCmd, err := e.Build("testdata/hello/hello.emojic")
	if err != nil {
		t.Fatal(err)
	}
	out, err := buildCmd.CombinedOutput()
	if err != nil {
		t.Error(err)
	}
	t.Log(string(out))
}

func TestBuildAndRun(t *testing.T) {
	e := &buildExec{}
	buildCmd, err := e.Build("testdata/hello/hello.emojic")
	if err != nil {
		t.Fatal(err)
	}
	out, err := buildCmd.CombinedOutput()
	if err != nil {
		t.Error(err)
	}
	t.Log(string(out))
	runCmd, err := e.Run("testdata/hello/hello.emojic")
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
