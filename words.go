// Package words provides a way to read system word lists
package words

import (
	"bufio"
	"io"
	"os"
	"sort"

	"golang.org/x/text/language"
	"golang.org/x/text/search"
)

var (
	// Unix standard words files
	// https://en.wikipedia.org/wiki/Words_(Unix)
	unixStandardWordsFileLocations = []string{
		"/usr/share/dict/words", // Location in MacOS
		"/usr/dict/words",       // Location in other linux distributions
	}

	// https://superuser.com/a/136267
	localDictionaryLocations = []string{
		"~/Library/Spelling/LocalDictionary", // Location local Dictionary in MacOS
	}
)

// NewWordMapOptions configures the way the word list and maps should be contructed.
type baseOptions struct {
	// IncludeLocalDictionary tells the library if it should include
	// user defined words
	IncludeLocalDictionary bool `default:"false"`

	// AdditionalWordFiles gives you the ability to provide paths to
	// extra word files.
	AdditionalWordFiles []string `default:"[]"`
}

// NewWordMapOptions configures the way the word list and maps should be contructed.
type NewWordMapOptions struct {
	baseOptions
}

func parseWordList(file io.Reader, wordMap *map[string]struct{}) (err error) {
	// Read the file line by line and populate the wordMap map
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		(*wordMap)[scanner.Text()] = struct{}{}
	}
	err = scanner.Err()

	return
}

// NewWordMap reads the systems word list(s) and returns them.
// It returns the map of system words and any errors encountered.
func NewWordMap(options *NewWordMapOptions) (map[string]struct{}, error) {
	if options == nil {
		options = &NewWordMapOptions{
			baseOptions{
				IncludeLocalDictionary: false,
			},
		}
	}

	wordListLocations := append(unixStandardWordsFileLocations, options.AdditionalWordFiles...)
	if options.IncludeLocalDictionary {
		wordListLocations = append(wordListLocations, localDictionaryLocations...)
	}

	wordMap := make(map[string]struct{})

	for _, wordListLocation := range wordListLocations {
		_, err := os.Stat(wordListLocation)
		if err != nil {
			continue
		}

		wordsFile, err := os.Open(wordListLocation)
		if err != nil {
			return map[string]struct{}{}, err
		}

		err = parseWordList(wordsFile, &wordMap)
		if err != nil {
			return map[string]struct{}{}, err
		}
	}

	return wordMap, nil
}

// NewWordListOptions configures the way the word list and maps should be contructed.
type NewWordListOptions struct {
	baseOptions

	// IgnoreSort defines if the function should sort the list or merge the files of
	// words in the way the happen to appear.
	IgnoreSort bool `default:"false"`
}

// NewWordList reads the systems word list(s) and returns them.
// It returns the sorted list of unique system words and any errors encountered.
func NewWordList(options *NewWordListOptions) ([]string, error) {
	if options == nil {
		options = &NewWordListOptions{
			baseOptions: baseOptions{
				IncludeLocalDictionary: false,
			},
			IgnoreSort: false,
		}
	}

	wordMap, err := NewWordMap(&NewWordMapOptions{
		baseOptions: options.baseOptions,
	})
	if err != nil {
		return []string{}, err
	}

	wordList := make([]string, 0, len(wordMap))
	for word := range wordMap {
		wordList = append(wordList, word)
	}

	if !options.IgnoreSort {
		sort.Strings(wordList)
	}

	return wordList, nil
}

// IsValidWordOptions defines the parameters of the IsValidWord function.
type IsValidWordOptions struct {
	baseOptions

	// All string comparisons use the golang.org/x/text/search#Matcher.EqualString comparator

	// IgnoreCase tells the library if it should deduplicate words that are
	// are eqivalent with respect to case ("Hello" == "hello").
	IgnoreCase bool `default:"false"`

	// IgnoreDiacritics tells the library if it should deduplicate words that are
	// are eqivalent with respect to diacritics ("Aö" == "Ao").
	IgnoreDiacritics bool `default:"false"`

	// Language specifies the language of the system and it's word lists.
	Language language.Tag `default:"language.English"`
}

// IsValidWord takes a word and comparison options and tells you if given word is
// a valid word according to the local word lists.
// It returns the validness of the word and any errors encountered.
func IsValidWord(word string, options *IsValidWordOptions) (bool, error) {
	if options == nil {
		options = &IsValidWordOptions{
			IgnoreCase:       false,
			IgnoreDiacritics: false,
			Language:         language.English,
		}
	}

	wordListLocations := append(unixStandardWordsFileLocations, options.AdditionalWordFiles...)
	if options.IncludeLocalDictionary {
		wordListLocations = append(wordListLocations, localDictionaryLocations...)
	}

	matcherOptions := []search.Option{}
	if options.IgnoreCase {
		matcherOptions = append(matcherOptions, search.IgnoreCase)
	}
	if options.IgnoreDiacritics {
		matcherOptions = append(matcherOptions, search.IgnoreDiacritics)
	}
	matcher := search.New(options.Language, matcherOptions...)

	for _, wordListLocation := range wordListLocations {
		_, err := os.Stat(wordListLocation)
		if err != nil {
			continue
		}

		wordsFile, err := os.Open(wordListLocation)
		if err != nil {
			return false, err
		}

		scanner := bufio.NewScanner(wordsFile)
		for scanner.Scan() {
			if matcher.EqualString(word, scanner.Text()) {
				return true, nil
			}
		}

		if err = scanner.Err(); err != nil {
			return false, err
		}
	}

	return false, nil
}
