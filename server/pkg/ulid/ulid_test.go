package ulid_test

import (
	"os"
	"sync"
	"testing"

	"github.com/ZyoGo/Backend-Challange/pkg/ulid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	ulidService *ulid.ULIDGenerator
)

func TestMain(m *testing.M) {
	ulidService = ulid.NewULIDGenerator()
	os.Exit(m.Run())
}

func TestGenerateULID(t *testing.T) {
	t.Run("Should generate ULID with single id", func(t *testing.T) {
		resultId := ulidService.Generate()
		assert.NotNil(t, resultId)
	})

	t.Run("Should generate ULID for multiple ID with goroutine, make sure doesnt have duplicate ULID/ID", func(t *testing.T) {
		var (
			wg      sync.WaitGroup
			workers = 10
			count   = 1000
		)

		idMap := make(map[string]int)
		var mu sync.Mutex

		for i := 0; i < workers; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for j := 0; j < count; j++ {
					id := ulidService.Generate()

					mu.Lock()
					idMap[id]++
					mu.Unlock()
				}
			}()
		}
		wg.Wait()

		for id, count := range idMap {
			require.Equalf(t, 1, count, "Duplicate ID found: %s (Count: %d)", id, count)
		}
	})
}
