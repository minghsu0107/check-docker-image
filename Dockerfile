FROM golang:1.15-alpine3.14 as builder

# Install git + SSL ca certificates
# Git is required for fetching the dependencies
# Ca-certificates is required to call HTTPS endpoints
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# Create appuser
ENV USER=appuser
ENV UID=10001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

RUN mkdir -p /app
WORKDIR /app

COPY go.mod .
ENV GO111MODULE=on
RUN go mod download
RUN go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o check-image .

FROM scratch

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /app/check-image /go/bin/check-image

# Use an unprivileged user
USER appuser:appuser

ENTRYPOINT ["/go/bin/check-image"]