package goptional

func NewEmptyOptional[ValueType any]() Optional[ValueType] {
	return Optional[ValueType]{}
}

func NewOptional[ValueType any](value ValueType) Optional[ValueType] {
	return Optional[ValueType]{
		value: &value,
	}
}

func NewOptionalPtr[ValueType any](value *ValueType) Optional[ValueType] {
	return Optional[ValueType]{
		value: value,
	}
}

type Optional[Valuetype any] struct {
	value *Valuetype
}

func (o *Optional[Valuetype]) Present() bool {
	return o.value != nil
}

func (o *Optional[Valuetype]) Get() Valuetype {
	return *o.value
}

func (o *Optional[Valuetype]) If(fun func(Valuetype) error) (err error) {
	if o.Present() {
		err = fun(o.Get())
	}
	return
}

func (o *Optional[Valuetype]) Else(fun func() error) (err error) {
	if o.Present() {
		err = fun()
	}
	return
}
