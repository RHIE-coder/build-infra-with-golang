import 'dart:math';

import 'package:http/http.dart';
import 'package:web3dart/crypto.dart';
import 'package:web3dart/web3dart.dart';
import 'package:dotenv/dotenv.dart';

Future<void> main() async {
  var env = DotEnv(includePlatformEnvironment: true)..load();
  String privKey = env['PRIVATE_KEY']!;
  Credentials credentials = EthPrivateKey.fromHex(privKey);
  print(credentials.address);
  String recipientAddress = '0x1Bc8D4d2A7069965CA0436667903aF4cf0f3A144';

  BigInt value = BigInt.from(1000000000000000); // 0.001 ETH in wei
  int gasPrice = (10 * pow(10, 9)).toInt(); // 1 Gwei in we
  int gasLimit = 21000;
  String hexNonce = '0x2';
  int nonce = int.parse(hexNonce.substring(2), radix: 16);

  Transaction transaction = Transaction(
    to: EthereumAddress.fromHex(recipientAddress),
    value: EtherAmount.inWei(value),
    gasPrice: EtherAmount.inWei(BigInt.from(gasPrice)),
    maxGas: gasLimit,
    nonce: nonce,
  );

  String rpcUrl = 'https://ropsten.infura.io/v3/YOUR_PROJECT_ID';
  final client = Web3Client(rpcUrl, Client());
  var result =
      await client.signTransaction(credentials, transaction, chainId: 11155111);
  print(result);
  print(bytesToHex(result));
  client.dispose();
}
