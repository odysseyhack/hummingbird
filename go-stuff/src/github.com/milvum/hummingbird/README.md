# Overview
This directory should allow us to properly make use of our existing grpc definitions, using Skywire as a bridge.

From a high level, we can define the following logical system parts;
- A grpc Hummingbird server that accepts input from mobile devices
- A skywire mainnet 'App' that sends 'something' forward (perhaps the bytes of the grpc request?)
- A skywire mainnet 'App' that receives 'something' over skywire (and hopefully deserializes this)
- A grpc Hummingbird client that sends the parsed input to mobile devices

Furthermore, with a certain interval updates need to be send to Hiber
mission control.  This is done by periodically seeding a HiberBridge
grpc service with gathered data, and calling out to the (hopefully
working) grpc server hosted on our rpi.

# Building

Make sure you have your direnv set up properly.
To build the chat-client:
`go build ./cmd/chat-client`

To build the chat-server:
`go build ./cmd/chat-server`

Run the resulting binaries with `-h` to read about any runtime
configuration options.
