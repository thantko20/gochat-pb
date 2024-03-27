# FROM alpine:latest

# ARG PB_VERSION=0.22.5

# RUN apk add --no-cache \
#   unzip \
#   ca-certificates

# # download and unzip PocketBase
# # ADD https://github.com/pocketbase/pocketbase/releases/download/v${PB_VERSION}/pocketbase_${PB_VERSION}_linux_amd64.zip /tmp/pb.zip
# # RUN unzip /tmp/pb.zip -d /pb/

# COPY ./pb_migrations /pb/pb_migrations

# # uncomment to copy the local pb_hooks dir into the image
# # COPY ./pb_hooks /pb/pb_hooks

# EXPOSE 8080

# # start PocketBase
# CMD ["/pb/pocketbase", "serve", "--http=0.0.0.0:8080"]

FROM golang:latest as builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY main.go .
RUN go build -v -o /run-app .


FROM debian:latest

RUN apt update && apt -y install ca-certificates
COPY --from=builder /run-app /pb/run-app

# COPY ./pb/pb_migrations /pb/pb_migrations

CMD [ "/pb/run-app", "serve", "--http=0.0.0.0:8080" ]