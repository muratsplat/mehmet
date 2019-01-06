package php

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Server struct {
	Protocol   string `json:"SERVER_PROTOCOL"`
	Method     string `json:"REQUEST_METHOD"`
	URI        string `json:"REQUEST_URI"`
	Query      string `json:"QUERY_STRING"`
	Server     string `json:"SERVER_SOFTWARE"`
	ServerName string `json:"SERVER_NAME"`
	Host       string `json:"HTTP_HOST"`
	RemoteAddr string `json:"REMOTE_ADDR"`
	RemotePort string `json:"REMOTE_PORT"`
}

type Inject struct {
	Server    string
	GlobalGET string
	PATH      string
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
	resolveRemoteAddr := ResolveRemoteAddr(p.Request.RemoteAddr)
	// https://secure.php.net/manual/tr/reserved.variables.server.php
	s := &Server{
		RemoteAddr: resolveRemoteAddr.Ip,
		RemotePort: resolveRemoteAddr.Port,
		ServerName: p.Request.URL.Hostname(),
		Host:       p.Request.Host,
		Protocol:   p.Request.Proto,
		Method:     p.Request.Method,
		URI:        fmt.Sprintf("%s?%s", p.Request.URL.Path, p.Request.URL.RawQuery),
		Query:      p.Request.URL.RawQuery,
		Server:     "mehmet",
	}

	globalGet := p.Request.URL.RawQuery
	serverInject, err := json.Marshal(s)
	if err != nil {
		return "", err
	}

	var inject Inject
	inject.Server = string(serverInject)
	inject.GlobalGET = globalGet
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
$_MEHMET_SERVER_INJECT = <<<JSON
{{.Server}}
JSON;
$_MEHMET_SERVER = json_decode($_MEHMET_SERVER_INJECT, true);
unset($_MEHMET_SERVER_INJECT);
$_SERVER = array_merge($_SERVER, $_MEHMET_SERVER);
unset($_MEHMET_SERVER);
parse_str("{{.GlobalGET}}", $_MEHMET_GET_INJECT);
$_GET = array_merge($_GET, $_MEHMET_GET_INJECT);
unset($_MEHMET_GET_INJECT);
require_once("{{.PATH}}");`
