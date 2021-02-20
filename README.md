# Truss ![Build Status](https://github.com/ben1009/truss/workflows/Go/badge.svg?branch=master)

Truss handles the painful parts of services, freeing you to focus on the
business logic.

![Everything all the time forever](http://i.imgur.com/FtvVeBG.jpg)

## Install

Currently, there is no binary distribution of Truss, you must install from
source.

To install this software, you must:

1. Install protoc 3 or newer. The easiest way is to
download a release from [github](https://github.com/gogo/protobuf/releases)
and add to `$PATH`.
Otherwise [install from source](https://github.com/gogo/protobuf)
1. Install Truss with

	```
	go get -u -d github.com/ben1009/truss
	cd $GOPATH/src/github.com/ben1009/truss
	make dependencies
	make
	```
	On Windows, do the following instead:
	```
	go get -u -d github.com/ben1009/truss
	cd %GOPATH%/src/github.com/ben1009/truss
	wininstall.bat
	```

## Usage

Using Truss is easy. You define your service with [gRPC](http://www.grpc.io/)
and [protoc buffers](https://developers.google.com/protocol-buffers/docs/proto3),
and Truss uses that definition to create an entire service. You can even
add [http annotations](
https://github.com/googleapis/googleapis/blob/928a151b2f871b4239b7707e1bb59258df3fe10a/google/api/http.proto#L36)
for HTTP 1.1/JSON transport!

Then you open the `handlers/handlers.go`,
add you business logic, and you're good to go.

Try Truss for yourself to see the service that is generated:
```
for creating new service
truss -c --servicename serviceName 
```

See [USAGE.md](./USAGE.md) and [TUTORIAL.md](./TUTORIAL.md) for more details.

## Developing

See [DEVELOPING.md](./DEVELOPING.md) for details.
