
FROM golang

WORKDIR /app

COPY . .
RUN go mod edit -dropreplace=github.com/projet-m2-siris-unistra/smart-park/backend
RUN go build -mod vendor -ldflags "-linkmode external -extldflags -static" -a main.go

FROM scratch
COPY --from=0 /app/main /main
CMD ["/main"]
