import 'package:flutter/material.dart';
import 'package:hummingbird/screens/status_screen.dart';
import 'package:hummingbird/screens/urgency_selection.dart';

import 'package:hummingbird/status_enum.dart';
import 'package:hummingbird/generated/hummingbird.pbenum.dart';
import 'package:hummingbird/widgets/colored_orb.dart';

class StatusButton extends StatelessWidget {
  final Status status;
  final Urgency urgency;
  final ValueChanged<Urgency> handleToggle;

  StatusButton(
      {Key key,
      @required this.status,
      @required this.urgency,
      this.handleToggle})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    Temp data = _statusToIcon(status);
    return Container(
      decoration: BoxDecoration(
        gradient: LinearGradient(
            begin: Alignment.topCenter,
            end: Alignment.bottomCenter,
            colors: [Color(0x88282C30), Color(0xFF1E2227)]),
      ),
      child: Material(
        color: Colors.transparent,
        child: InkWell(
          onTap: () async {
            final newUrgency = await Navigator.push(
                    context,
                    MaterialPageRoute(
                        builder: (context) => UrgencySelection(status))) ??
                urgency;

            handleToggle(newUrgency);
          },
          child: Column(
            mainAxisAlignment: MainAxisAlignment.spaceEvenly,
            children: <Widget>[
              Icon(
                data.iconData,
                color: Colors.white,
                size: 30,
              ),
              Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: <Widget>[
                  ColoredOrb(urgencyToColor(urgency)),
                  Container(
                    constraints: BoxConstraints(maxWidth: 80),
                    child: Text(
                      data.text,
                      style: TextStyle(color: Colors.white),
                    ),
                  ),
                ],
              )
            ],
          ),
        ),
      ),
    );
  }

  Temp _statusToIcon(Status status) {
    switch (status) {
      case Status.ELECTRICITY:
        return const Temp(
          const IconData(59698, fontFamily: 'MaterialIcons'),
          "Electricity",
        );
        break;
      case Status.FIRE:
        return const Temp(
          const IconData(58416, fontFamily: 'MaterialIcons'),
          "Fire",
        );
        break;
      case Status.FIRST_AID:
        return const Temp(
          const IconData(57669, fontFamily: 'MaterialIcons'),
          "First aid",
        );
        break;
      case Status.FOOD:
        return const Temp(
          const IconData(58746, fontFamily: 'MaterialIcons'),
          "Food",
        );
        break;
      case Status.PROTECTION:
        return const Temp(
          const IconData(58416, fontFamily: 'MaterialIcons'),
          "Heating",
        );
        break;
      case Status.SHELTER:
        return const Temp(
          const IconData(57393, fontFamily: 'MaterialIcons'),
          "Building Collapse",
        );
        break;
      case Status.WATER:
        return const Temp(
          const IconData(59537, fontFamily: 'MaterialIcons'),
          "Water",
        );
        break;
      default:
        throw "Unsupported case: " + status.toString();
    }
  }

  Urgency _nextUrgency(Urgency urgency) {
    switch (urgency) {
      case Urgency.LOW:
        return Urgency.MEDIUM;
        break;
      case Urgency.MEDIUM:
        return Urgency.HIGH;
        break;
      case Urgency.HIGH:
        return Urgency.LOW;
        break;
      default:
        throw "Uncovered urgency case " + urgency.toString();
    }
  }
}

class Temp {
  final IconData iconData;
  final String text;

  const Temp(this.iconData, this.text);
}
