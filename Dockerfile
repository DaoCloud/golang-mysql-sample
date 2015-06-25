FROM google/golang
MAINTAINER Sakeven "sakeven.jiang@daocloud.io"

# Build app
WORKDIR /gopath/app
ENV GOPATH /gopath/app
ADD . /gopath/app/src/golang-mysql-sample

RUN go get -t golang-mysql-sample
RUN go install golang-mysql-sample

EXPOSE 80
CMD ["/gopath/app/bin/golang-mysql-sample"]
