FROM golang:1.11.4-alpine3.8 as build

COPY . /go/src/github.com/d7561985/heroku_boilerplate/
WORKDIR /go/src/github.com/d7561985/heroku_boilerplate/
RUN apk add --no-cache git; \
    go get ./...; \
    go build -o /heroku cmd/server/main.go

FROM alpine:3.8

COPY --from=build /heroku /app

# init certificates for https connection
RUN apk add --no-cache libstdc++ \
	ca-certificates

# add heroku user
RUN adduser -D -u 1000 heroku
USER heroku

CMD /app