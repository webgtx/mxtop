#!/bin/sh

# project info
NAME='mxtop'
MODULE_NAME='github.com/webgtx/mxtop'
VERSION='0.0.1'
COMMIT_HASH="$(git rev-parse --short HEAD)"

# go build flags
if [ -z "$DEBUG_VAR" ]; then
    DEBUG_VAR="OFF"
else
    DEBUG_VAR="ON"
fi
GO_VARS="\
    -X ${MODULE_NAME}/internal/info.Version=${VERSION} \
    -X ${MODULE_NAME}/internal/info.CommitHash=${COMMIT_HASH} \
    -X ${MODULE_NAME}/internal/info.Debug=${DEBUG_VAR} \
"
GO_LINKER_FLAGS="\
    -s \
    -w \
    ${GO_VARS} \
"

build() {
    if ! command -v go &> /dev/null; then
        echo 'go compiler doesnt finded in system'
        exit 1
    fi
    go build "-ldflags=${GO_LINKER_FLAGS}" "cmd/${NAME}/${NAME}.go"

}

main() {
    if [ -z "$1" ]; then
        return 1
    fi

    local job="$1"
    shift 1
    "${job}" "$@"
}

main "$@"