#
# Makefile
# hzsunshx, 2015-02-11 13:17
#

dev: clean
	npm start

prod: bundle binary

bundle:
	webpack -p --progress --config config/webpack.config.prod.js

binary:
	go get github.com/jteeuwen/go-bindata/...
	(go-bindata -o ../backend/modules/bindata.go -pkg modules bundle/... public/... templates/...)
	(cd ../backend; go get -tags "bindata")
	(cd ../backend; go build -tags "bindata")

install-deps:
	sudo apt-get update -qq
	sudo apt-get install -qq nodejs npm

deps:
	npm install

cross-build:
	GOOS=windows GOARCH=386 go build
	GOOS=linux GOARCH=386 go build -o fileserv-linux-386
	GOOS=linux GOARCH=amd64 go build -o fileserv-linux-amd64

webpack:
	webpack

clean:
	rm -f bundle/bundle.js
# vim:ft=make
#
