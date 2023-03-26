# Words
A golang library to interface with system dictionaries and word lists.

[![Go Reference](https://pkg.go.dev/badge/github.com/kavfixnel/words.svg)](https://pkg.go.dev/github.com/kavfixnel/words)

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
    // Get a map[string]struct{} object of all words known by the system
    wordMap, err := words.NewWordMap(nil)
    if err != nil {
        panic(err)
    }

    fmt.Println(len(wordMap))

    // Get a []string of all known system words by the system
    wordList, err := words.NewWordList(nil)
    if err != nil {
        panic(err)
    }

    fmt.Println(len(wordList))
}
```
```bash
~/examples/words ❯ go run main.go
235976
235976
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
    isValid, err := words.IsValidWord(mysteryWord, nil)
    if err != nil {
        panic(err)
    }

    modifier := ""
    if !isValid {
        modifier = " not"
    }

    fmt.Printf("%s is%s a valid word\n", mysteryWord, modifier)

    // Ability to check variations of words
    mysteryWord = "ÄbraCADabra"
	isValid, err = words.IsValidWord(mysteryWord, &words.IsValidWordOptions{
		IgnoreCase:       true,
		IgnoreDiacritics: true,
	})
    if err != nil {
        panic(err)
    }

    modifier = ""
    if !isValid {
        modifier = " not"
    }

    fmt.Printf("%s is%s a valid word\n", mysteryWord, modifier)
}
```
```bash
~/examples/words ❯ go run main.go   
abracadabra is a valid word
ÄbraCADabra is a valid word
```
