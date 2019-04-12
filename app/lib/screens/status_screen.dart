import 'dart:async';

import 'package:flutter/material.dart';

import 'package:grpc/grpc.dart';

import 'package:hummingbird/generated/hummingbird.pb.dart';
import 'package:hummingbird/generated/hummingbird.pbgrpc.dart';
import 'package:hummingbird/network.dart';

class StatusScreen extends StatefulWidget {
  StatusScreen({Key key, this.title}) : super(key: key);

  final String title;

  @override
  _StatusScreenState createState() => _StatusScreenState();
}

class _StatusScreenState extends State<StatusScreen> {
  int _counter = 0;

  void _incrementCounter() {
    setState(() {
      _counter++;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            Text(
              'Tap icons to cast your needs in your status.',
            ),
            GridView.count(
              crossAxisCount: 3,
              shrinkWrap: true,
              padding: EdgeInsets.all(20),
              children: <Widget>[
                const Text('Water'),
                const Text('Food'),
                const Text('Electricity'),
                const Text('Fire'),
                const Text('Derp'),
                const Text('Medicine'),
              ],
            ),
            RaisedButton(
              child: Text("Cast again"),
              onPressed: () async {
                final channel =
                    Network.createInsecureChannel('192.168.8.190', 50051);
                final stub = HummingbirdClient(channel);

                final name = 'world';

                try {
                  var response =
                      await stub.castStatuses(StatusRequest()..origin = name);
                  print('Greeter client received: ${response.success}');
                } catch (e) {
                  print('Caught error: $e');
                }
                await channel.shutdown();
              },
            )
          ],
        ),
      ),
    );
  }
}
