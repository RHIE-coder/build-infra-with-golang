# P(rocess) M(anager) 2

## # 설치
```bash
npm i pm2
```

## # 주요 명령어

### - 실행 

```js
module.exports = {
  apps: [{ // apps은 고정
    name: 'app',
    script: './app.js',
    instances: 0, //instance 항목값을 '0'으로 설정하면 CPU 코어 수 만큼 프로세스를 생성하겠다는 뜻
    exec_mode: 'cluster'
  }]
}
```

```bash
npx pm2 start server.config.js # ecosystem.config.js (default)
pm2 reload all # Zero Downtime Reload

npx pm2 start npm --name "my-apps" -- run start
```

### - 조회 

```bash
npx pm2 ls
```

### - 로그

```sh
npx pm2 logs
```

```sh
$ pm2 logs APP-NAME       # Display APP-NAME logs
$ pm2 logs --json         # JSON output
$ pm2 logs --format       # Formated output

$ pm2 flush               # Flush all logs
$ pm2 reloadLogs          # Reload all logs
```

#### Log Rotate

```sh
npx pm2 install pm2-logrotate
```



### - 설정

#### `max_size` (Defaults to 10M): When a file size becomes higher than this value it will rotate it (its possible that the worker check the file after it actually pass the limit) . You can specify the unit at then end: 10G, 10M, 10K
#### `retain` (Defaults to 30 file logs): This number is the number of rotated logs that are keep at any one time, it means that if you have retain = 7 you will have at most 7 rotated logs and your current one.
#### `compress` (Defaults to false): Enable compression via gzip for all rotated logs
#### `dateFormat` (Defaults to YYYY-MM-DD_HH-mm-ss) : Format of the data used the name the file of log
#### `rotateModule` (Defaults to true) : Rotate the log of pm2's module like other apps
#### `workerInterval` (Defaults to 30 in secs) : You can control at which interval the worker is checking the log's size (minimum is 1)
#### `rotateInterval` (Defaults to 0 0 * * * everyday at midnight): This cron is used to a force rotate when executed. We are using node-schedule to schedule cron, so all valid cron for node-schedule is valid cron for this option. Cron style :
#### `TZ` (Defaults to system time): This is the standard tz database timezone used to offset the log file saved. For instance, a value of Etc/GMT+1, with an hourly log, will save a file at hour 14 GMT with hour 13 (GMT+1) in the log name.

```
*    *    *    *    *    *
┬    ┬    ┬    ┬    ┬    ┬
│    │    │    │    │    |
│    │    │    │    │    └ day of week (0 - 7) (0 or 7 is Sun)
│    │    │    │    └───── month (1 - 12)
│    │    │    └────────── day of month (1 - 31)
│    │    └─────────────── hour (0 - 23)
│    └──────────────────── minute (0 - 59)
└───────────────────────── second (0 - 59, OPTIONAL)
```

 - `pm2 set pm2-logrotate:max_size 1K (1KB)`
 - `pm2 set pm2-logrotate:compress true (compress logs when rotated)`
 - `pm2 set pm2-logrotate:rotateInterval '*/1 * * * *' (force rotate every minute)`

### - Start Up

```sh
# Generate Startup Script
$ pm2 startup

# Freeze your process list across server restart
$ pm2 save

# Remove Startup Script
$ pm2 unstartup
```

### - 모니터링

```sh
npx pm2 describe <id|app_name>
npx pm2 monit
```

### - 삭제

```sh
$ pm2 stop     <app_name|namespace|id|'all'|json_conf>
$ pm2 restart  <app_name|namespace|id|'all'|json_conf>
$ pm2 delete   <app_name|namespace|id|'all'|json_conf>
```

 - Managing Process

```sh
$ pm2 restart app_name
$ pm2 reload app_name
$ pm2 stop app_name
$ pm2 delete app_name
```

## # 주의해야할 점

각 코어마다의 Context는 공유가 안된다.

그러므로 메모리 위주로 작업을 할 때 주의해야 한다.

https://pm2.keymetrics.io/docs/usage/pm2-doc-single-page/