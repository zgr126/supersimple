package main

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"log"
)

//
func cryptoByte(b []byte) string {
	h := sha256.New()
	_b := h.Sum(b)

	return hex.EncodeToString(_b)
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoui64(b []byte) uint64 {
	log.Print(b)
	return binary.BigEndian.Uint64(b)
}
