package run_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/theothertomelliott/emojicode-playground/pkg/run"
	"github.com/theothertomelliott/emojicode-playground/pkg/run/dockerexec"
)

func TestRunHello(t *testing.T) {
	r := run.New(dockerexec.New(), "./testdata")

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
	fmt.Printf("Data: %s\n", string(data))
	pipeR.Close()

	if string(data) != `Hello world!
Привет мир!
你好，世界！
` {
		t.Errorf("Output not as expected: %s", string(data))
	}
}

const helloCode = `
🏁 🍇
  😀🔤Hello world!🔤❗️
  😀🔤Привет мир!🔤❗️
  😀🔤你好，世界！🔤❗️
🍉
`
