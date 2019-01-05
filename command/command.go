package command

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"time"

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
	now := time.Now()
	eval := php.NewPHPInject(r, p.Path)
	cmd := exec.Command("php", "-r", eval.String())
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
		log.Printf("PHP error reason (%v)", string(out))
		log.Printf("Execute Error:(%v)", err)
		panic(err)
	}

	reader := bytes.NewBuffer(out)
	_, err = io.Copy(w, reader)
	log.Printf("PHP execution time: %s", time.Now().Sub(now).String())
	if err != nil {
		panic(err)
	}
}
