import 'package:flutter/material.dart';
import 'screens/status_screen.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Hummingbird',
      theme: ThemeData(
          primaryColor: Colors.blue,
          brightness: Brightness.dark,
          primarySwatch: Colors.blue,
          backgroundColor: Color(0xFF1E2227),
          buttonTheme: ButtonThemeData(
            buttonColor: Colors.blue,
          )),
      home: StatusScreen(title: 'Hummingbird'),
    );
  }
}
