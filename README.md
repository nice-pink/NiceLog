# What

NiceLog supports:

- Command line logging
- Remote logging

Enrich your logs with:

- Severity
- Timestamps
- Any additional maps
    - Add common data -> added to each log
    - Add custom map per log

# Local logging

Setup how your logs should be looking and start logging.

*Example*

```
# info log (INFO prefix)
log.Info("I just want to inform you ...")

# error log (red colored ERROR prefix)
log.Error("Trrrrrrouble!!!")

# success log (green colored INFO prefix)
log.Success("This was a great success!")

# enrich log with data
log.InfoD(map[string]string{"ping":"pong"}, "Ping is pong")
```

# Remote logging

Currently supports:
- Socket (tcp, udp)
- Http post
- NdJson

Connect the logger and start logging.

```
# connect to tcp socket
log.Connect(log.Connection{
    Address:    "http://localhost:3443",
    Protocol:   "tcp",
})
log.Info("info log")

# connect to ndjson compatible client (e.g. victoria logs)
log.Connect(log.Connection{
    Address:    "http://localhost:9428/insert/jsonline",
    Protocol:   "ndjson",
    StreamName: "test",
})
log.Info("info log")
```