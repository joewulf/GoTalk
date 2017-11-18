#!/bin/sh

SRC_DIR=./
DST_DIR=./generated


if [ $1 = "c++" ]; then
	mkdir -p $DST_DIR/cpp
	protoc -I=$SRC_DIR --cpp_out=$DST_DIR/cpp/ $SRC_DIR/*.proto
elif [ $1 = "java" ]; then
	mkdir -p $DST_DIR/java
	protoc -I=$SRC_DIR --java_out=$DST_DIR/java/ $SRC_DIR/*.proto
elif [ $1 = "python" ]; then
	mkdir -p $DST_DIR/python
	protoc -I=$SRC_DIR --python_out=$DST_DIR/python/ $SRC_DIR/*.proto
elif [ $1 = "go" ]; then
	#for golang, it's a little manual work, see https://github.com/golang/protobuf/issues/39
	mkdir -p $DST_DIR/golang
	protoc -I=$SRC_DIR --go_out=$DST_DIR/golang/ $SRC_DIR/IM.BaseDefine.proto
	protoc -I=$SRC_DIR --go_out=$DST_DIR/golang/ $SRC_DIR/IM.File.proto
	protoc -I=$SRC_DIR --go_out=$DST_DIR/golang/ $SRC_DIR/IM.Login.proto
	protoc -I=$SRC_DIR --go_out=$DST_DIR/golang/ $SRC_DIR/IM.Other.proto
	protoc -I=$SRC_DIR --go_out=$DST_DIR/golang/ $SRC_DIR/IM.SwitchService.proto
	protoc -I=$SRC_DIR --go_out=$DST_DIR/golang/ $SRC_DIR/IM.Buddy.proto
	protoc -I=$SRC_DIR --go_out=$DST_DIR/golang/ $SRC_DIR/IM.Group.proto
	protoc -I=$SRC_DIR --go_out=$DST_DIR/golang/ $SRC_DIR/IM.Message.proto
	protoc -I=$SRC_DIR --go_out=$DST_DIR/golang/ $SRC_DIR/IM.Server.proto
else
	echo "unknown parameter, please use ./create [c++ java ptyhon go]"
fi

