package messenger

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"os"
	"strings"
)

type address struct {
	Id uuid.UUID
	Name string
}

func newRandomAddress() address {
	return newAddress(uuid.NewV4())
}

func newAddress(id uuid.UUID) address {
	address := new(address)
	address.Id = id
	hostname, err := os.Hostname()
	if err != nil {
		return *address
	}
	address.Name = hostname

	return *address
}

func (address address) String() string {
	return fmt.Sprintf("%s(%s)",
		strings.Split(address.Id.String(), "-")[0],
		address.Name,
	)
}