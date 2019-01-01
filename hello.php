<?php
declare(strict_types=1);

$stdin = fopen('php://stdin', 'r');
$content = stream_get_contents($stdin);

phpinfo();

printf("%s", $phpinfo); 
