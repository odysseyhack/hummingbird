#!/bin/bash

# Make ARM5 binary
cd ~/go-serial-test
make all

# Copy config file
cp ./config.json ./build/config.json

mv ./build/hummingbird-hiber-bridge-linux-arm ./build/hiber-bridge

# Update RPi
scp -r build pi@rpi:/home/pi/hiber