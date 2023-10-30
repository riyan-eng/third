# build
FROM golang:1.21.1-bullseye AS build-stage
WORKDIR /app
COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /binary

# release
FROM alpine:3.18.4 AS release-stage
WORKDIR /
COPY --from=build-stage /binary /binary
CMD ["/binary"]