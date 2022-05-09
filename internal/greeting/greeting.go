// Package greeting - use this package to print greeting to console
package greeting

import (
	"fmt"
)

// Greeting struct instance this to start printing greeting
type Greeting struct {
	Name string
}

// SayHello prints greeting to console
func (g *Greeting) SayHello() {
	if g.Name == "" {
		fmt.Printf("Hello \n")
	} else {
		fmt.Printf("Hello, %s \n", g.Name)
	}

	// Output: Hello, John
}
