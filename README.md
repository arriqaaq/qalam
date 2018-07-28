# qalam
Simple time based filewriter to log data to disk. Has buffer support to flush data from memory to disk to reduce disk IO


Features
--------
- Create a simple log file handler, will do log rotation automatically

Installing
----------

```
go get -u github.com/arriqaaq/qalam
```

Example
-------

Here's a full example of a qalam that writes to a file with a (%Y%m%d) format:

You can run this example from a terminal:

```sh
go run example/main.go
```

```go
package main

import (
	"github.com/arriqaaq/qalam"
	"log"
	"time"
)

func main() {
	config := qalam.NewConfig("./log.%Y%m%d.gz", time.Local, 1, true, 10*time.Millisecond)
	c, err := qalam.NewQalam(config)
	if err != nil {
		log.Fatalln(err)
	}
	c.Writeln([]byte("pogba"))
	c.Writeln([]byte("kante"))
	c.Close()
}

```

TODO
--------
- Add more test cases
- Add prometheus stats
- Add zap logger for debug purpose

References
--------
- https://www.devdungeon.com/content/working-files-go

Contact
-------
Farhan [@arriqaaq](http://twitter.com/arriqaaq)
