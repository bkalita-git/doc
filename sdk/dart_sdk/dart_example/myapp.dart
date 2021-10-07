//dart compile exe myapp.dart -o myapp


import 'dart:async';

import 'dart:io';

void main(){
 HttpServer.bind('127.0.0.1', 4444)
     .then((server) => print('${server.isBroadcast}'))
     .catchError(print);
}
