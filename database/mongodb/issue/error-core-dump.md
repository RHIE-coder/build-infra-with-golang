# ISSUE: Mongod fail with `core-dump`

```sh
mongod.service - MongoDB Database Server
   Loaded: loaded (/lib/systemd/system/mongod.service; disabled; vendor preset: enabled)
   Active: failed (Result: core-dump) since Tue 2022-09-06 23:25:06 KST; 42min ago
     Docs: https://docs.mongodb.org/manual
 Main PID: 26775 (code=dumped, signal=ILL)

 9월 06 23:25:06 ksrhie-VirtualBox systemd[1]: Started MongoDB Database Server.
 9월 06 23:25:06 ksrhie-VirtualBox systemd[1]: mongod.service: Main process exited, code=dumped, status=4/ILL
 9월 06 23:25:06 ksrhie-VirtualBox systemd[1]: mongod.service: Failed with result 'core-dump'.
```

## ! 원인

https://www.mongodb.com/community/forums/t/status-of-mongod-returns-core-dump/13076

 - `/proc/cpuinfo`에서 `avx` 및 `avx2`가 `grep`되어야 함.
 - 만일 없다면 CPU가 version 5를 지원하지 않음

```txt
For anyone else with this problem, try running the following:

grep avx /proc/cpuinfo
grep avx2 /proc/cpuinfo
If neither print anything to the console, your CPU doesn’t support v5.

You can install version 4.4 from the Version 4.4 MongoDB docs at: https://docs.mongodb.com/v4.4/tutorial/install-mongodb-on-ubuntu/

Thank you again.
```