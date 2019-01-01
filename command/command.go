package command

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os/exec"

	"github.com/muratsplat/mehmet/php"
)

type Worker interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// PHP worker
type PHP struct {
	Path string
}

func NewPHP(path string) Worker {
	return &PHP{
		Path: path,
	}
}

func (p *PHP) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	cmd := exec.Command("php", "-r", php.EvalCode)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	byt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	go func() {
		defer stdin.Close()
		_, err := stdin.Write(byt)
		if err != nil {
			panic(err)
		}
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	reader := bytes.NewBuffer(out)

	_, err = io.Copy(w, reader)

	if err != nil {
		panic(err)
	}
}
