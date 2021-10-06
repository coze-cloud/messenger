package messenger

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"os"
	"strings"
)

type address struct {
	id uuid.UUID
	name string
}

func newAddress() *address {
	name, _ := os.Hostname()

	return &address{
		id: uuid.NewV4(),
		name: name,
	}
}

func (address address) String() string {
	return fmt.Sprintf("%s(%s)",
		strings.Split(address.id.String(), "-")[0],
		address.name,
	)
}