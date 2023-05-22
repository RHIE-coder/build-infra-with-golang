package utils

import "github.com/google/uuid"

func GetUUID4() string {
	id, _ := uuid.NewRandom()
	return id.String()
}
