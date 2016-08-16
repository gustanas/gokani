# Example

``` go
package main

import (
	"fmt"
	"github.com/gustanas/gokani"
)

func main() {
	c := gokani.Client{Key: "MY_API_KEY"}

	fmt.Println(c.UserInformation())
	fmt.Println(c.StudyQueue())
}
```
