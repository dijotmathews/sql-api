FROM golang:1.20 AS builder

COPY ./ /app
COPY ./.env.vm /app/.env

WORKDIR /app
RUN ls
# Toggle CGO on your app requirement
RUN CGO_ENABLED=0 go build -ldflags '-s -w -extldflags "-static"' -o /app/lending-los main.go
RUN CGO_ENABLED=0 go build -ldflags '-s -w -extldflags "-static"' -o /app/migrate-lending-los database/migrate/migrate.go
RUN CGO_ENABLED=0 go build -ldflags '-s -w -extldflags "-static"' -o /app/seed-los database/seeder/seeder.go

# Use below if using vendor
# RUN CGO_ENABLED=0 go build -mod=vendor -ldflags '-extldflags "-static"' -o /app/lending-los *.go

FROM debian:stable-slim
LABEL MAINTAINER Dijo T Mathews <dijo.mathews@bankopen.co>

# Following commands are for installing CA certs (for proper functioning of HTTPS and other TLS)
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates  \
    netbase \
    && rm -rf /var/lib/apt/lists/ \
    && apt-get autoremove -y && apt-get autoclean -y

# Add new user 'appuser'. App should be run without root privileges as a security measure
RUN adduser --home "/home/appuser" --disabled-password appuser \
    --gecos "appuser,-,-,-"
USER appuser

# COPY --from=builder /app/internal/server/http/web /home/appuser/app/web
RUN mkdir -p /home/appuser/app
COPY --from=builder /app/lending-los /home/appuser/app
COPY --from=builder /app/.env.vm /home/appuser/app/.env
COPY --from=builder /app/migrate-lending-los /home/appuser/app
COPY --from=builder /app/seed-los /home/appuser/app
# ENV TEMPLATES_BASEPATH=/home/appuser/app/web/templates

WORKDIR /home/appuser/app

# Since running as a non-root user, port bindings < 1024 are not possible
# 8000 for HTTP; 8443 for HTTPS;
EXPOSE 9090

CMD ["./lending-los"]