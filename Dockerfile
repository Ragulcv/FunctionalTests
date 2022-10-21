FROM golang:latest

RUN  mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/ragulcv/FunctionalTests/main
RUN cd /build && git clone https://github.com/Ragulcv/FunctionalTests.git

RUN cd /build/FunctionalTests/main && go build

EXPOSE 8080

ENTRYPOINT [ "/buildFunctionalTests/main/main" ]