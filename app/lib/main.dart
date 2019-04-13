import 'package:flutter/material.dart';
import 'package:hummingbird/screens/chat.dart';
import 'package:hummingbird/widgets/text_switch.dart';
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
        home: TabbedHumm());
  }
}

class TabbedHumm extends StatefulWidget {
  @override
  State createState() => TabbedHummState();
}

class TabbedHummState extends State<TabbedHumm>
    with SingleTickerProviderStateMixin {
  TabController _tabController;

  @override
  void initState() {
    super.initState();
    _tabController = TabController(vsync: this, length: 2);
  }

  @override
  void dispose() {
    _tabController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) => Scaffold(
        appBar: AppBar(
          title: Text('Hummingbird'),
          actions: <Widget>[
            TextSwitch(
              tabController: _tabController,
            ),
          ],
        ),
        body: TabBarView(
          controller: _tabController,
          children: <Widget>[
            StatusScreen(title: 'Hummingbird'),
            Chat(),
          ],
        ),
      );
}
