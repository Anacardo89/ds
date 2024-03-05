package arrlst

import "errors"

var (
	ErrIdxOutOfBounds = errors.New("index out of bounds")
	ErrNotFound       = errors.New("value not found")
)

type Arrlst struct {
	len int
	cap int
	arr []interface{}
}

func New(cap int) Arrlst {
	return Arrlst{
		len: 0,
		cap: cap,
		arr: make([]interface{}, cap),
	}
}

func (a *Arrlst) grow() {
	a.cap = a.cap + int(a.cap/2)
	arr := make([]interface{}, a.cap)
	for i := 0; i < a.len; i++ {
		arr[i] = a.arr[i]
	}
	a.arr = arr
}

func (a *Arrlst) Prepend(val interface{}) {
	if a.len == a.cap {
		a.grow()
	}
	a.len++
	for i := a.len - 1; i > 0; i-- {
		a.arr[i] = a.arr[i-1]
	}
	a.arr[0] = val
}

func (a *Arrlst) InsertAt(val interface{}, idx int) error {
	if idx < 0 || idx > a.len {
		return ErrIdxOutOfBounds
	}
	if a.len == a.cap {
		a.grow()
	}
	if idx == a.len {
		a.arr[idx] = val
		a.len++
		return nil
	}
	for i := a.len - 1; i >= idx; i-- {
		a.arr[i+1] = a.arr[i]
	}
	a.arr[idx] = val
	a.len++
	return nil
}

func (a *Arrlst) Append(val interface{}) {
	if a.len == a.cap {
		a.grow()
	}
	a.arr[a.len] = val
	a.len++
}

func (a *Arrlst) Remove(val interface{}) (interface{}, error) {
	idx := -1
	for i := 0; i < a.len; i++ {
		if a.arr[i] == val {
			idx = i
		}
	}
	if idx < 0 {
		return nil, ErrNotFound
	}
	ret := a.arr[idx]
	for i := idx; i < a.len-1; i++ {
		a.arr[i] = a.arr[i+1]
	}
	a.arr[a.len-1] = nil
	a.len--
	return ret, nil
}

func (a *Arrlst) Get(idx int) (interface{}, error) {
	if idx < 0 || idx > a.len-1 {
		return nil, ErrIdxOutOfBounds
	}
	return a.arr[idx], nil
}

func (a *Arrlst) RemoveAt(idx int) (interface{}, error) {
	if idx < 0 || idx > a.len-1 {
		return nil, ErrIdxOutOfBounds
	}
	ret := a.arr[idx]
	for i := idx; i < a.len-1; i++ {
		a.arr[i] = a.arr[i+1]
	}
	a.arr[a.len-1] = nil
	a.len--
	return ret, nil
}
