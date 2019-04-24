FROM centos:latest

USER root

WORKDIR /go

RUN yum update -y && yum install -y epel-release \
	vim \
	wget \
	git \
	zsh

RUN wget https://dl.google.com/go/go1.12.3.linux-amd64.tar.gz && tar vzfx go1.12.3.linux-amd64.tar.gz && mv /go/go /usr/bin/ && mkdir -p /DevSpace

ENV PATH $PATH:/usr/bin/go/bin
ENV GOPATH /go
ENV PATH $PATH:/go/bin
ENV TZ Asia/Tokyo


