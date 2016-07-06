package metrics

import (
	"errors"
)

var ErrImplemented = errors.New("there is implemenation already injected")

var ctorImpl InternalNew = nil

// name is dot spearated path
// must be uniqe, use system naming, and unit postfix, examples:
//   ipfs.blockstore.bloomcache.bloom.miss.total
//   ipfs.routing.dht.notresuingstream.total
//
// both arguemnts are obligatory
func New(name, helptext string) Creator {
	if ctorImpl == nil {
		return &noop{}
	} else {
		return ctorImpl(name, helptext)
	}
}

type InternalNew func(string, string) Creator

func InjectImpl(newimpl InternalNew) error {
	if ctorImpl != nil {
		return ErrImplemented
	} else {
		ctorImpl = newimpl
		return nil
	}
}
