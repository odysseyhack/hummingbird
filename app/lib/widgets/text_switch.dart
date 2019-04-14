import 'package:flutter/material.dart';
import 'package:hummingbird/widgets/colored_orb.dart';

class TextSwitch extends StatefulWidget {
  final bool value;
  final ValueChanged<bool> onToggle;
  final TabController tabController;

  TextSwitch({Key key, this.value, this.onToggle, @required this.tabController})
      : super(key: key);

  @override
  State createState() => TextSwitchState();
}

class TextSwitchState extends State<TextSwitch> {
  @override
  Widget build(BuildContext context) => GestureDetector(
        onTap: () {
          var i = widget.tabController.index == 0 ? 1 : 0;
          widget.tabController.animateTo(i);
          setState(() {});
        },
        child: Container(
          constraints: BoxConstraints(),
          padding: EdgeInsets.all(5),
          width: 85,
          decoration: BoxDecoration(
              color: Colors.black, borderRadius: BorderRadius.circular(50)),
          child: widget.tabController.index == 1
              ? _buildChatRow()
              : _buildHiberRow(),
        ),
      );

  Widget _buildChatRow() => Row(
        mainAxisAlignment: MainAxisAlignment.end,
        mainAxisSize: MainAxisSize.min,
        children: <Widget>[
          Text("Chat"),
          ColoredOrb(
            Colors.white,
            margin: EdgeInsets.only(left: 5),
          ),
        ],
      );

  Widget _buildHiberRow() => Row(
        mainAxisAlignment: MainAxisAlignment.start,
        mainAxisSize: MainAxisSize.min,
        children: <Widget>[
          ColoredOrb(Colors.white, margin: EdgeInsets.only(right: 10)),
          Text("Hiber"),
        ],
      );
}
