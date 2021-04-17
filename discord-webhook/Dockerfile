FROM golang:latest as build

WORKDIR /src

ADD go.mod go.sum ./

RUN go mod download

ADD . .

RUN go build -o /main src/main.go

FROM golang:latest

COPY --from=build /main /main

ENTRYPOINT ["/main"]
