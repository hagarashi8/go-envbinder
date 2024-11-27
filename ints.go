package envbinder

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

func atoi(s string, bitSize int) (int64, error) {
	return strconv.ParseInt(s, 10, bitSize)
}

func lookupInt[T constraints.Integer](name string, bitsize int) (T, error) {
	s, err := lookupString(name)
	if err != nil {
		return 0, err
	}
	i, err := atoi(s, bitsize)
	if err != nil {
		return 0, err
	}
	return T(i), nil
}

func (b *EnvBinder) Int8(v *int8, name string) *EnvBinder {
	i, err := lookupInt[int8](name, 8)
	if err != nil {
		b.addError(err)
		return b
	}
	*v = i
	return b
}

func (b *EnvBinder) Int8OrDef(v *int8, name string, def int8) *EnvBinder {
	i, err := lookupInt[int8](name, 8)
	if err != nil {
		*v = def
		return b
	}
	*v = i
	return b
}

func (b *EnvBinder) Int8s(v *[]int8, name string, sep string) *EnvBinder {
	s, err := lookupStrings(name, sep)
	slice := []int8{}
	b.addError(err)
	for _, si := range s {
		i, err := atoi(si, 8)
		if err != nil {
			b.addError(err)
			break
		}
		slice = append(slice, int8(i))
	}
	*v = slice
	return b
}

func (b *EnvBinder) Int8sOrDef(v *[]int8, name string, sep string, def []int8) *EnvBinder {
	s, err := lookupStrings(name, sep)
	slice := []int8{}
	b.addError(err)
	for _, si := range s {
		i, err := atoi(si, 8)
		if err != nil {
			slice = def
			break
		}
		slice = append(slice, int8(i))
	}
	*v = slice
	return b
}

func (b *EnvBinder) Int16(v *int16, name string) *EnvBinder {
	i, err := lookupInt[int16](name, 16)
	if err != nil {
		b.addError(err)
		return b
	}
	*v = i
	return b
}

func (b *EnvBinder) Int16OrDef(v *int16, name string, def int16) *EnvBinder {
	i, err := lookupInt[int16](name, 16)
	if err != nil {
		*v = def
		return b
	}
	*v = i
	return b
}

func (b *EnvBinder) Int16s(v *[]int16, name string, sep string) *EnvBinder {
	s, err := lookupStrings(name, sep)
	slice := []int16{}
	b.addError(err)
	for _, si := range s {
		i, err := atoi(si, 16)
		if err != nil {
			b.addError(err)
			break
		}
		slice = append(slice, int16(i))
	}
	*v = slice
	return b
}

func (b *EnvBinder) Int16sOrDef(v *[]int16, name string, sep string, def []int16) *EnvBinder {
	s, err := lookupStrings(name, sep)
	slice := []int16{}
	b.addError(err)
	for _, si := range s {
		i, err := atoi(si, 16)
		if err != nil {
			slice = def
			break
		}
		slice = append(slice, int16(i))
	}
	*v = slice
	return b
}

func (b *EnvBinder) Int32(v *int32, name string) *EnvBinder {
	i, err := lookupInt[int32](name, 32)
	if err != nil {
		b.addError(err)
		return b
	}
	*v = i
	return b
}

func (b *EnvBinder) Int32OrDef(v *int32, name string, def int32) *EnvBinder {
	i, err := lookupInt[int32](name, 32)
	if err != nil {
		*v = def
		return b
	}
	*v = i
	return b
}

func (b *EnvBinder) Int32s(v *[]int32, name string, sep string) *EnvBinder {
	s, err := lookupStrings(name, sep)
	slice := []int32{}
	b.addError(err)
	for _, si := range s {
		i, err := atoi(si, 32)
		if err != nil {
			b.addError(err)
			break
		}
		slice = append(slice, int32(i))
	}
	*v = slice
	return b
}

func (b *EnvBinder) Int32sOrDef(v *[]int32, name string, sep string, def []int32) *EnvBinder {
	s, err := lookupStrings(name, sep)
	slice := []int32{}
	b.addError(err)
	for _, si := range s {
		i, err := atoi(si, 32)
		if err != nil {
			slice = def
			break
		}
		slice = append(slice, int32(i))
	}
	*v = slice
	return b
}

func (b *EnvBinder) Int64(v *int64, name string) *EnvBinder {
	i, err := lookupInt[int64](name, 64)
	if err != nil {
		b.addError(err)
		return b
	}
	*v = i
	return b
}

func (b *EnvBinder) Int64OrDef(v *int64, name string, def int64) *EnvBinder {
	i, err := lookupInt[int64](name, 64)
	if err != nil {
		*v = def
		return b
	}
	*v = i
	return b
}

func (b *EnvBinder) Int64s(v *[]int64, name string, sep string) *EnvBinder {
	s, err := lookupStrings(name, sep)
	slice := []int64{}
	b.addError(err)
	for _, si := range s {
		i, err := atoi(si, 64)
		if err != nil {
			b.addError(err)
			break
		}
		slice = append(slice, int64(i))
	}
	*v = slice
	return b
}

func (b *EnvBinder) Int64sOrDef(v *[]int64, name string, sep string, def []int64) *EnvBinder {
	s, err := lookupStrings(name, sep)
	slice := []int64{}
	b.addError(err)
	for _, si := range s {
		i, err := atoi(si, 64)
		if err != nil {
			slice = def
			break
		}
		slice = append(slice, int64(i))
	}
	*v = slice
	return b
}

func (b *EnvBinder) Int(v *int8, name string) *EnvBinder {
	i, err := lookupInt[int8](name, 0)
	if err != nil {
		b.addError(err)
		return b
	}
	*v = i
	return b
}

func (b *EnvBinder) IntOrDef(v *int, name string, def int) *EnvBinder {
	i, err := lookupInt[int](name, 0)
	if err != nil {
		*v = def
		return b
	}
	*v = i
	return b
}

func (b *EnvBinder) Ints(v *[]int, name string, sep string, def []int) *EnvBinder {
	s, err := lookupStrings(name, sep)
	slice := []int{}
	b.addError(err)
	for _, si := range s {
		i, err := atoi(si, 0)
		if err != nil {
			b.addError(err)
			break
		}
		slice = append(slice, int(i))
	}
	*v = slice
	return b
}

func (b *EnvBinder) IntsOrDef(v *[]int, name string, sep string, def []int) *EnvBinder {
	s, err := lookupStrings(name, sep)
	slice := []int{}
	b.addError(err)
	for _, si := range s {
		i, err := atoi(si, 0)
		if err != nil {
			slice = def
			break
		}
		slice = append(slice, int(i))
	}
	*v = slice
	return b
}
