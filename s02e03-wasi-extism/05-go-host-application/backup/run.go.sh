#!/bin/bash
clear
bat $0 --line-range 5:
echo ""
# args: wasm_file function_name config
./hostapp ../03-go-plugin/go-plugin.wasm \
say_hello \
"Bob Morane"

echo ""