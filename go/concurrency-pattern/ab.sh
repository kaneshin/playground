#!/bin/sh
ab -n 2000 -c 2000 -p ./data.json -T "application/json; charset=utf-8" "http://localhost:8080/"
