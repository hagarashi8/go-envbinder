package envbinder

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var InvalidBoolValue = fmt.Errorf("Invalid Value for Bool")

type EnvBinder struct {
	errs []error
}

func NewEnvBinder() (b EnvBinder) {
	b.errs = make([]error, 0)
	return
}

func (b *EnvBinder) addError(err error) {
	b.errs = append(b.errs, err)
}

func (b *EnvBinder) lookupString(name string) (s string, err error) {
	s, ok := os.LookupEnv(name)
	if !ok {
		return s, fmt.Errorf("No env var: %s", name)
	}
	return
}

func (b *EnvBinder) String(v *string, name string) *EnvBinder {
	s, err := b.lookupString(name)
	b.addError(err)
	*v = s
	return b
}

func (b *EnvBinder) StringOrDef(v *string, name string, def string) *EnvBinder {
	s, err := b.lookupString(name)
	if err != nil {
		*v = def
		return b
	}
	*v = s
	return b
}

func (b *EnvBinder) lookupBool(name string) (bool, error) {
	s, err := b.lookupString(name)
	if err != nil {
		return false, err
	}
	s = strings.ToLower(s)

	if slices.Contains([]string{"yes", "enable", "enabled", "true"}, s) {
		return true, nil
	} else if slices.Contains([]string{"no", "disable", "disabled", "false"}, s) {
		return false, nil
	} else {
		return false, InvalidBoolValue
	}
}

func (b *EnvBinder) Bool(v *bool, name string) *EnvBinder {
	r, err := b.lookupBool(name)
	*v = r
	b.addError(err)
	return b
}

func (b *EnvBinder) BoolOrDef(v *bool, name string, def bool) *EnvBinder {
	r, err := b.lookupBool(name)
	if err != nil {
		*v = def
		return b
	}
	*v = r
	return b
}

func (b *EnvBinder) lookupInt(name string) (int, error) {
	s, err := b.lookupString(name)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(s)
}

func (b *EnvBinder) Int(v *int, name string) *EnvBinder {
	i, err := b.lookupInt(name)
	if err != nil {
		b.addError(err)
		return b
	}
	*v = i
	return b
}

func (b *EnvBinder) IntOrDef(v *int, name string, def int) *EnvBinder {
	i, err := b.lookupInt(name)
	if err != nil {
		*v = def
		return b
	}
	*v = i
	return b
}

func (b *EnvBinder) BindError() error {
	if len(b.errs) > 0 {
		b.errs = b.errs[1:]
		return b.errs[0]
	}
	return nil
}

func (b *EnvBinder) BindErrors() []error {
	b.errs = make([]error, 0)
	return b.errs
}
