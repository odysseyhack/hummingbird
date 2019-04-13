import 'package:flutter/material.dart';

class ColoredOrb extends StatelessWidget {
  final Color color;
  final EdgeInsets margin;

  ColoredOrb(this.color, {this.margin = const EdgeInsets.only(right: 10)});

  @override
  Widget build(BuildContext context) => Container(
        height: 13,
        width: 13,
        decoration: BoxDecoration(
          color: color,
          shape: BoxShape.circle,
        ),
        margin: margin,
      );
}
