#!/usr/bin/env bash
set -eEuo pipefail

PICO_PROJECT_SETUP_PATH="/etc/pico_project_setup"

sudo mkdir -p "$PICO_PROJECT_SETUP_PATH"

sudo cp templates/* "$PICO_PROJECT_SETUP_PATH"

go build