FROM golang:1.10

WORKDIR /go/src/github.com/currency-converter
COPY . .

RUN go get -u github.com/golang/dep/cmd/dep \
    && dep ensure -v \
    && go build \
    && go install \
    && cp shared/config/docker.yml shared/config/default.yml

EXPOSE 8080
VOLUME [ "/var/log/locum-server" ]

ENTRYPOINT [ "Locum-Server","-migrate", "-log_dir", "/var/log/locum-server", "-alsologtostderr", "-stderrthreshold", "warning", "-v", "2" ]
#ENTRYPOINT ["LocumServer"]