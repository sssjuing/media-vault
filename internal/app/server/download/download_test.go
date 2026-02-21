package download

import (
	"crypto/md5"
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"
)

func TestHashSet(t *testing.T) {
	s := mapset.NewSet[[16]byte]()
	assert.Equal(t, true, s.Add(md5.Sum([]byte("math"))), "should be equal")
	assert.Equal(t, false, s.Add(md5.Sum([]byte("math"))), "should be equal")
	s.Add(md5.Sum([]byte("english")))
	s.Add(md5.Sum([]byte("biology")))
	assert.Equal(t, true, s.Contains(md5.Sum([]byte("english"))))
	s.Remove(md5.Sum([]byte("english")))
	assert.Equal(t, false, s.Contains(md5.Sum([]byte("english"))))
}
