FROM golang:1.18-bullseye as builder
WORKDIR /opt/app
COPY go.mod ./
COPY ./*.go ./
COPY ./handler/*.go ./handler/
COPY ./infra/*.go ./infra/
COPY ./repository/*.go ./repository/
COPY ./usecase/*.go ./usecase/
COPY ./domain/*.go ./domain/
RUN go mod download
RUN go mod tidy
RUN go build -trimpath -ldflags="-w -s" -o "app"

FROM gcr.io/distroless/base-debian11
COPY --from=builder /opt/app/app /app
CMD ["/app"]