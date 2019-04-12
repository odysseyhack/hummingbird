import 'package:flutter/material.dart';

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
              onPressed: () {
                // call grpc
              },
            )
          ],
        ),
      ),
    );
  }
}
