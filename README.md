[![Build Status](https://travis-ci.com/muratsplat/mehmet.svg?branch=master)](https://travis-ci.com/muratsplat/mehmet)
# Mehmet 
This is simple http server to serve your PHP code to the World ! Dont waste your time to server your PHP application by doing a lot of configuration. 

This app can be used instead of `Nginx`+`PHP-FPM` or `Apache` + `PHP-Mod`.

More details will be added....

# Requiretments
* PHP in  the `PATH` on your OS
 

## How to run
```sh
$ ./mehmet -path ~/code/personal/foo/public/index.php
2019/01/06 00:13:14 Server is starting.
2019/01/06 00:13:14 :8090 listening...
2019/01/06 00:13:21 PHP execution time: 228.824778ms
2019/01/06 00:13:22 PHP execution time: 92.766119ms
2019/01/06 00:13:24 PHP execution time: 164.773723ms
2019/01/06 00:13:25 PHP execution time: 151.359473ms
2019/01/06 00:13:25 PHP execution time: 101.984133ms
```

## TODO
 * ~~Handling URL parameters~~
 * Handling POST request
 * Serving static assets
