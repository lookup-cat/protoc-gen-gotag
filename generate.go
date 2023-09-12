package main

//go:generate protoc --proto_path=./third  --proto_path=. ./tagger/tagger.proto --go_out=. --go_opt=paths=source_relative
