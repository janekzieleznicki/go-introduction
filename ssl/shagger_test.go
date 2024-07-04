package ssl

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "crypto/sha256"

)
func TestSha256(t *testing.T){
  assert := assert.New(t)
  assert.Equal(expected, Sha256(input), "Hash should produce expected")
}
func TestSha256Calledback(t *testing.T){
  Sha256Calledback("aaaa")
}
var input = "Decidable-Unsavory-Marmalade-Onward-Bazooka-Supply-Hardness-Boondocks-Cosmic-Improving"
var expected = "872fa4d06aeac5798bd7a1412e32786196fae598bf7dbc1667f1f65a0f6cb4e6"
func BenchmarkSha256(b *testing.B){
  for i := 0; i < b.N; i++ {
    Sha256(input)
  }
}
func BenchmarkSha256Byte(b *testing.B){
  byted := []byte(input)
  output := []byte{}
  for i := 0; i < b.N; i++ {
    Sha256Byte(byted, &output)
  }
}
func BenchmarkGoLibSha(b *testing.B){
  byted := []byte(input)
  output := []byte{}
  for i := 0; i < b.N; i++ {
    h := sha256.New()
    h.Write(byted)
    h.Sum(output)
  }
}