FROM golang:1.17-alpine AS builder
MAINTAINER kineo2k <kineo2k@gmail.com>

WORKDIR /build
COPY . ./

RUN go mod download
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o go-pipeline-sample -a -installsuffix cgo -ldflags '-s' go_pipeline_sample.go



FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/go-pipeline-sample .

ENTRYPOINT ["./go-pipeline-sample", "production"]
