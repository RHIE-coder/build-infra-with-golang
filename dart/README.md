# Dart

## [ Install ]

 - ref: https://dart.dev/get-dart

```sh
sudo apt-get update
sudo apt-get install apt-transport-https
wget -qO- https://dl-ssl.google.com/linux/linux_signing_key.pub | sudo gpg --dearmor -o /usr/share/keyrings/dart.gpg
echo 'deb [signed-by=/usr/share/keyrings/dart.gpg arch=amd64] https://storage.googleapis.com/download.dartlang.org/linux/debian stable main' | sudo tee /etc/apt/sources.list.d/dart_stable.list

sudo apt-get update
sudo apt-get install dart
```

`~/.bashrc`

>```sh
>export PATH="$PATH:/usr/lib/dart/bin"
>```


## [ Hello World ]

```sh
dart create <project_name>
dart pub add uuid
```

OR

```sh
mkdir <project_name>
cd <project_name>
```

`pubspec.yaml`

```yaml
name: external_ethers
version: 1.0.0

environment:
  sdk: '>=2.19.6 <3.0.0'
```

`.gitignore`

```
.dart_tool
```