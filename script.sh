#! /bin/sh

nohup sh -c mongod --dbpath /test > mongoLog &

export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

go run main.go > serverLog

echo "\n========================================"
echo "\nMongo Log\n"
cat mongoLog
echo "=========================================="

echo "\nserver Log\n"
cat serverLog
echo "=========================================="