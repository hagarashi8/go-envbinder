package envbinder

type EnvBinder struct {
	errs []error
}

func NewEnvBinder() (b EnvBinder) {
	b.errs = make([]error, 0)
	return
}

func (b *EnvBinder) addError(err error) {
	if err == nil {
		return
	}
	b.errs = append(b.errs, err)
}

func (b *EnvBinder) BindError() error {
	if len(b.errs) > 0 {
		b.errs = b.errs[1:]
		return b.errs[0]
	}
	return nil
}

func (b *EnvBinder) BindErrors() []error {
	errors := b.errs
	b.errs = make([]error, 0)
	return errors
}
