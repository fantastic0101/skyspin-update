
BUILD_TIME=`date '+%Y-%m-%d %H:%M:%S %z'`
GO_VERSION=`go version`
prefix='duck/lazy.'
LDFLAGS="-X '${prefix}BUILD_TIME=${BUILD_TIME}' -X '${prefix}GO_VERSION=${GO_VERSION}' -X '${prefix}AUTHOR=${AUTHOR}' -s -w"

source header.sh

for service in $services; do
    echo build $service
    
    if go build -ldflags "${LDFLAGS}" -o bin/$service ./service/$service; then
        echo "编译成功 $ret 1"
    else
        echo "编译失败 $ret 1"
    fi
done