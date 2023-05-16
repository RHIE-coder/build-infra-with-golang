```sh
$ useradd -m -d /home/testuser/ -s /bin/bash -G sudo testuser

-m creates the home directory if it does not exist.
-d overrides the default home directory location. 
-s sets the login shell for the user.
-G expects a comma-separated list of groups that the user should belong to.

$ passwd owen
$ passwd rhie
```

```sh
$ userdel -r rhie

-r remove home directory and mail spool
```

```
groupadd <name>
groupdel <name>

gpasswd -a <user> <group-name> # 추가
groups <user> # 확인
gpasswd -d <user> <group-name> # 제거
```

```sh
chown :mygroup -R /home/test
```