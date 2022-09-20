# Log

Log provides a `Log` function to log message in `~/.:bin/:bin.log` (by default).

## Usage

```go
package mydomain

import (
	"github.com/owner/repository/core/domain/log"
	"github.com/vite-cloud/go-zoup"
)

func Do() error {
	log.Log(zoup.DebugLevel, "message", zoup.Fields{
		"key": "value",
    })
	
	return nil
}
```