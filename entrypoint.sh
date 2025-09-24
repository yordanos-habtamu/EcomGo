#!/bin/sh

if [ -n "$MYSQL_PUBLIC_URL" ]; then
  # Use the Railway-provided URL directly for migrate
  MIGRATE_DSN="$MYSQL_PUBLIC_URL"
else
  echo "MYSQL_PUBLIC_URL is not set"
  exit 1
fi

echo "Running migrations with DSN: $MIGRATE_DSN"
migrate -path /app/migrations -database "$MIGRATE_DSN" up

echo "Starting app..."
exec ./myapp