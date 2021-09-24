package xxhash

import "github.com/OneOfOne/xxhash"

func XXHash32(key []byte) uint32 {
	h := xxhash.New32()
	h.Write(key)
	return h.Sum32()
}

func XXHash64(key []byte) uint64 {
	h := xxhash.New64()
	h.Write(key)
	return h.Sum64()
}
