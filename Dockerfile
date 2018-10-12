FROM golang:jessie
WORKDIR /go
# RUN go get ...
RUN go get github.com/gorilla/mux
RUN go get github.com/rs/cors
RUN go get gopkg.in/russross/blackfriday.v2
RUN go get -u github.com/go-redis/redis
RUN go get github.com/gorilla/securecookie
RUN go get github.com/avelino/slugify
RUN go get github.com/brianvoe/gofakeit
RUN go get github.com/json-iterator/go
RUN go get -u github.com/go-redis/cache
RUN go get github.com/vmihailenco/msgpack
RUN go get github.com/prometheus/client_golang/prometheus
RUN go get golang.org/x/oauth2

# Copy the server code into the container
COPY . /go

EXPOSE 443

RUN go build
ENTRYPOINt ["./go"]