import 'package:flutter/material.dart';

import 'package:hummingbird/screens/chat_room.dart';

class Chat extends StatelessWidget {
  @override
  Widget build(BuildContext context) => GestureDetector(
        onTap: () => Navigator.push(
              context,
              MaterialPageRoute(builder: (context) => ChatRoom()),
            ),
        child: DecoratedBox(
            decoration: BoxDecoration(
                image: DecorationImage(image: AssetImage('assets/Chat.png')))),
      );
}
