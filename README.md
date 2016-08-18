Wanikani API wrapper written in Go.

Endpoints available: 

- [x] UserInformation
- [x] StudyQueue
- [x] LevelProgression
- [x] SRSDistribution
- [ ] RecentUnlocksList
- [ ] CriticalItemsList
- [ ] RadicalsList
- [ ] KanjiList
- [ ] VocabularyList

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
