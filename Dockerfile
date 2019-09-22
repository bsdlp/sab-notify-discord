FROM golang:alpine as builder
RUN apk update && apk upgrade && apk add --no-cache ca-certificates
RUN update-ca-certificates
ADD ./cmd/sab-notify-discord ./cmd
ENV GOOS=linux GOARCH=amd64 CGO_ENABLED=0
RUN go build -a -ldflags="-w -s" -installsuffix cgo -o /go/sab-notify-discord ./cmd

FROM scratch
COPY --from=builder /go/sab-notify-discord /sab-notify-discord
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/sab-notify-discord"]
