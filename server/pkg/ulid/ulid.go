package ulid

import (
	"io"
	"math/rand"
	"sync"
	"time"

	"github.com/oklog/ulid/v2"
)

type ULIDGenerator struct {
	pool *sync.Pool
}

type generator struct {
	r io.Reader
}

func NewULIDGenerator() *ULIDGenerator {
	pool := &sync.Pool{
		New: func() interface{} {
			return &generator{r: ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)}
		},
	}
	return &ULIDGenerator{pool: pool}
}

func (g *generator) New() ulid.ULID {
	return ulid.MustNew(ulid.Timestamp(time.Now()), g.r)
}

func (u *ULIDGenerator) Generate() string {
	g := u.pool.Get().(*generator)
	id := g.New().String()
	u.pool.Put(g)
	return id
}
