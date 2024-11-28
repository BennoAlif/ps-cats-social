# syntax=docker/dockerfile:1

ARG GO_VERSION=1.18
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS build
WORKDIR /src

# Install golang-migrate
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.17.1

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

ARG TARGETARCH

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 GOARCH=$TARGETARCH go build -o /bin/server ./src

FROM alpine:latest AS migration
# Copy migration files
COPY ./db/migrations /migrations

# Install golang-migrate in the migration stage
RUN apk add --no-cache curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz | tar xvz

FROM alpine:latest AS final

RUN --mount=type=cache,target=/var/cache/apk \
    apk --update add \
        ca-certificates \
        tzdata \
        && \
        update-ca-certificates

ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    appuser

# Copy migration binary and files
COPY --from=migration /migrate /bin/migrate
COPY --from=migration /migrations /migrations

# Copy server binary
COPY --from=build /bin/server /bin/

ARG DATABASE_URL
ENV DATABASE_URL=$DATABASE_URL

COPY <<EOF /entrypoint.sh
#!/bin/sh
set -e

# Run database migrations
/bin/migrate -path /migrations -database "$DATABASE_URL" up

# Start the server
exec /bin/server
EOF

# Use root to set permissions, then switch back to appuser
USER root
RUN chmod +x /entrypoint.sh
USER appuser

EXPOSE 8080

ENTRYPOINT ["/entrypoint.sh"]