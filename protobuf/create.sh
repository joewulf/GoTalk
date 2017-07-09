#!/bin/sh
SRC_DIR=./
DST_DIR=./gen

#C++
mkdir -p $DST_DIR/cpp
protoc -I=$SRC_DIR --cpp_out=$DST_DIR/cpp/ $SRC_DIR/*.proto

#JAVA
mkdir -p $DST_DIR/java
protoc -I=$SRC_DIR --java_out=$DST_DIR/java/ $SRC_DIR/*.proto

#PYTHON
mkdir -p $DST_DIR/python
protoc -I=$SRC_DIR --python_out=$DST_DIR/python/ $SRC_DIR/*.proto

#GOLANG
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

