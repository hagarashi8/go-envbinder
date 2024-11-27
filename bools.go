package envbinder

import (
	"fmt"
	"slices"
	"strings"
)

var InvalidBoolValue = fmt.Errorf("Invalid Value for Bool")

var truthyStrings = []string{"yes", "enable", "enabled", "true"}
var falsyStrings = []string{"no", "disable", "disabled", "false"}

func lookupBool(name string) (bool, error) {
	s, err := lookupString(name)
	if err != nil {
		return false, err
	}
	s = strings.ToLower(s)
	return parseBool(s)
}

func parseBool(s string) (bool, error) {
	if slices.Contains(truthyStrings, s) {
		return true, nil
	} else if slices.Contains(falsyStrings, s) {
		return false, nil
	} else {
		return false, InvalidBoolValue
	}
}

func (b *EnvBinder) Bool(v *bool, name string) *EnvBinder {
	r, err := lookupBool(name)
	b.addError(err)
	*v = r
	return b
}

func (b *EnvBinder) BoolOrDef(v *bool, name string, def bool) *EnvBinder {
	r, err := lookupBool(name)
	if err != nil {
		*v = def
		return b
	}
	*v = r
	return b
}

func (b *EnvBinder) Bools(v *[]bool, name string, sep string) *EnvBinder {
	s, err := lookupStrings(name, sep)
	slice := []bool{}
	b.addError(err)
	for _, sbl := range s {
		bl, err := parseBool(sbl)
		if err != nil {
			b.addError(err)
			break
		}
		slice = append(slice, bl)
	}
	*v = slice
	return b
}

func (b *EnvBinder) BoolsOrDef(v *[]bool, name string, sep string, def []bool) *EnvBinder {
	s, err := lookupStrings(name, sep)
	slice := []bool{}
	b.addError(err)
	for _, sbl := range s {
		bl, err := parseBool(sbl)
		if err != nil {
			slice = def
			break
		}
		slice = append(slice, bl)
	}
	*v = slice
	return b
}
