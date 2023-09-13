#!/bin/sh

# Wait for PostgreSQL to be ready
while ! nc -z postgres 5432; do
  echo 'Waiting for PostgreSQL to become available...'
  sleep 1
done

./main
