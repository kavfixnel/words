package words

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Overwrite standard locations with local testing files
	unixStandardWordsFileLocations = []string{
		"tests/dictionary1",
	}
	localDictionaryLocations = []string{
		"tests/customWords",
	}

	code := m.Run()
	os.Exit(code)
}

func TestNewWordMap(t *testing.T) {
	t.Run("Get default map does not return", func(t *testing.T) {
		defaultLanguageMap, err := NewWordMap(nil)
		assert.NoError(t, err)

		assert.Len(t, defaultLanguageMap, 100)
	})

	t.Run("Default map does not include empty string", func(t *testing.T) {
		defaultLanguageMap, err := NewWordMap(nil)
		assert.NoError(t, err)

		assert.NotContains(t, defaultLanguageMap, "")
	})

	t.Run("Get language map with local dictionary", func(t *testing.T) {
		_, err := NewWordMap(&NewWordMapOptions{
			baseOptions{
				IncludeLocalDictionary: true,
			},
		})

		assert.NoError(t, err)
	})

	t.Run("Language map with local dictionary is larger than default", func(t *testing.T) {
		defaultLanguageMap, err := NewWordMap(nil)
		assert.NoError(t, err)

		languageMapWithLocalDictionary, err := NewWordMap(&NewWordMapOptions{
			baseOptions{
				IncludeLocalDictionary: true,
			},
		})
		assert.NoError(t, err)

		assert.Less(t, len(defaultLanguageMap), len(languageMapWithLocalDictionary))
	})

	t.Run("Language map with empty local dictionary is equal in size to default", func(t *testing.T) {
		localDictionaryLocations = []string{
			"tests/emptyCustomWords",
		}

		defaultLanguageMap, err := NewWordMap(nil)
		assert.NoError(t, err)

		languageMapWithLocalDictionary, err := NewWordMap(&NewWordMapOptions{
			baseOptions{
				IncludeLocalDictionary: true,
			},
		})
		assert.NoError(t, err)

		assert.Equal(t, len(defaultLanguageMap), len(languageMapWithLocalDictionary))
	})

	t.Run("Get default map with multiple standard word files does not return error", func(t *testing.T) {
		unixStandardWordsFileLocations = []string{
			"tests/dictionary1",
			"tests/dictionary2",
		}

		defaultLanguageMap, err := NewWordMap(nil)
		assert.NoError(t, err)

		assert.Len(t, defaultLanguageMap, 335)
	})

	t.Run("Additional word file gets included", func(t *testing.T) {
		defaultLanguageMap, err := NewWordMap(&NewWordMapOptions{
			baseOptions{
				AdditionalWordFiles: []string{
					"tests/additionalWordFile",
				},
			},
		})
		assert.NoError(t, err)

		assert.Contains(t, defaultLanguageMap, "additionalWord1")
	})

	t.Run("AdditionalWordFile that does not exist does not throw error", func(t *testing.T) {
		defaultLanguageMap, err := NewWordMap(&NewWordMapOptions{
			baseOptions{
				AdditionalWordFiles: []string{
					"tests/additionalWordFile",
					"tests/additionalWordFileDoesNotExist",
				},
			},
		})
		assert.NoError(t, err)

		assert.Contains(t, defaultLanguageMap, "additionalWord1")
	})
}

func TestNewWordList(t *testing.T) {
	// Overwrite standard locations with local testing files
	unixStandardWordsFileLocations = []string{
		"tests/dictionary1",
	}
	localDictionaryLocations = []string{
		"tests/customWords",
	}

	t.Run("Get default list does not return", func(t *testing.T) {
		defaultLanguageMap, err := NewWordList(nil)
		assert.NoError(t, err)

		assert.Len(t, defaultLanguageMap, 100)
	})

	t.Run("Default list does not include empty string", func(t *testing.T) {
		defaultLanguageMap, err := NewWordList(nil)
		assert.NoError(t, err)

		assert.NotContains(t, defaultLanguageMap, "")
	})

	t.Run("Get language list with local dictionary", func(t *testing.T) {
		_, err := NewWordList(&NewWordListOptions{
			baseOptions: baseOptions{
				IncludeLocalDictionary: true,
			},
			IgnoreSort: false,
		})

		assert.NoError(t, err)
	})

	t.Run("Language list with local dictionary is larger than default", func(t *testing.T) {
		defaultLanguageMap, err := NewWordList(nil)
		assert.NoError(t, err)

		languageMapWithLocalDictionary, err := NewWordList(&NewWordListOptions{
			baseOptions: baseOptions{
				IncludeLocalDictionary: true,
			},
			IgnoreSort: false,
		})
		assert.NoError(t, err)

		assert.Less(t, len(defaultLanguageMap), len(languageMapWithLocalDictionary))
	})

	t.Run("Language list with empty local dictionary is equal in size to default", func(t *testing.T) {
		localDictionaryLocations = []string{
			"tests/emptyCustomWords",
		}

		defaultLanguageMap, err := NewWordList(nil)
		assert.NoError(t, err)

		languageMapWithLocalDictionary, err := NewWordList(&NewWordListOptions{
			baseOptions: baseOptions{
				IncludeLocalDictionary: true,
			},
			IgnoreSort: false,
		})
		assert.NoError(t, err)

		assert.Equal(t, len(defaultLanguageMap), len(languageMapWithLocalDictionary))
	})

	t.Run("Get default list with multiple standard word files does not return error", func(t *testing.T) {
		unixStandardWordsFileLocations = []string{
			"tests/dictionary1",
			"tests/dictionary2",
		}

		defaultLanguageMap, err := NewWordList(nil)
		assert.NoError(t, err)

		assert.Len(t, defaultLanguageMap, 335)
	})

	t.Run("Additional word file gets included", func(t *testing.T) {
		defaultLanguageMap, err := NewWordList(&NewWordListOptions{
			baseOptions: baseOptions{
				AdditionalWordFiles: []string{
					"tests/additionalWordFile",
				},
			},
		})
		assert.NoError(t, err)

		assert.Contains(t, defaultLanguageMap, "additionalWord1")
	})
}

func TestIsValidWord(t *testing.T) {
	t.Run("A valid word returns true with nil Options", func(t *testing.T) {
		isValid, err := IsValidWord("Abba", nil)
		assert.NoError(t, err)

		assert.True(t, isValid)
	})

	t.Run("A valid word returns true", func(t *testing.T) {
		isValid, err := IsValidWord("Abba", &IsValidWordOptions{
			baseOptions: baseOptions{
				IncludeLocalDictionary: false,
			},
			IgnoreCase:       false,
			IgnoreDiacritics: false,
		})
		assert.NoError(t, err)

		assert.True(t, isValid)
	})

	t.Run("An invalid word returns false", func(t *testing.T) {
		isValid, err := IsValidWord("jhdsoagnbeiv", &IsValidWordOptions{
			baseOptions: baseOptions{
				IncludeLocalDictionary: false,
			},
			IgnoreCase:       false,
			IgnoreDiacritics: false,
		})
		assert.NoError(t, err)

		assert.False(t, isValid)
	})

	t.Run("Word variations", func(t *testing.T) {
		t.Run("A valid, but differently capitalized word returns true", func(t *testing.T) {
			isValid, err := IsValidWord("ABBA", &IsValidWordOptions{
				baseOptions: baseOptions{
					IncludeLocalDictionary: false,
				},
				IgnoreCase:       true,
				IgnoreDiacritics: false,
			})
			assert.NoError(t, err)

			assert.True(t, isValid)
		})

		t.Run("A valid word, with a diacritical symbol returns true", func(t *testing.T) {
			isValid, err := IsValidWord("Äbba", &IsValidWordOptions{
				baseOptions: baseOptions{
					IncludeLocalDictionary: false,
				},
				IgnoreCase:       false,
				IgnoreDiacritics: true,
			})
			assert.NoError(t, err)

			assert.True(t, isValid)
		})
	})

	t.Run("Additional word file gets included", func(t *testing.T) {
		defaultLanguageMap, err := IsValidWord("additionalWord1", &IsValidWordOptions{
			baseOptions: baseOptions{
				AdditionalWordFiles: []string{
					"tests/additionalWordFile",
				},
			},
		})
		assert.NoError(t, err)

		assert.True(t, defaultLanguageMap)
	})

	t.Run("Local Dictionary gets included and used", func(t *testing.T) {
		localDictionaryLocations = []string{"tests/customWords"}

		defaultLanguageMap, err := IsValidWord("isthisarealword", &IsValidWordOptions{
			baseOptions: baseOptions{
				IncludeLocalDictionary: true,
			},
		})
		assert.NoError(t, err)

		assert.True(t, defaultLanguageMap)
	})

	t.Run("Local Dictionary gets included and used", func(t *testing.T) {
		isValidWord, err := IsValidWord("thisisnotavalidword", &IsValidWordOptions{
			baseOptions: baseOptions{
				AdditionalWordFiles: []string{
					"tests/additionalWordFileDoesNotExist",
				},
			},
		})
		assert.NoError(t, err)

		assert.False(t, isValidWord)
	})
}
