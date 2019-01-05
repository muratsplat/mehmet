package php

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"text/template"
)

type Server struct {
	Protocol string `json:"SERVER_PROTOCOL"`
	Method   string `json:"REQUEST_METHOD"`
	URI      string `json:"REQUEST_URI"`
}

type Inject struct {
	JSON string
	PATH string
}

type PHPInject struct {
	Request       *http.Request
	PHPScriptPath string
}

func NewPHPInject(r *http.Request, scriptPath string) *PHPInject {
	return &PHPInject{
		Request:       r,
		PHPScriptPath: scriptPath,
	}
}

func (p *PHPInject) String() string {
	eval, err := p.Eval()
	if err != nil {
		log.Printf("Fatal Error: %v", err)
		panic(err)
	}
	return eval

}

func (p *PHPInject) Eval() (string, error) {
	s := &Server{
		Protocol: p.Request.Proto,
		Method:   p.Request.Method,
		URI:      p.Request.RequestURI,
	}

	jsonInject, err := json.Marshal(s)
	if err != nil {
		return "", err
	}

	var inject Inject
	inject.JSON = string(jsonInject)
	inject.PATH = p.PHPScriptPath

	evil, err := GetEvil(inject)
	if err != nil {
		return "", err
	}
	return evil, nil
}

func GetEvil(i Inject) (string, error) {
	t, err := template.New("test").Parse(evalCode)
	if err != nil {
		panic(err)
	}

	buffer := new(bytes.Buffer)
	err = t.Execute(buffer, i)
	if err != nil {
		return buffer.String(), err
	}
	return buffer.String(), nil
}

// https://secure.php.net/manual/en/reserved.variables.server.php
const evalCode = `
$_MEHMET_INJECT = <<<JSON
{{.JSON}}
JSON;
$_MEHMET_SERVER = json_decode($_MEHMET_INJECT, true);
unset($_MEHMET_INJECT);
$_SERVER = array_merge($_SERVER, $_MEHMET_SERVER);
unset($_MEHMET_SERVER);
require_once("{{.PATH}}");`
