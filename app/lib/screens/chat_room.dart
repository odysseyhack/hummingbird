import 'package:flutter/material.dart';

class ChatRoom extends StatefulWidget {
  @override
  State createState() => ChatRoomState();
}

class ChatRoomState extends State<ChatRoom> {
  @override
  Widget build(BuildContext context) => Scaffold(
        appBar: AppBar(
          title: Text("Duncan Grint"),
        ),
        body: Container(
          constraints: BoxConstraints.expand(),
          decoration: BoxDecoration(
              gradient: LinearGradient(
                  colors: [Color(0xFF0072FF), Color(0xFF00C3FF)])),
          child: Stack(
            children: <Widget>[
              Positioned(
                bottom: 20,
                left: 20,
                right: 20,
                child: TextField(
                  style: TextStyle(color: Colors.black),
                  decoration: InputDecoration(
                    filled: true,
                    fillColor: Colors.white,
                    border: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(5),
                    ),
                    hintText: 'Write a message',
                    hintStyle: TextStyle(color: Color(0xFFCECFD0)),
                  ),
                ),
              )
            ],
          ),
        ),
      );
}
