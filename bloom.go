package bloom

import (
	"fmt"
	"math"
	"errors"
)

type bitarray []bool

type BloomFilter struct {
	bits bitarray
	numHashFunctions int
	funnel
}

type funnel interface{} // TODO(aoeu): What is this supposed to be?

func NewBloomFilter(numHashFunctions int) (*BloomFilter, error) {
	n := numHashFunctions
	if n < 0 {
		return new(BloomFilter), errors.New(fmt.Sprintf("A bloom filter must have more than 0 hash functions: %v\n", n)) 
	}
	if n >= 255 {
		return new(BloomFilter), errors.New(fmt.Sprintf("A bloom filter must have less than 255 hash functions: %v\n", n))
	}
	b := new(BloomFilter)
	b.numHashFunctions = n
	b.funnel = funnel{}
	return b, nil
}


func (b *BloomFilter) expectedFalsePositiveProbability {
	bitSize := 1 // TODO(aoeu): What value is this supposed to be?
	return math.Pow(len(b.bits) / bitSize, b.numHashFunctions)
}

// Computes the optimal k (number of hashes per element inserted in Bloom filter), given the
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

func NewBloomFilter(expectedInsertions float642, falsePostiveRate float64) {
	numBits := optimalNumOfBits(expectedInsertions, falsePositiveRate)
	numHashFuctions := optimalNumOfHashFunctions(expectedInsertions, numBits)
	b, err := NewBloomFilter(numHashFunctions)
`	if err != nil {
		return b, err
	}	
	b.bits = make(bitarray, numBits)
	return b, nil
}


var seed C.int // TODO(aoeu): What is this supposed to be?

func (b *BloomFilter) put(i interface{}) bool {
	bitSize := 1 // TODO(aoeu): What is this supposed to be?
	key := "TODO" // TODO(aoeu): What is this value supposed to be?
	hash1 := C.murmur3_32(C.CString(key), bitSize, seed) // TODO(aoeu): How many bits is this?
	hash2 := C.murmur3_32(C.CString(key), bitSize, seed)
	// uint32_t murmur3_32(const char *key, uint32_t len, uint32_t seed) 
	changedBits := false
	for i := 0; i < b.numHashFunctions; i++ {
		combinedHash := hash1 + (i * hash2)	
	}
	if _, ok := b.bitarray[combinedHash % bitSize]; !ok {
		changedBits = true
	}
	return changedBits
}

/*
  MURMUR128_MITZ_32() {
    @Override
    public <T> boolean put(
        T object, Funnel<? super T> funnel, int numHashFunctions, BitArray bits) {
      long bitSize = bits.bitSize();
      long hash64 = Hashing.murmur3_128().hashObject(object, funnel).asLong();
      int hash1 = (int) hash64;
      int hash2 = (int) (hash64 >>> 32);

      boolean bitsChanged = false;
      for (int i = 1; i <= numHashFunctions; i++) {
        int combinedHash = hash1 + (i * hash2);
        // Flip all the bits if it's negative (guaranteed positive number)
        if (combinedHash < 0) {
          combinedHash = ~combinedHash;
        }
        bitsChanged |= bits.set(combinedHash % bitSize);
      }
      return bitsChanged;
    }
*/


/*
    @Override
    public <T> boolean mightContain(
        T object, Funnel<? super T> funnel, int numHashFunctions, BitArray bits) {
      long bitSize = bits.bitSize();
      long hash64 = Hashing.murmur3_128().hashObject(object, funnel).asLong();
      int hash1 = (int) hash64;
      int hash2 = (int) (hash64 >>> 32);

      for (int i = 1; i <= numHashFunctions; i++) {
        int combinedHash = hash1 + (i * hash2);
        // Flip all the bits if it's negative (guaranteed positive number)
        if (combinedHash < 0) {
          combinedHash = ~combinedHash;
        }
        if (!bits.get(combinedHash % bitSize)) {
          return false;
        }
      }
      return true;
    }
  },
*/

/*
#include <stdint.h>

uint32_t murmur3_32(const char *key, uint32_t len, uint32_t seed) {
	static const uint32_t c1 = 0xcc9e2d51;
	static const uint32_t c2 = 0x1b873593;
	static const uint32_t r1 = 15;
	static const uint32_t r2 = 13;
	static const uint32_t m = 5;
	static const uint32_t n = 0xe6546b64;
 
	uint32_t hash = seed;
 
	const int nblocks = len / 4;
	const uint32_t *blocks = (const uint32_t *) key;
	int i;
	for (i = 0; i < nblocks; i++) {
		uint32_t k = blocks[i];
		k *= c1;
		k = (k << r1) | (k >> (32 - r1));
		k *= c2;
 
		hash ^= k;
		hash = ((hash << r2) | (hash >> (32 - r2))) * m + n;
	}
 
	const uint8_t *tail = (const uint8_t *) (key + nblocks * 4);
	uint32_t k1 = 0;
 
	switch (len & 3) {
	case 3:
		k1 ^= tail[2] << 16;
	case 2:
		k1 ^= tail[1] << 8;
	case 1:
		k1 ^= tail[0];
 
		k1 *= c1;
		k1 = (k1 << r1) | (k1 >> (32 - r1));
		k1 *= c2;
		hash ^= k1;
	}
 
	hash ^= len;
	hash ^= (hash >> 16);
	hash *= 0x85ebca6b;
	hash ^= (hash >> 13);
	hash *= 0xc2b2ae35;
	hash ^= (hash >> 16);
 
	return hash;
}
*/
import "C"
