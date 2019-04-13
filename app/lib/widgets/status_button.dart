import 'package:flutter/material.dart';

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
          onTap: () => handleToggle(_nextUrgency(urgency)),
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
                  ColoredOrb(_urgencyToColor(urgency)),
                  Text(
                    data.text,
                    style: TextStyle(color: Colors.white),
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
          const IconData(58940, fontFamily: 'MaterialIcons'),
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
          const IconData(58746, fontFamily: 'MaterialIcons'),
          "Protection",
        );
        break;
      case Status.SHELTER:
        return const Temp(
          const IconData(58746, fontFamily: 'MaterialIcons'),
          "Shelter",
        );
        break;
      case Status.TOOLS:
        return const Temp(
          const IconData(58746, fontFamily: 'MaterialIcons'),
          "Tools",
        );
        break;
      case Status.TRANSPORT:
        return const Temp(
          const IconData(58746, fontFamily: 'MaterialIcons'),
          "Transport",
        );
        break;
      case Status.WATER:
        return const Temp(
          const IconData(58746, fontFamily: 'MaterialIcons'),
          "Water",
        );
        break;
      default:
        throw "Unsupported case: " + status.toString();
    }
  }

  Color _urgencyToColor(Urgency urgency) {
    switch (urgency) {
      case Urgency.LOW:
        return const Color(0xFF00DF80);
        break;
      case Urgency.MEDIUM:
        return const Color(0xFFFFC125);
        break;
      case Urgency.HIGH:
        return const Color(0xFFFF004E);
      default:
        throw "No urgency to color case for " + urgency.toString();
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
