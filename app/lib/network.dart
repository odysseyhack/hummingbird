import 'package:grpc/grpc.dart';

class Network {
  static ClientChannel createInsecureChannel(String url, int port) {
    return ClientChannel(url,
        port: 50051,
        options: const ChannelOptions(
          credentials: const ChannelCredentials.insecure(),
        ));
  }
}
