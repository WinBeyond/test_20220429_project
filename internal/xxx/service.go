// Package xxx provides xxx logic implement
package xxx

import (
	"fmt"
)

// Xxx
type Xxx struct{
}

// NewService
func NewService() *Xxx {
	service := &Xxx {
		// init custom fileds
	}
	return service
}

func (s *Xxx) Helloworld() {
	fmt.Printf("Hello World!\n")
	// implement business logic here ...
	// ...
	return
}