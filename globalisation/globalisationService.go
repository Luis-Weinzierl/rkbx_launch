package globalisation

///
/// This package offers a fallback mechanism as an extension of fyne's
/// built in translation system. It also allows for an ID-based text system
/// without any hard-coded strings in the code
///

import (
	"encoding/json"

	"fyne.io/fyne/v2/lang"
)

var defaultLanguage = map[string]string{}

func Get(id string) string {
	return lang.X(id, defaultLanguage[id])
}

func LoadDefaultLanguage(jsonStr []byte) {
	json.Unmarshal(jsonStr, &defaultLanguage)
}
