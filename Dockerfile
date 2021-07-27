FROM golang:1.16

ENV PATH $GOPATH/bin:/root/.yarn/bin:$PATH
 
RUN apt-get update && \
    apt-get install git