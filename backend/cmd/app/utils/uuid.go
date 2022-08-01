package utils

import (
	"encoding/base64"
	"github.com/google/uuid"
)

func GenerateUUID() string {
	uuid4 := uuid.New()
	uuid4Binary, _ := uuid4.MarshalBinary()
	uuid4EncodedBase64 := base64.RawURLEncoding.EncodeToString(uuid4Binary)
	return uuid4EncodedBase64
}
