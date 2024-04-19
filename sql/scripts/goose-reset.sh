#!/bin/bash

if [ -f .env ]; then
    source .env
fi

cd sql/migrations
goose turso "$AUTH_DATABASE_URL" reset
