package part

import(
	"fmt"
	"strings"
)

type Part struct {
    Id   int    // Named field (aggregation)
    Name string // Named field (aggregation)
}

func (part *Part) LowerCase() {
    part.Name = strings.ToLower(part.Name)
}

func (part *Part) UpperCase() {
    part.Name = strings.ToUpper(part.Name)
}

func (part Part) String() string {
    return fmt.Sprintf("%d %q", part.Id, part.Name)
}

func (part Part) HasPrefix(prefix string) bool {
    return strings.HasPrefix(part.Name, prefix)
}
