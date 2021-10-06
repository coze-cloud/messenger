package messenger

import (
	uuid "github.com/satori/go.uuid"
	"os"
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