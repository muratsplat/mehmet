package command

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

type Worker interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// PHP worker
type PHP struct {
	// Options
}

func NewPHP() Worker {
	return &PHP{}
}

func (p *PHP) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	cmd := exec.Command("php", "hello.php")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	byt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	log.Println(string(byt))
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

	log.Printf("Out %s", out)
	reader := bytes.NewBuffer(out)

	_, err = io.Copy(w, reader)

	if err != nil {
		panic(err)
	}
}
