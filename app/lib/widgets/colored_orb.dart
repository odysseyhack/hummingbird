import 'package:flutter/material.dart';

class ColoredOrb extends StatelessWidget {
  final Color color;

  ColoredOrb(this.color);

  @override
  Widget build(BuildContext context) => Container(
        height: 13,
        width: 13,
        decoration: BoxDecoration(
          color: color,
          shape: BoxShape.circle,
        ),
        margin: EdgeInsets.only(right: 10),
      );
}
