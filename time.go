package envbinder

import "time"

func lookupTime(name string, layout string) (t time.Time, err error) {
	s, err := lookupString(name)
	if err != nil {
		return
	}
	t, err = time.Parse(layout, s)
	return
}

func (b *EnvBinder) Time(v *time.Time, name string, layout string) *EnvBinder {
	t, err := lookupTime(name, layout)
	b.addError(err)
	*v = t
	return b
}

func (b *EnvBinder) TimeOrDef(v *time.Time, name string, layout string, def time.Time) *EnvBinder {
	t, err := lookupTime(name, layout)
	if err != nil {
		*v = def
		return b
	}
	*v = t
	return b
}

func (b *EnvBinder) Duration(v *time.Duration, name string) *EnvBinder {
	s, err := lookupString(name)
	if err != nil {
		b.addError(err)
		return b
	}
	d, err := time.ParseDuration(s)
	b.addError(err)
	*v = d
	return b
}

func (b *EnvBinder) DurationOrDef(v *time.Duration, name string, def time.Duration) *EnvBinder {
	s, err := lookupString(name)
	if err != nil {
		*v = def
		return b
	}
	d, err := time.ParseDuration(s)
	if err != nil {
		*v = def
		return b
	}
	*v = d
	return b
}
