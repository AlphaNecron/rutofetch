package rutofetch

import (
	"fmt"
)

type Info struct {
	Name  string
	Value string
}

func (i Info) print() {
	fmt.Printf("%s%s: %s%s\n", Yellow.Fg(), i.Name, White.Fg(), i.Value)
}
