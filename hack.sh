#!/usr/bin/env bash

set -e

build () {
    docker build -t simpleproxy:$(cat version) .
}

run () {
    docker run -p 80:3000 -ti simpleproxy:$(cat version) $@
}

case $1 in
"build")
    build
    ;;

"run")
    shift
    run $@
    ;;
*)
    echo -e "Usage:\n ./hack build\n ./hack run <args>"
    exit 1
    ;;
esac