#!/bin/sh

if [ -n "$MYSQL_PUBLIC_URL" ]; then
  # Convert mysql://user:pass@host:port/db to mysql://user:pass@tcp(host:port)/db
  MIGRATE_DSN=$(echo "$MYSQL_PUBLIC_URL" | sed -E 's|mysql://([^:]+):([^@]+)@([^:]+):([0-9]+)/([^?]+)|mysql://\1:\2@tcp(\3:\4)/\5|')
else
  echo "MYSQL_PUBLIC_URL is not set"
  exit 1
fi

echo "Running migrations with DSN: $MIGRATE_DSN"
migrate -path /app/migrations -database "$MIGRATE_DSN" up

echo "Starting app..."
exec ./myapp