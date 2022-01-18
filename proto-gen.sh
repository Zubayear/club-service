#!/bin/bash
echo 'generating stub from proto file'

`protoc --proto_path=. --micro_out=. --go_out=:. client/player.proto`