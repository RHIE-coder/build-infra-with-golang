# Flutter

## [ Install ]

 - ref: https://docs.flutter.dev/get-started/install

```sh
apt install cmake # dependency

wget https://storage.googleapis.com/flutter_infra_release/releases/stable/linux/flutter_linux_3.7.12-stable.tar.xz
tar xf ~/flutter_linux_3.7.12-stable.tar.xz
```

`~/.bashrc`

> ```sh
> export=$PATH:$HOME/flutter/bin
> ```

```sh
source ~/.bashrc
flutter precache
```

## [ Hello World ]

```sh
flutter create <project_name>
```

