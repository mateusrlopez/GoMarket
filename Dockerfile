FROM golang:1.20 AS builder

WORKDIR /usr/src/app

ENV GO11MODULE on
ENV CGO_ENABLED 0
ENV GOOS linux

COPY . .

RUN go build -o bin/main cmd/main.go

FROM scratch

COPY --from=builder /usr/src/app/bin/main .

CMD [ "main" ]
