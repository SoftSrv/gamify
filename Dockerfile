FROM golang:1.13

ENV SRC_DIR=/usr/local/go/src/github.com/softsrv/gamify
ADD . $SRC_DIR 
RUN ls $SRC_DIR 
RUN cd $SRC_DIR; go get github.com/julienschmidt/httprouter 
RUN cd $SRC_DIR; go mod tidy
RUN cd $SRC_DIR; go mod vendor 
RUN cd $SRC_DIR/cmd/gamify; go build
WORKDIR /app
RUN mv $SRC_DIR/cmd/gamify/gamify /app/ 

EXPOSE 8080
CMD ["./gamify"] 
