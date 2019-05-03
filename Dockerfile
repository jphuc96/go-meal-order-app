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

RUN apk --no-cache add ca-certificates

# Set of environments
ENV PORT=${PORT}
ENV DB_HOST=${DB_HOST}
ENV DB_PORT=${DB_PORT}
ENV DB_NAME=${DB_NAME}
ENV DB_USER=${DB_USER}
ENV DB_PASSWORD=${DB_PASSWORD}
ENV DB_SSL=${DB_SSL}
ENV GOOGLE_REDIRECT_URL=${GOOGLE_REDIRECT_URL}
ENV GOOGLE_OAUTH_CLIENT_ID=${GOOGLE_OAUTH_CLIENT_ID}
ENV GOOGLE_OAUTH_CLIENT_SECRET=${GOOGLE_OAUTH_CLIENT_SECRET}

COPY --from=builder /go/bin/migrate .
COPY --from=builder /go/bin/server .

EXPOSE 8080

CMD ["./server"]