package envbinder

import (
	"fmt"
	"os"
	"strings"
)

func lookupString(name string) (s string, err error) {
	s, ok := os.LookupEnv(name)
	if !ok {
		return s, fmt.Errorf("No env var: %s", name)
	}
	return
}

func lookupStrings(name string, sep string) (slice []string, err error) {
	s, err := lookupString(name)
	slice = strings.Split(s, sep)
	return
}

func (b *EnvBinder) String(v *string, name string) *EnvBinder {
	s, err := lookupString(name)
	b.addError(err)
	*v = s
	return b
}

func (b *EnvBinder) StringOrDef(v *string, name string, def string) *EnvBinder {
	s, err := lookupString(name)
	if err != nil {
		*v = def
		return b
	}
	*v = s
	return b
}

func (b *EnvBinder) Strings(v *[]string, name string, sep string) *EnvBinder {
	s, err := lookupStrings(name, sep)
	b.addError(err)
	*v = s
	return b
}

func (b *EnvBinder) StringsOrDef(v *[]string, name string, sep string, def []string) *EnvBinder {
	s, err := lookupStrings(name, sep)
	if err != nil {
		*v = def
		return b
	}
	*v = s
	return b
}
