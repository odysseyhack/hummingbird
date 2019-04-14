# Hummingbird, smart decentralised communication platform

In this repository you will find:
* ./app -> the community native Android app (Flutter SDK)
* ./dashboard -> emergency workers dashboard, overview of verified data retrieved by Hiber's low earth orbit satellites
* ./hiber-bridge -> golang utility to bridge gRPC API request to the Hiber's hardware modem through UART
* ./hiber-faker -> Hiber utility (outside scope hackathon) to simulate satellites: virtual desktop satellite
* ./protobuf -> gRPC API protocol buffer
* ./vendor -> vendored code

# Architecture

![Architecture overview](https://github.com/odysseyhack/hummingbird/blob/develop/Hummingbird_Architecture1.png)

![Architecture overview](https://github.com/odysseyhack/hummingbird/blob/develop/Hummingbird_Architecture2.png)
