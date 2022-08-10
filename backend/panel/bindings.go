package panel

import "fmt"

type Bindings struct {
}

func (b *Bindings) Greet(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}
