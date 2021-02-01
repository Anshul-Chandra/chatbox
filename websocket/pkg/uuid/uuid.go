package uuid

import (
	"log"

	"github.com/google/uuid"
)

// GetUUID generates a new uuid
func GetUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		log.Printf("unable to generate a uuid: %s", err.Error())
		return "", err
	}
	return id.String(), nil
}
