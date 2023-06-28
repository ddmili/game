
FROM golang:1.19 as builder

# ENV NODE_ENV build

WORKDIR /home/node

COPY . /home/node 

RUN  go version \
    && go mod tidy \
    &&  cd cmd/app \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /home/node/main  -tags v1 -v main.go

# ---

FROM alpine:latest

# ENV NODE_ENV prod

WORKDIR /home/node

# RUN apk --no-cache add ca-certificates

COPY --from=builder /home/node/main /home/node/

EXPOSE 8021
CMD ["./main"]