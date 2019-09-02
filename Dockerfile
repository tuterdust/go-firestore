FROM golang:1.12.0-alpine


ENV ROOTPATH=$GOPATH/src/github.com/tuterdust/go-firestore/src
RUN mkdir -p $GOPATH/src/github.com/tuterdust/go-firestore
ADD . $GOPATH/src/github.com/tuterdust/go-firestore
WORKDIR $ROOTPATH
RUN cd $ROOTPATH
RUN pwd
CMD go run *.go  

