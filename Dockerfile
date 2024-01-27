FROM golang:alpine AS build

WORKDIR /usr/src/app

COPY . .

RUN go build -o /go/hello ./cmd/app

FROM scratch

COPY --from=build /go/hello /go/hello

CMD ["/go/hello"]