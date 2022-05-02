FROM golang:1.17-alpine3.15 AS builder
WORKDIR /app
COPY main.go .
#RUN export GOPROXY=https://proxy.golang.com.cn,direct
RUN go build  main.go

# run
FROM alpine:3.15.4 AS runner
WORKDIR /app

COPY --from=builder /app/main ./app
RUN chmod 775 main

EXPOSE 8080

ENTRYPOINT ["./main"]




