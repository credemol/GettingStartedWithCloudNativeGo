FROM golang:1.9.2-alpine3.7
MAINTAINER YG Kim(credemol@gmail.com)

ENV SOURCES ${GOPATH}/src/GettingStartedWithCloudNativeGo/

COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go install

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT GettingStartedWithCloudNativeGo
