package locker

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
)

func TestLocker(t *testing.T) {
	loc := New()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	resources := []string{"a", "b", "c", "d"}
	values := []int{0, 0, 0, 0}
	a := 0
	b := 0
	c := 0
	d := 0

	// run multiple goroutines to test lock
	var g errgroup.Group
	for i := 0; i < 1000; i++ {
		read := r.Intn(4) == 0
		if !read {
			values[0]++
		}
		g.Go(func() error {
			singleAction(loc, resources[0], read, func() {
				a++
			})
			return nil
		})
	}
	for i := 0; i < 1000; i++ {
		read := r.Intn(4) == 0
		if !read {
			values[1]++
		}
		g.Go(func() error {
			singleAction(loc, resources[1], read, func() {
				b++
			})
			return nil
		})
	}
	for i := 0; i < 1000; i++ {
		read := r.Intn(4) == 0
		if !read {
			values[2]++
		}
		g.Go(func() error {
			singleAction(loc, resources[2], read, func() {
				c++
			})
			return nil
		})
	}
	for i := 0; i < 1000; i++ {
		read := r.Intn(4) == 0
		if !read {
			values[3]++
		}
		g.Go(func() error {
			singleAction(loc, resources[3], read, func() {
				d++
			})
			return nil
		})
	}
	err := g.Wait()
	assert.NoError(t, err, "should have no error.")

	assert.Equal(t, a, values[0], "a wrong")
	assert.Equal(t, b, values[1], "b wrong")
	assert.Equal(t, c, values[2], "c wrong")
	assert.Equal(t, d, values[3], "d wrong")

	assert.Len(t, loc.info, 0, "should be empty")
}

func singleAction(loc *Locker, resource string, read bool, writeFun func()) {
	if read {
		defer loc.ReadLock(resource)()
	} else {
		defer loc.WriteLock(resource)()
		writeFun()
	}
}

func BenchmarkLocker(b *testing.B) {
	resources := []string{"a", "b", "c", "d"}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	loc := New()

	for i := 0; i < b.N; i++ {
		num := r.Intn(4)
		read := num%2 == 0
		singleAction(loc, resources[num], read, func() {})
	}
}
