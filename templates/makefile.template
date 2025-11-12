# Makefile

BUILD_DIR := build
PROJECT_NAME := $(shell cat project_name.txt)

.PHONY: all build flash clean

all: build flash

build:
	mkdir -p $(BUILD_DIR)
	cd $(BUILD_DIR) && cmake ..
	cd $(BUILD_DIR) && make

flash:
	openocd -f interface/cmsis-dap.cfg \
		-f target/rp2040.cfg \
		-c "adapter speed 5000" \
		-c "program $(BUILD_DIR)/$(PROJECT_NAME).elf verify reset exit"

clean:
	rm -rf $(BUILD_DIR)