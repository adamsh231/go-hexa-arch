package libraries

import "sync"

type IPubSubLibrary interface {
	Publish()
	Subscribe(wg *sync.WaitGroup, handler func(message []byte))
}