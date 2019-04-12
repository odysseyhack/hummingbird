#!/bin/bash

# Make ARM5 binary
cd ~/go-serial-test
make all

# Copy config file
cp ./config.json ./build/config.json

# Update RPi
scp -r build pi@rpi:/home/pi/hiber