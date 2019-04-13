import 'dart:ui';

import 'package:hummingbird/generated/hummingbird.pbenum.dart';

enum Status {
  WATER,
  FIRST_AID,
  FOOD,
  SHELTER,
  PROTECTION,
  FIRE,
  ELECTRICITY,
}

Color urgencyToColor(Urgency urgency) {
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

String statusToString(Status status) {
  switch (status) {
    case Status.ELECTRICITY:
      return "Electricity";
      break;
    case Status.FIRE:
      return "Fire";
      break;
    case Status.FIRST_AID:
      return "First aid";
      break;
    case Status.FOOD:
      return "Food";
      break;
    case Status.PROTECTION:
      return "Heating";
      break;
    case Status.WATER:
      return "Water";
      break;
    case Status.SHELTER:
      return "Building Collapse";
    default:
      throw "Unsupported case: " + status.toString();
  }
}

List<UrgencyStatement> statusToUrgencyStatements(Status status) {
  var list = List<UrgencyStatement>();

  switch (status) {
    case Status.WATER:
      list.add(UrgencyStatement(
          Urgency.MEDIUM, "I haven't been able to drink in the last 6h"));
      list.add(UrgencyStatement(Urgency.HIGH, "Super thirsty"));
      break;
    case Status.SHELTER:
      list..add(UrgencyStatement(Urgency.LOW, "Nothing out of the ordinary"));
      list
        ..add(UrgencyStatement(Urgency.MEDIUM, "A building near me collapsed"));
      list
        ..add(UrgencyStatement(Urgency.HIGH,
            "People are buried underneath a nearby collapsed building"));
      break;
    default:
      list
        ..add(UrgencyStatement(Urgency.LOW, "I'm okay."))
        ..add(UrgencyStatement(Urgency.MEDIUM, "Please help me out"))
        ..add(UrgencyStatement(Urgency.HIGH, "For the love of god!"));
      break;
  }
  return list;
}

class UrgencyStatement {
  final String statement;
  final Urgency urgency;

  UrgencyStatement(this.urgency, this.statement);
}
