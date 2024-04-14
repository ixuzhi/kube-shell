# https://docs.docker.com/language/golang/build-images/

FROM golang:1.20-alpine as builder
ENV GO111MODULE=on
WORKDIR /work
COPY go.mod .
COPY go.sum .
COPY . .
RUN go mod download && go build -o /app && go run *.go
# ENV HTTP_PORT=8080
# EXPOSE 8080
ENTRYPOINT ["tail", "-f", "/dev/null"]
# ENTRYPOINT ["sleep"]
# CMD ["/bin/bash"]
# CMD ["sh", "-c", "sleep infinity"]

# # BINARIES
# FROM alpine:latest
# COPY --from=builder /app /app
# ENTRYPOINT ["/app"]

