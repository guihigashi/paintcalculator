FROM golang:1.19-alpine AS build

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY cmd cmd
COPY pkg pkg

RUN go build -v -o /usr/local/bin/app ./cmd/web

# deploy
FROM alpine:latest

COPY --from=build /usr/local/bin/app /usr/local/bin/app

CMD ["app"]
