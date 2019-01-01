package php

const EvalCode = `
// https://secure.php.net/manual/en/reserved.variables.server.php
$_SERVER['SERVER_PROTOCOL']= 'HTTP/1.1';
$_SERVER['REQUEST_METHOD'] = 'GET';
$_SERVER['REQUEST_URI'] = '/foo';
require_once('/Users/muod/code/personal/foo/public/index.php');`
