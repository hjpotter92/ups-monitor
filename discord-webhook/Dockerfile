FROM	golang:1.19.3-alpine3.16	as	build

WORKDIR	/src

ADD	go.mod	go.sum	./

RUN	go mod download

ADD	.	.

RUN	go build -o /main src/main.go

FROM	alpine:3.16.3

LABEL maintainer="hjpotter92"

COPY --from=build	/main	/main

ENTRYPOINT	["/main"]
