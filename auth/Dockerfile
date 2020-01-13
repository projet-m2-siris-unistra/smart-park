FROM golang

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN go build -ldflags "-linkmode external -extldflags -static" -a main.go

FROM gcr.io/distroless/base
COPY --from=0 /app/main /main
CMD ["/main"]
