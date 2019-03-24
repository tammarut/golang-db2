#Set Path DB2driver
#!/bin/sh
export DB2HOME=/home/khunpleum/go/src/github.com/tammarut/go-sql/_clidriver/linux
export CGO_CFLAGS=-I$DB2HOME/include
export CGO_LDFLAGS=-L$DB2HOME/lib
export LD_LIBRARY_PATH=$DB2HOME/lib