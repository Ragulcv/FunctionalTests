FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN cd /test && git clone https://github.com/Ragulcv/FunctionalTests.git

RUN cd /build/FunctionalTests/main && go build
RUN go run test-v