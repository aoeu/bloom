package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"
	"log"
	"math/rand"
)

// TODO(aoeu): There's libraries with bit arrays for bit torrent - "bit set" - check them out.
type bitarray []bool

type BloomFilter struct {
	bitarray
	seeds  []int64
	hasher hash.Hash32
}

func New(numHashFunctions int) (*BloomFilter, error) {
	k := numHashFunctions
	m := 1024
	if k < 1 {
		return new(BloomFilter), errors.New(fmt.Sprintf("A bloom filter must have more than 0 hash functions: %v\n", k))
	}
	if k >= 255 {
		return new(BloomFilter), errors.New(fmt.Sprintf("A bloom filter must have less than 255 hash functions: %v\n", k))
	}
	b := &BloomFilter{
		bitarray: make(bitarray, m),
		seeds:    make([]int64, k, k),
		hasher:   fnv.New32(),
	}
	rng := rand.New(rand.NewSource(777))
	for i := 0; i < k; i++ {
		b.seeds[i] = rng.Int63()
	}
	return b, nil
}

func (b *BloomFilter) Put(value interface{}) error {
	for i := 0; i < len(b.seeds); i++ {
		b.hasher.Reset()
		if err := binary.Write(b.hasher, binary.LittleEndian, b.seeds[i]); err != nil {
			return err
		}
		if err := binary.Write(b.hasher, binary.LittleEndian, value); err != nil {
			return err
		}
		h := b.hasher.Sum32()
		offset := h % uint32(len(b.bitarray))
		fmt.Printf("%v - %v\n", h, offset)
		b.bitarray[offset] = true
	}
	return nil
}

func (b *BloomFilter) MightContain(value interface{}) bool {
	for i := 0; i < len(b.seeds); i++ {
		b.hasher.Reset()
		binary.Write(b.hasher, binary.LittleEndian, b.seeds[i])
		binary.Write(b.hasher, binary.LittleEndian, value)
		offset := b.hasher.Sum32() % uint32(len(b.bitarray))
		if !b.bitarray[offset] {
			return false
		}
	}
	return true
}

/*
 * Cheat sheet:
 *
 * m: total bits
 * n: expected insertions
 * b: m/n, bits per insertion
 * p: expected false positive probability
 * k: number of hashes
 *
 * 1) Optimal k = b * ln2
 * 2) p = (1 - e ^ (-kn/m))^k
 * 3) For optimal k: p = 2 ^ (-k) ~= 0.6185^b
 * 4) For optimal k: m = -nlnp / ((ln2) ^ 2)
 */

/*
func (b *BloomFilter) expectedFalsePositiveProbability {
	bitSize := 1 // TODO(aoeu): What value is this supposed to be?
	return math.Pow(len(b.bits) / bitSize, b.numHashFunctions)
	// TODO(aoeu): This should be (1 - e ^ (-kn/m)) ^ k
}

/ Computes the optimal k (number of hashes per element inserted in Bloom filter), given the
// expected insertions and total number of bits in the Bloom filter.
//
// See http://en.wikipedia.org/wiki/File:Bloom_filter_fp_probability.svg for the formula.
//
func optimalNumOfHashFunctions(expectedInsertions float64, totalBits float64) float64 {
	return math.Max(1, (expectedInsertions / totalBits) * math.Log(2))
}


func optimalNumOfBits(expectedInsertions float64, falsePositiveRate float64) float64 {
	if falsePostiveRate == 0.0 {
		return 0.0
	}
	log2 := math.Log(2)
	return -n * math.Log(falsePositiveRate) / (log2 * log2)
}

func New(expectedInsertions float64, falsePostiveRate float64) {
	numBits := optimalNumOfBits(expectedInsertions, falsePositiveRate)
	numHashFuctions := optimalNumOfHashFunctions(expectedInsertions, numBits)
	b, err := New(numHashFunctions)
`	if err != nil {
		return b, err
	}
	b.bits = make(bitarray, numBits)
	return b, nil
}


*/

func main() {
	bf, err := New(5)
	if err != nil {
		log.Fatal(err)
	}

	if err := bf.Put([]byte("foo")); err != nil {
		log.Fatal(err)
	}
	fmt.Println(bf.MightContain([]byte("foo")))

	bf.Put([]byte("bar"))
	fmt.Println(bf.MightContain([]byte("bar")))

	fmt.Println(bf.MightContain([]byte("baz")))
	fmt.Println(bf.MightContain([]byte("qux")))
}
