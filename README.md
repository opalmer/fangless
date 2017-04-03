# Fangless

Fangless is a small wrapper around [spf13's viper](https://github.com/spf13/viper)
package.

## Usage
### Download

This package can be downloaded in standard go fashion
or vendored using tags.

```console
go get github.com/opalmer/fangless
```

### API Usage
```go
package main

import (
	"github.com/opalmer/fangless"
	"github.com/spf13/viper"
)

func main() {
	snake := fangless.New(viper.New())
}
```

The `snake` variable will provide access to Viper's
own read functions with read-only locks. You may also
call `snake.Set` or `snake.SetDefault` to set a value
safely inside of a goroutine. If additional control
is needed direct access is allowed too:

```go
package main

import (
	"github.com/opalmer/fangless"
	"github.com/spf13/viper"
)

func main() {
	snake := fangless.New(viper.New())
	v := snake.Viper() // locks the config
	defer snake.Unlock()
	// do stuff with v
	
}
```