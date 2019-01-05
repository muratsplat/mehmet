package php

import (
	"net/http/httptest"
	"testing"
)

func TestFirst(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	eval := NewPHPInject(req, "/app/hello.php")

	const expect = `
$_MEHMET_INJECT = <<<JSON
{"SERVER_PROTOCOL":"HTTP/1.1","REQUEST_METHOD":"GET","REQUEST_URI":"http://example.com/foo"}
JSON;
$_MEHMET_SERVER = json_decode($_MEHMET_INJECT, true);
unset($_MEHMET_INJECT);
$_SERVER = array_merge($_SERVER, $_MEHMET_SERVER);
unset($_MEHMET_SERVER);
require_once("/app/hello.php");`
	script, err := eval.Eval()

	if err != nil {
		t.Fatal(err)
	}
	if script != expect {
		t.Fatalf("Expect: \r\t%v\n Got:\r\t%v", expect, script)
	}
}
