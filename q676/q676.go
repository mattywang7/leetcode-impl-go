package q676

type MagicDictionary []string

func Constructor() MagicDictionary {
	return MagicDictionary{}
}

func (m *MagicDictionary) BuildDict(dictionary []string) {
	*m = dictionary
}

func (m *MagicDictionary) Search(searchWord string) bool {
next:
	for _, word := range *m {
		if len(word) != len(searchWord) {
			continue
		}
		diff := false
		for i := range word {
			if word[i] != searchWord[i] {
				if diff {
					continue next
				}
				diff = true
			}
		}
		if diff {
			return true
		}
	}
	return false
}
