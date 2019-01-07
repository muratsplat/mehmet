package php

import (
	"net/http/httptest"
	"testing"
)

func TestFirst(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	eval := NewPHPInject(req, "/app/hello.php")

	const expect = `
$_MEHMET_SERVER_INJECT = <<<JSON
{"SERVER_PROTOCOL":"HTTP/1.1","REQUEST_METHOD":"GET","REQUEST_URI":"/foo?","QUERY_STRING":"","SERVER_SOFTWARE":"mehmet","SERVER_NAME":"example.com","HTTP_HOST":"example.com","REMOTE_ADDR":"192.0.2.1","REMOTE_PORT":"1234"}
JSON;
$_MEHMET_SERVER = json_decode($_MEHMET_SERVER_INJECT, true);
unset($_MEHMET_SERVER_INJECT);
$_SERVER = array_merge($_SERVER, $_MEHMET_SERVER);
unset($_MEHMET_SERVER);
parse_str("", $_MEHMET_GET_INJECT);
$_GET = array_merge($_GET, $_MEHMET_GET_INJECT);
unset($_MEHMET_GET_INJECT);
require_once("/app/hello.php");`
	script, err := eval.Eval()

	if err != nil {
		t.Fatal(err)
	}
	if script != expect {
		t.Fatalf("Expect: \r\t%v\n Got:\r\t%v", expect, script)
	}
}

func TestQuery(t *testing.T) {
	req := httptest.NewRequest("GET", "http://test.com/foo?bar=foo", nil)
	eval := NewPHPInject(req, "/app/hello.php")

	const expect = `
$_MEHMET_SERVER_INJECT = <<<JSON
{"SERVER_PROTOCOL":"HTTP/1.1","REQUEST_METHOD":"GET","REQUEST_URI":"/foo?bar=foo","QUERY_STRING":"bar=foo","SERVER_SOFTWARE":"mehmet","SERVER_NAME":"test.com","HTTP_HOST":"test.com","REMOTE_ADDR":"192.0.2.1","REMOTE_PORT":"1234"}
JSON;
$_MEHMET_SERVER = json_decode($_MEHMET_SERVER_INJECT, true);
unset($_MEHMET_SERVER_INJECT);
$_SERVER = array_merge($_SERVER, $_MEHMET_SERVER);
unset($_MEHMET_SERVER);
parse_str("bar=foo", $_MEHMET_GET_INJECT);
$_GET = array_merge($_GET, $_MEHMET_GET_INJECT);
unset($_MEHMET_GET_INJECT);
require_once("/app/hello.php");`
	script, err := eval.Eval()

	if err != nil {
		t.Fatal(err)
	}
	if script != expect {
		t.Fatalf("Expect: \r\t%v\n Got:\r\t%v", expect, script)
	}
}
