#!/bin/bash
#########################################################################
# Author: (zfly1207@126.com)
# Created Time: 2021-07-19 10:42:53
# File Name: gen.sh
# Description: 
#########################################################################

find . -type f -name '*.proto' | xargs -n 1 protoc --proto_path=. --go_out=plugins=grpc:.

# vim: set noexpandtab ts=4 sts=4 sw=4 :
