FROM golang:1.17-bullseye as builder
WORKDIR /opt/app
COPY go.mod ./
RUN go mod download
COPY ./*.go ./
COPY ./handler/*.go ./handler/
RUN go build -trimpath -ldflags="-w -s" -o "app"

FROM gcr.io/distroless/base-debian11
COPY --from=builder /opt/app/app /app
CMD ["/app"]