package parmesan

import (
  "testing"
)

const EMPTY_STR_SHA512 = "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e"

func TestSha512Hash(t *testing.T) {
  b := &BittrexClient{}
  got := b.hash([]byte{})
  if got != EMPTY_STR_SHA512 {
    t.Errorf("Calculated wrong hash for the empty string.")
  }
}
