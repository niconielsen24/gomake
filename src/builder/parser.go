package builder

import (
	"errors"
	"strings"
)

var (
	ErrCouldNotFindLang = errors.New("Could not find given language")
)

var lang_map = map[string]string{
	"Go":     "go",
	"Golang": "go",
	"Python": "python",
	"Zig":    "zig",
}

func ParseArguments(s string) (*BuildArguments, error) {
	ba := BuildArguments{}

	arguments := strings.Split(s, ":")
	if len(arguments) > 1 {
		if l, exists := lang_map[arguments[0]]; exists {
			ba.SetLang(l)
      techs := strings.Split(arguments[1], ",") 
       
		} else {
      return nil, ErrCouldNotFindLang
    }
	} else {
		if l, exists := lang_map[arguments[0]]; exists {
			ba.SetLang(l)
		} else {
			return nil, ErrCouldNotFindLang
		}
	}

	return &ba, nil
}
