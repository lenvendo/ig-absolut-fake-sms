FROM golang:1.16 as golang-builder

LABEL Component="lv-fakesms-build"

WORKDIR /src

COPY . .

RUN apt-get install -y ca-certificates
RUN make build-docker

FROM scratch
LABEL Component="lv-fakesmsn"
WORKDIR /app
COPY --from=golang-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=golang-builder /src/bin /app

CMD ["/app/app"]
