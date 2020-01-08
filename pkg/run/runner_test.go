package run_test

import (
	"context"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/theothertomelliott/emojicode-playground/pkg/run"
	"github.com/theothertomelliott/emojicode-playground/pkg/run/dockerexec"
)

func TestRunHello(t *testing.T) {
	r := run.New(dockerexec.New(), "./testdata", 10*time.Second)

	pipeR, pipeW, _ := os.Pipe()

	err := r.Run(context.Background(), []byte(helloCode), pipeW)
	if err != nil {
		t.Errorf("could not run code: %v", err)
	}
	pipeW.Close()

	data, err := ioutil.ReadAll(pipeR)
	if err != nil {
		t.Errorf("could not read data: %v", err)
	}
	pipeR.Close()

	if string(data) != `Hello world!
Привет мир!
你好，世界！
` {
		t.Errorf("Output not as expected: %s", string(data))
	}
}

func TestBuildFailure(t *testing.T) {
	r := run.New(dockerexec.New(), "./testdata", 10*time.Second)

	pipeR, pipeW, _ := os.Pipe()

	err := r.Run(context.Background(), []byte(helloBadSyntax), pipeW)
	if err == nil {
		t.Errorf("expected an error from running code")
	}
	pipeW.Close()

	data, err := ioutil.ReadAll(pipeR)
	if err != nil {
		t.Errorf("could not read data: %v", err)
	}
	pipeR.Close()

	// Expect that the output includes the problem line
	if !strings.Contains(string(data), "code.emojic:3:1") {
		t.Errorf("Output not as expected:\n%s", string(data))
	}
}

func TestRunTimeout(t *testing.T) {
	r := run.New(dockerexec.New(), "./testdata", 5*time.Second)

	pipeR, pipeW, _ := os.Pipe()

	err := r.Run(context.Background(), []byte(infiniteLoop), pipeW)
	if err == nil {
		t.Errorf("expected an error from running code")
	}
	pipeW.Close()

	data, err := ioutil.ReadAll(pipeR)
	if err != nil {
		t.Errorf("could not read data: %v", err)
	}
	pipeR.Close()

	// Expect that the output includes the problem line
	if len(string(data)) > 0 {
		t.Errorf("Expected no output, got:\n%s", string(data))
	}
}

const helloCode = `🏁 🍇
  😀🔤Hello world!🔤❗️
  😀🔤Привет мир!🔤❗️
  😀🔤你好，世界！🔤❗️
🍉
`

const helloBadSyntax = `🏁 🍇
  😀🔤Hello world!🔤
🍉
`

const infiniteLoop = `🏁 🍇
  🔁 👍 🍇
    😀 🔤It goes on and on and on🔤❗️
  🍉
🍉`
