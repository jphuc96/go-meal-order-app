#### Builder Stage
FROM golang:1.12.4-alpine3.9 as builder

WORKDIR /go/src/git.d.foundation/datcom/backend
COPY . .

# Enable go module
ENV GO111MODULE=on

# Set of commands
RUN apk add git make
RUN go get -v ./...
RUN make build
RUN go get -d -v ./...
RUN go install -v ./...

#### Runner Stage
FROM alpine:3.9

COPY cert.pem .
COPY privkey.pem .

RUN apk --no-cache add ca-certificates tzdata

#### Set timezone
RUN cp /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime
RUN echo "Asia/Ho_Chi_Minh" /etc/timezone

# Set of environments
ENV PORT=${PORT}
ENV DB_HOST=${DB_HOST}
ENV DB_PORT=${DB_PORT}
ENV DB_NAME=${DB_NAME}
ENV DB_USER=${DB_USER}
ENV DB_PASSWORD=${DB_PASSWORD}
ENV DB_SSL=${DB_SSL}

COPY --from=builder /go/bin/migrate .
COPY --from=builder /go/bin/server .

EXPOSE 80
EXPOSE 443

CMD ["./server"]