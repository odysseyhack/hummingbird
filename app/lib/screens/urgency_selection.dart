import 'package:flutter/material.dart';

import 'package:hummingbird/status_enum.dart';
import 'package:hummingbird/widgets/colored_orb.dart';

class UrgencySelection extends StatelessWidget {
  final Status status;

  UrgencySelection(this.status);

  @override
  Widget build(BuildContext context) {
    var statements = statusToUrgencyStatements(status);
    return Scaffold(
      appBar: AppBar(
        title: Text("Hummingbird"),
      ),
      body: Container(
        padding: EdgeInsets.all(20),
        constraints: BoxConstraints(minWidth: double.infinity),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: List<Widget>()
            ..add(Text("Describe your situation"))
            ..addAll(statements
                .map((us) => _buildStatement(context, us))
                .toList(growable: false)),
        ),
      ),
    );
  }

  Widget _buildStatement(BuildContext context, UrgencyStatement us) =>
      Container(
        margin: EdgeInsets.only(top: 20),
        child: RaisedButton(
          child: Container(
              height: 80,
              padding: EdgeInsets.all(10),
              constraints: BoxConstraints(minWidth: double.infinity),
              child: Row(
                children: <Widget>[
                  ColoredOrb(urgencyToColor(us.urgency)),
                  Expanded(child: Text(us.statement)),
                ],
              )),
          onPressed: () {
            Navigator.pop(context, us.urgency);
          },
        ),
      );
}
