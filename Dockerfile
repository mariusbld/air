FROM golang

WORKDIR /go/src/github.com/mariusbld/air
COPY . .

# RUN go get github.com/gorilla/mux
# RUN go get -u github.com/go-redis/redis

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 80

ENV AIR_CONFIG config/prod.json

CMD ["air"]
