package testcraft

import (
	"errors"
	"reflect"
)

type AttrGenerator[T any] func(instance *T) error

type Factory[T any] struct {
	object   T
	sequence map[string]int
	attrsGen []AttrGenerator[T]
	valuer   Valuer
	t        reflect.Type
}

func NewFactory[T any](object T) *Factory[T] {
	return &Factory[T]{
		object:   object,
		sequence: make(map[string]int),
		valuer:   defaultValuer(),
	}
}

func (f *Factory[T]) Attr(attrsGen ...AttrGenerator[T]) *Factory[T] {
	f.attrsGen = append(f.attrsGen, attrsGen...)
	return f
}

func (f *Factory[T]) Build() (T, error) {
	return f.build()
}

func (f *Factory[T]) MustBuild() T {
	v, err := f.build()
	if err != nil {
		panic(err)
	}
	return v
}

func (f *Factory[T]) Randomize() (T, error) {
	return f.randomize(false)
}

func (f *Factory[T]) MustRandomize() T {
	res, err := f.randomize(false)
	if err != nil {
		panic(err)
	}
	return reflect.Indirect(reflect.ValueOf(res)).Interface().(T)
}

func (f *Factory[T]) RandomizeWithAttrs() (T, error) {
	return f.randomize(true)
}

func (f *Factory[T]) MustRandomizeWithAttrs() T {
	res, err := f.randomize(true)
	if err != nil {
		panic(err)
	}
	return reflect.Indirect(reflect.ValueOf(res)).Interface().(T)
}

func (f *Factory[T]) typeOf() reflect.Type {
	if f.t == nil {
		f.t = reflect.TypeOf(f.object)
	}
	return f.t
}

func (f *Factory[T]) build() (T, error) {
	t := f.typeOf()
	v := reflect.New(t)
	tp := v.Interface().(*T)
	errs := f.applyAttrs(tp)
	return reflect.Indirect(v).Interface().(T), errors.Join(errs...)
}

func (f *Factory[T]) applyAttrs(tp *T) []error {
	var errs []error
	for _, attr := range f.attrsGen {
		err := attr(tp)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func (f *Factory[T]) randomize(applyAttr bool) (T, error) {
	res, err := randomize(f.object, f.valuer)
	if err != nil {
		return reflect.Indirect(reflect.ValueOf(res)).Interface().(T), err
	}
	if applyAttr {
		f.applyAttrs(res.(*T))
	}
	return reflect.Indirect(reflect.ValueOf(res)).Interface().(T), err
}
