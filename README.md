# Words
A golang library to interface with system dictionaries and word lists.

### Word of caution
Right now the system has only been tested with MacOS, though according to [The Unix word dictionary](https://en.wikipedia.org/wiki/Words_(Unix))
this library should also function on other unix systems.

## Usage
There are 2 main functionalities in this library:
1. Parsing and loading the system's word lists into memory
```go
package main

import (
    "fmt"

    "github.com/kavfixnel/words"
)

func main() {
    // Get a map[string]struct{} object of all known system words
    wordMap, err := words.NewWordMap(nil)
    if err != nil {
        panic(err)
    }

    fmt.Println(wordMap)

    // Get a []string of all known system words
    wordList, err := words.NewWordList(nil)
    if err != nil {
        panic(err)
    }

    fmt.Println(wordList)
}
```

2. Check if words are valid and known by the system
```go
package main

import (
    "fmt"

    "github.com/kavfixnel/words"
)

func main() {
    mysteryWord := "abracadabra"
    isValid, err := words.IsValidWord(nil)
    if err != nil {
        panic(err)
    }

    modifier := ""
    if !isValid {
        modifier = " not"
    }

    fmt.Printf("%s is%s a valid word\n", mysteryWord, modifier)
}
```
