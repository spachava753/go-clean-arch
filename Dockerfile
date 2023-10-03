FROM golang:1.20-alpine as builder

ARG VERSION

RUN mkdir -p /app/

WORKDIR /app

# The next to steps are important to speed up building

# This will only copy go.mod and go.sum
COPY go.* .

# This download the necessary modules
RUN go mod download

# Now we can copy the rest of the files. This ensures that the previous two steps will always be cached
COPY 02_clean_arch .

RUN CGO_ENABLED=0 go build -ldflags "-s -w \
    -X 'main.Version=${VERSION}' \
    -X 'main.BuildDate=$(date)'" \
    -a -o bin/app main.go

FROM alpine:3.18

RUN addgroup -S app \
    && adduser -S -G app app \
    && apk --no-cache add \
    ca-certificates curl netcat-openbsd

WORKDIR /home/app

COPY --from=builder /app/bin/app .
RUN chown -R app:app ./

USER app

CMD ["./app"]