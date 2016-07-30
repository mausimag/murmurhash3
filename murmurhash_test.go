package murmurhash

import (
	"log"
	"testing"
)

func TestMurmurhash(t *testing.T) {
	b := []byte("teste1")
	log.Printf(">%d", Murmurhash3_x86_32(b, len(b), 0))

	b2 := []byte("teste2")
	log.Printf(">%d", Murmurhash3_x86_32(b2, len(b2), 0))
}
