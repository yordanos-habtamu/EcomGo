#!/bin/sh

# Convert Railway's MYSQL_PUBLIC_URL to Go/MySQL driver format
if [ -n "$MYSQL_PUBLIC_URL" ]; then
  # Example: mysql://user:pass@host:port/db
  USER=$(echo "$MYSQL_PUBLIC_URL" | sed -E 's|mysql://([^:]+):.*|\1|')
  PASS=$(echo "$MYSQL_PUBLIC_URL" | sed -E 's|mysql://[^:]+:([^@]+)@.*|\1|')
  HOSTPORT=$(echo "$MYSQL_PUBLIC_URL" | sed -E 's|mysql://[^@]+@([^/]+)/.*|\1|')
  DBNAME=$(echo "$MYSQL_PUBLIC_URL" | sed -E 's|.*/([^?]+).*|\1|')
  MIGRATE_DSN="$USER:$PASS@tcp($HOSTPORT)/$DBNAME"
else
  echo "MYSQL_PUBLIC_URL is not set"
  exit 1
fi

echo "Running migrations with DSN: $MIGRATE_DSN"
migrate -path /app/migrations -database "$MIGRATE_DSN" up

echo "Starting app..."
exec ./myapp