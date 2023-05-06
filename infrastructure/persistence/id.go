package persistence

import (
	"github.com/oklog/ulid"
	"math/rand"
	"time"
)

func generateID() ulid.ULID {
	t := time.Now()
	return ulid.MustNew(ulid.Timestamp(t), ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0))
}
