FROM golang:1.10

WORKDIR /go/src/github.com/currency-converter
COPY . .

RUN go get -u github.com/golang/dep/cmd/dep \
    && dep ensure -v \
    && go build \
    && go install \
    && cp shared/config/docker.yml shared/config/default.yml

EXPOSE 8080
VOLUME [ "/var/log/converter-server" ]

ENTRYPOINT [ "currency-converter","-migrate", "-seed", "-log_dir", "/var/log/converter-server", "-alsologtostderr", "-stderrthreshold", "warning", "-v", "2" ]