#DB2 Library
"github.com/asifjalil/cli"

#How Available _clidriver
export DB2HOME=/home/rakhil/dsdriver
export CGO_CFLAGS=-I$DB2HOME/include
export CGO_LDFLAGS=-L$DB2HOME/lib 
  #Linux:
  export LD_LIBRARY_PATH=/home/rakhil/dsdriver/lib

#Where _clidriver
/home/khunpleum/go/src/github.com/tammarut/go-sql/_clidriver/linux

#Run Shell script
    source _shllscript/init.sh
    sh _shllscript/init.sh
    echo $DB2HOME

**DB2 is complicated than other SQL databases.**
**Because golang hasn't driver so you need to set driver.**
