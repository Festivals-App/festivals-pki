# Makefile for festivals-pki

VERSION=development
DATE=$(shell date +"%d-%m-%Y-%H-%M")
REF=refs/tags/development
export

build:
	go build

test:
	 go run festivalspki-test main.go

update:
	systemctl stop festivals-gateway
	cp festivals-gateway /usr/local/bin/festivals-gateway
	systemctl start festivals-gateway

uninstall:
	systemctl stop festivals-gateway
	rm /usr/local/bin/festivals-gateway
	rm /etc/festivals-gateway.conf
	rm /etc/systemd/system/festivals-gateway.service

run:
	./festivals-gateway

stop:
	killall festivals-gateway

clean:
	rm -r festivals-gateway