package jobs

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
)

func sha256Encode(pwd string) string {
	h := sha256.New()
	io.WriteString(h, pwd)
	return hex.EncodeToString(h.Sum(nil))
}
