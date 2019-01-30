FROM golang:latest

RUN go get -u github.com/lib/pq github.com/jinzhu/gorm

RUN mkdir $GOPATH/src/model
COPY app/src/model $GOPATH/src/model/
WORKDIR $GOPATH/src/model
RUN go build && go install

RUN mkdir $GOPATH/src/repository
COPY app/src/repository $GOPATH/src/repository/
WORKDIR $GOPATH/src/repository
RUN go build && go install

RUN mkdir $GOPATH/src/config
COPY app/src/config $GOPATH/src/config/
WORKDIR $GOPATH/src/config
RUN go build && go install

RUN mkdir $GOPATH/src/route
COPY app/src/route $GOPATH/src/route/
WORKDIR $GOPATH/src/route
RUN go build && go install

COPY app/src/main.go /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]