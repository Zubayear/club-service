#!/bin/bash
echo 'generating'

`protoc --proto_path=. --micro_out=./client/ --go_out=:./client/ client/proto/player.proto`