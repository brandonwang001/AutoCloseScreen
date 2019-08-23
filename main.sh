#!/usr/bin/env bash
function makeRelativeDyLib() {
    local dylib="auto_close_screen"
    local dependency
    for dependency in $(otool -L ${dylib} | otool -L auto_close_screen | cut -d " " -f 1)
    do
        echo $dependency
        $(cp ${dependency} ./lib)
        install_name_tool -change ${dependency} @loader_path/lib/$(basename ${dependency}) ${dylib}
    done
}

function main() {
    makeRelativeDyLib
}

main
