FROM golang:1.13 as build-env

WORKDIR /go/src/app
ADD . /go/src/app/

RUN cd /go/src/app/ && go mod download

RUN go build -o /go/bin/app

FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/app /
CMD ["/app"]
