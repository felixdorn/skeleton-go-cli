# Store

Store provides a simple interface to store and retrieve configuration or data created or stored by `:bin` (by default, store uses `~/.:bin` as the root directory for storing data).

## Usage

```go
package mydomain

import (
	"fmt"
	"github.com/owner/repository/core/domain/store"
)

const Store = store.Store("dirname-unique-across-app")

func Do() error {
	dir, err := Store.Dir()
	if err != nil {
		return err
	}

	fmt.Println(dir)
	// Output: /home/user/.:bin/dirname-unique-across-app
	
	file, err := Store.Open("filename", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
        return err
    }
	
	// do something with the file (*os.File)
	
	
	return nil
}
```