# Gokani (Work in Progress)

[Wanikani](https://www.wanikani.com/) API wrapper written in Go.

Endpoints available: 

- UserInformation
- StudyQueue
- LevelProgression
- SRSDistribution
- RecentUnlocksList
- CriticalItemsList
- RadicalsList
- KanjiList
- VocabularyList

# Example

``` go
package main

import (
	"fmt"
	"github.com/gustanas/gokani"
)

func main() {
	c := gokani.Client{Key: "MY_API_KEY"}

	levels := []string{"1", "2"}
	v, e := c.KanjiList(levels)

	if e != nil {
		// handle error
	}
	fmt.Println(v[0].Kunyomi)
}
```
