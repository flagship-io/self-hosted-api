FROM golang:1-alpine

WORKDIR /go/src/github.com/flagship-io/self-hosted-api

ADD . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /bin/app .


FROM gcr.io/distroless/base

COPY --from=0 ./bin/app /bin/app

CMD ["/bin/app"]