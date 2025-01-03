package matcher

var synonyms = map[string][]string{
	"java":   {"jvm", "jdk", "java programming"},
	"golang": {"go", "golang programming"},
}

func SynonymMatch(word string, keywords []string) bool {
	for _, keyword := range keywords {
		if contains(synonyms[keyword], word) {
			return true
		}
	}
	return false
}
