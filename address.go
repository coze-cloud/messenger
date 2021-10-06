package messenger

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"os"
	"strings"
)

type address struct {
	Id   uuid.UUID `json:"id"`
	Name string `json:"name"`
}

func newAddress() *address {
	name, _ := os.Hostname()

	return &address{
		Id:   uuid.NewV4(),
		Name: name,
	}
}

func (address address) String() string {
	return fmt.Sprintf("%s(%s)",
		strings.Split(address.Id.String(), "-")[0],
		address.Name,
	)
}