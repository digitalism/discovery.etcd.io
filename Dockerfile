FROM 
MAINTAINER Meng Zhuo "mengzhuo1203@gmail.com"

RUN apt-get update
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y golang

RUN mkdir /srv/discovery

EXPOSE 8087

ADD . /srv/discovery
RUN cd /srv/discovery && ./build
COPY ./init/discovery.service /etc/systemd/user/

RUN systemctl enable /etc/systemd/user/discovery.service

RUN systemctl daemon-reload
RUN systemctl start discovery.service
