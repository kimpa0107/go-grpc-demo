function genProto {
    DOMAIN=$1

    PROTO_PATH=./proto
    GO_OUT_PATH=./${PROTO_PATH}/${DOMAIN}/gen

    mkdir -p $GO_OUT_PATH

    # generate grpc go file
    protoc -I=$PROTO_PATH --go_out=plugins=grpc,paths=source_relative:$GO_OUT_PATH ${DOMAIN}.proto

}

genProto blog