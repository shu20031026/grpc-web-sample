#!/bin/sh

protoc proto/helloworld.proto \
    --js_out=import_style=commonjs:cliant/helloworld \
    --grpc-web_out=import_style=commonjs+dts,mode=grpcwebtext:cliant/helloworld \
    --go-grpc_out=api/helloworld