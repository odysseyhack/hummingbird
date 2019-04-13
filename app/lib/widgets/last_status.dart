import 'package:flutter/material.dart';

import 'package:hummingbird/generated/hummingbird.pb.dart';
import 'package:hummingbird/widgets/colored_orb.dart';

class LastStatus extends StatelessWidget {
  final Statuses status;
  final DateTime timestamp;

  LastStatus({Key key, this.status, this.timestamp}) : super(key: key);

  @override
  Widget build(BuildContext context) => Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: <Widget>[
          Text(
            "My last cast status",
            style: TextStyle(fontWeight: FontWeight.bold, fontSize: 20),
          ),
          Row(
            children: <Widget>[
              ColoredOrb(Colors.red),
              Text("Badly needs food, needs fire"),
            ],
          ),
          RaisedButton(
            onPressed: () {},
            color: Colors.blue,
            child: Container(
                constraints: BoxConstraints(minWidth: double.infinity),
                child: Center(child: Text("Cast again"))),
          ),
        ],
      );
}
