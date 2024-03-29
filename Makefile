SHELL := /bin/bash
BASEDIR = $(shell pwd)
GOPROXY=https://mirrors.cloud.tencent.com/go/

# build with version infos
versionDir = "minepin/com/v"
gitTag = $(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
buildDate = $(shell TZ=Asia/Shanghai date +%FT%T%z)
gitCommit = $(shell git log --pretty=format:'%H' -n 1)
gitTreeState = $(shell if git status|grep -q -E 'clean|干净';then echo clean; else echo dirty; fi)

ldflags=" -linkmode external -extldflags '-static' -s -w -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} -X ${versionDir}.gitCommit=${gitCommit} -X ${versionDir}.gitTreeState=${gitTreeState}"

all: gotool
	@go build -a -v -ldflags ${ldflags} -o app .
clean:
	rm -f minegin
	find . -name "[._]*.s[a-w][a-z]" | xargs -i rm -f {}
gotool:
	gofmt -w .
	go vet . | grep -v vendor;true
ca:
	openssl req -new -nodes -x509 -out conf/server.crt -keyout conf/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=127.0.0.1/emailAddress=xxxxx@qq.com"
air:
	if [ -e ./config.yaml ] ; then cp config.yaml ./tmp/; fi
	air

help:
	@echo "make - compile the source code"
	@echo "make clean - remove binary file and vim swp files"
	@echo "make gotool - run go tool 'fmt' and 'vet'"
	@echo "make ca - generate ca files"

.PHONY: clean gotool ca help air
