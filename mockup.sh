#!/bin/bash

curl -XPOST "http://localhost:8066/write" \
    --data-binary 'http,status=200,hostname=alice rtime=60'