package goitertools

import "github.com/pchchv/express/optionext"

type fakeIterator struct {
	max int
}

func (f *fakeIterator) Next() optionext.Option[int] {
	f.max--
	if f.max < 0 {
		return optionext.None[int]()
	}

	return optionext.Some(f.max)
}
