FROM golang:1.22.4-alpine AS build

RUN mkdir /app
WORKDIR /app

COPY ./backend/timeline/ ./
RUN go mod download
RUN go build -o /myapp

FROM gcr.io/distroless/base-debian12:latest-amd64
WORKDIR /app
COPY --from=build /myapp /myapp
EXPOSE 9000

ENTRYPOINT ["/myapp"]
