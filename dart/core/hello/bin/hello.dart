import 'package:uuid/uuid.dart';
import 'package:hello/hello.dart' as hello;

void main(List<String> arguments) {
  print('Hello world: ${hello.calculate()}!');
  var uuid = Uuid();
  print(uuid.v4());
}
