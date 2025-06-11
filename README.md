# Server Module

⚡ For geekros server module ⚡

## License

[![License:Apache2.0](https://img.shields.io/badge/License-Apache2.0-yellow.svg)](https://opensource.org/licenses/Apache2.0)

## Example

```golang
package service

import "github.com/geekros/server/pkg/server"

var Get = &server.Server{}
```

```golang
...
service.Get = server.New()
service.Get.Start(port, mode, ReadTimeout, WriteTimeout, func() {
    log.Println(color.Gray.Text("started server"))
}, func() {
    log.Println(color.Gray.Text("exited server"))
})
...
```