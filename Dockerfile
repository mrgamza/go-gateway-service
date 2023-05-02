FROM golang:1.15

ENV GOPATH /go
ARG service=gateway-service

WORKDIR $GOPATH/src/${service}

COPY go.mod .
COPY go.sum .
RUN go mod download

ADD . /go/src/${service}

EXPOSE 80

CMD ["go", "run", "gateway-service"]