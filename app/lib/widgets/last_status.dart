import 'dart:math';

import 'package:flutter/material.dart';

import 'package:hummingbird/generated/hummingbird.pb.dart';
import 'package:hummingbird/status_enum.dart';
import 'package:hummingbird/widgets/colored_orb.dart';

class LastStatus extends StatelessWidget {
  final Statuses statuses;
  final DateTime timestamp;

  LastStatus({Key key, this.statuses, this.timestamp}) : super(key: key);

  @override
  Widget build(BuildContext context) => Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: <Widget>[
          Text(
            "My last cast status",
            style: TextStyle(fontWeight: FontWeight.bold, fontSize: 20),
          ),
          Container(
            height: 60,
            child: Row(
              children: <Widget>[
                ColoredOrb(urgencyToColor(_getMostUrgent(statuses))),
                Expanded(
                  child: Text(
                    _statusesToText(_collectStatuses(statuses)),
                  ),
                ),
              ],
            ),
          ),
          RaisedButton(
            onPressed: () {},
            color: Colors.blue,
            child: Container(
                constraints: BoxConstraints(minWidth: double.infinity),
                child: Center(child: Text("Broadcast my wellbeing"))),
          ),
        ],
      );

  Urgency _getMostUrgent(Statuses status) {
    int highestUrgency = 0;
    for (var i = 0; i < Status.values.length; i++) {
      Urgency urgency = status.getField(i + 1) as Urgency;
      highestUrgency = max(highestUrgency, urgency.value);
    }

    return Urgency.valueOf(highestUrgency);
  }

  Map<Urgency, List<Status>> _collectStatuses(Statuses statuses) {
    var map = Map<Urgency, List<Status>>();

    for (var i = 0; i < Status.values.length; i++) {
      Urgency urgency = statuses.getFieldOrNull(i + 1) ?? Urgency.LOW;
      map.update(urgency, (list) => list..add(Status.values[i]),
          ifAbsent: () => [Status.values[i]]);
    }

    return map;
  }

  String _statusesToText(Map<Urgency, List<Status>> map) {
    var text = "";

    if (map.containsKey(Urgency.HIGH)) {
      text += "badly needs " + map[Urgency.HIGH].map(statusToString).join(", ");
    }

    if (map.containsKey(Urgency.MEDIUM)) {
      if (text.isNotEmpty) {
        text += ", ";
      }
      text += "needs " + map[Urgency.MEDIUM].map(statusToString).join(", ");
    }

    if (text.isEmpty) {
      text = "Doing okay";
    } else {
      text = text[0].toUpperCase() + text.substring(1);
    }

    return text;
  }
}
