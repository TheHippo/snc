[![Build Status](https://travis-ci.org/TheHippo/snc.svg?branch=master)](https://travis-ci.org/TheHippo/snc)

# snc

Like `netcat`, but with TLS.

**Is it secure?**

No way!

## Usage

Client:
```
snc host port
```

Server options:
```
Usage of snc:
  -b, --bind string
        ip adress to bind (default "0.0.0.0")
  -l, --listen value
        port to listen (default 8888)
  -v, --version
        display snc version
```


## Installation

```
go get -u github.com/TheHippo/snc/cmd/snc
```

## Speed

```
$ dd if=/dev/urandom of=file.txt bs=1048576 count=1000

$ pv file.txt |  ./snc -l 9999
1e+03MB 0:00:12 [  83MB/s] [========================================================================>] 100%            

$ pv file.txt |  nc -l 9999
1e+03MB 0:00:03 [ 256MB/s] [========================================================================>] 100%  
