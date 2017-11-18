#!/bin/sh
CPP_DIR=../server/src/base/pb/protocol
DST_DIR=./generated

#C++
cp $DST_DIR/cpp/* $CPP_DIR/

rm -rf ./gen
