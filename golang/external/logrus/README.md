# install

```sh
go get -u github.com/sirupsen/logrus
```

# api

```go
log.SetFormatter(&log.TextFormatter{}) // default
log.SetFormatter(&log.TextFormatter{
    DisableColors: true,
    FullTimestamp: true,
    TimestampFormat: time.RFC3339,
})
log.SetFormatter(&log.JSONFormatter{})

// Output to stdout instead of the default stderr
// Can be any io.Writer, see below for File example
log.SetOutput(os.Stdout)

// Only log the warning severity or above.
log.SetLevel(log.DebugLevel)
```

```go
log.Printf("hello")
log.WithFields(log.Fields{
    "animal": "walrus",
}).Info("A walrus appears")
// INFO[2023-05-31T09:40:05+09:00] A walrus appears                              animal=walrus
```