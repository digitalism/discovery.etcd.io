FROM centos:centos7
MAINTAINER Meng Zhuo "mengzhuo1203@gmail.com"

RUN yum install golang -y
RUN mkdir /srv/discovery

EXPOSE 8087

ADD . /srv/discovery
RUN cd /srv/discovery && ./build
COPY ./init/discovery.service /etc/systemd/user/

ENTRYPOINT ['/srv/discovery/discovery']
