FROM golang:1.13 AS builder
# Use ADD in case of a URL only.
COPY . /app/backend
WORKDIR /app/backend
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /main .

# final stage
FROM alpine:3.10
COPY --from=builder /main ./
RUN chmod +x ./main
ENTRYPOINT ["./main"]
EXPOSE 8000
