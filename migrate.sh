#!/bin/bash
migrate -source "file://./examples/default/db/migrations" -database "postgres://postgres:password@127.0.0.1:54321/example?sslmode=disable" $@
