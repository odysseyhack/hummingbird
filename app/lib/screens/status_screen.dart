import 'dart:async';

import 'package:flutter/material.dart';

import 'package:hummingbird/generated/hummingbird.pb.dart';
import 'package:hummingbird/generated/hummingbird.pbgrpc.dart';
import 'package:hummingbird/widgets/last_status.dart';
import 'package:hummingbird/widgets/status_button.dart';
import 'package:hummingbird/status_enum.dart';
import 'package:hummingbird/widgets/text_switch.dart';

class StatusScreen extends StatefulWidget {
  StatusScreen({Key key, this.title}) : super(key: key);

  final String title;

  @override
  _StatusScreenState createState() => _StatusScreenState();
}

class _StatusScreenState extends State<StatusScreen> {
  Statuses statuses = Statuses()
    ..wATER = Urgency.LOW
    ..fIRSTAID = Urgency.LOW
    ..fOOD = Urgency.LOW
    ..sHELTER = Urgency.LOW
    ..pROTECTION = Urgency.LOW
    ..fIRE = Urgency.LOW
    ..eLECTRICITY = Urgency.LOW
    ..tOOLS = Urgency.LOW
    ..tRANSPORT = Urgency.LOW;

  @override
  Widget build(BuildContext context) {
    return  Container(
        color: Color(0xFF1E2227),
        child: Column(
          children: <Widget>[
            Expanded(
              child: Column(
                children: <Widget>[
                  Container(
                    padding: EdgeInsets.all(15),
                    child: Text(
                      'Tap icons to cast your needs in your status.',
                    ),
                  ),
                  Container(
                    color: Color(0x66FFFFFF),
                    child: GridView.count(
                      crossAxisCount: 3,
                      shrinkWrap: true,
                      crossAxisSpacing: 1,
                      mainAxisSpacing: 1,
                      children: _buildStatuses(context),
                    ),
                  ),
                ],
              ),
            ),
            Container(
              constraints: BoxConstraints(minWidth: double.infinity),
              height: 1,
              color: Color(0x44FFFFFF),
            ),
            Container(
              constraints: const BoxConstraints(minWidth: double.infinity),
              padding: EdgeInsets.symmetric(horizontal: 40, vertical: 8),
              color: Colors.black,
              child: LastStatus(
                statuses: statuses,
              ),
            ),
          ],
        ),
      );
  }

  List<StatusButton> _buildStatuses(BuildContext context) {
    var buttons = List<StatusButton>();
    for (var i = 0; i < Status.values.length; i++) {
      buttons.add(StatusButton(
        status: Status.values[i],
        urgency: statuses.getFieldOrNull(i + 1) ?? Urgency.LOW,
        handleToggle: (Urgency urgency) {
          setState(() => statuses.setField(i + 1, urgency));
        },
      ));
    }
    return buttons;
  }
}
