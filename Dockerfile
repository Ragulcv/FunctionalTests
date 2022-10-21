FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN cd /build && git clone https://github.com/Ragulcv/FunctionalTests.git

RUN cd \Users\SNEKA\FunctionalTests\main && go build
RUN go run test-v