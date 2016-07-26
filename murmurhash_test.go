package murmurhash

import (
	"log"
	"testing"
)

func TestMurmurhash(t *testing.T) {
	b := []byte("test")
	v := Murmurhash3_x86_32(b, len(b), 0)
	log.Printf(">%d", v)
}
