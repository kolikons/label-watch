ARG BASE_IMAGE=scratch

FROM docker.io/library/golang:1.16 as builder
ADD . /go/src/github.com/kolikons/label-watch
WORKDIR /go/src/github.com/kolikons/label-watch
RUN go mod vendor && go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o label-watch .

FROM ${BASE_IMAGE}
LABEL maintainer="kolikons@gmail.com"
COPY --from=builder /go/src/github.com/kolikons/label-watch/label-watch ./
ENTRYPOINT ["./label-watch"]
