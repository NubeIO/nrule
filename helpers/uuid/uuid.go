package uuid

import (
	"crypto/rand"
	"fmt"
	"io"
)

func ShortUUID(prefix ...string) string {
	u := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, u)
	if n != len(u) || err != nil {
		return "-error-uuid-"
	}
	uuid := fmt.Sprintf("%x%x", u[0:4], u[4:6])
	if len(prefix) > 0 {
		uuid = fmt.Sprintf("%s_%s", prefix[0], uuid)
	}
	return uuid
}
