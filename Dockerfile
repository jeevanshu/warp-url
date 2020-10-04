FROM golang:1.14.4-alpine3.12 AS build

RUN apk update \
    && apk add --no-cache git libc6-compat\
    && apk add --no-cache --repository=http://dl-cdn.alpinelinux.org/alpine/edge/community \
    && apk add --no-cache grpc \
    && go get google.golang.org/grpc

WORKDIR $GOPATH/src
COPY . .
RUN go build -o $GOPATH/bin/main cmd/server/main.go
RUN chmod +x $GOPATH/bin/main


FROM golang:1.14.4-alpine3.12
RUN apk update \
    && apk add --no-cache git libc6-compat\
    && apk add --no-cache --repository=http://dl-cdn.alpinelinux.org/alpine/edge/community \
    && apk add --no-cache grpc \
    && go get google.golang.org/grpc

COPY --from=build $GOPATH/bin/main $GOPATH/bin

CMD [ "main" ]


