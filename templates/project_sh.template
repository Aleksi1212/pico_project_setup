#!/usr/bin/env bash
set -eEuo pipefail

if [ $# -eq 0 ]; then
    echo "Error: No arguments provided. Arguments: <build | run | debug | clean>"
    exit 1
fi

BUILD_DIR="build"
PROJECT_NAME=$(<project_name.txt)

check_build_dir() {
    if [ ! -d "$BUILD_DIR" ]; then
        if [ "${1:-}" == "create" ]; then
            mkdir -p "$BUILD_DIR"
        else
            echo "Error: Build directory does not exist. Please run './project.sh build' first."
            exit 1
        fi
    fi
}

if [ "$1" == "build" ]; then
    check_build_dir "create"

    cd "$BUILD_DIR"
    cmake ..
    make -j$(nproc)

    exit 0
fi

if [ "$1" == "run" ]; then
    check_build_dir

    cd "$BUILD_DIR"
    openocd -f interface/cmsis-dap.cfg -f target/rp2040.cfg -c "adapter speed 8000; program $PROJECT_NAME.elf verify reset exit"

    exit 0
fi

if [ "$1" == "debug" ]; then
    check_build_dir

    gnome-terminal -- bash -c "gdb-multiarch $BUILD_DIR/$PROJECT_NAME.elf; exec bash"
    openocd -f interface/cmsis-dap.cfg -f target/rp2040.cfg -c "adapter speed 8000"
fi

if [ "$1" == "clean" ]; then
    rm -rf "$BUILD_DIR"
    exit 0
fi
