# Makefile for festivals-pki

VERSION=development
DATE=$(shell date +"%d-%m-%Y-%H-%M")
REF=refs/tags/development
export

build:
	go build

test:
	 go test