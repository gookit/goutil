package strutil_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestMicroTimeID(t *testing.T) {
	fmt.Println("microTimeID:")
	for i := 0; i < 10; i++ {
		id := strutil.MicroTimeID()
		fmt.Println(id, "len:", len(id))
		assert.NotEmpty(t, id)
	}

	id := strutil.MicroTimeID()
	fmt.Println("mtID :", id, "len:", len(id))
	b16id := strutil.Base10Conv(id, 16)
	fmt.Println("b16id:", b16id, "len:", len(b16id))
	b32id := strutil.Base10Conv(id, 32)
	fmt.Println("b32id:", b32id, "len:", len(b32id))
	b36id := strutil.Base10Conv(id, 36)
	fmt.Println("b36id:", b36id, "len:", len(b36id))
	b62id := strutil.Base10Conv(id, 62)
	fmt.Println("b62id:", b62id, "len:", len(b62id))
}

func TestMicroTimeBaseID(t *testing.T) {
	t.Run("Base 16", func(t *testing.T) {
		fmt.Println("MicroTimeHexID:")
		idMap := make(map[string]bool)
		for i := 0; i < 10; i++ {
			id := strutil.MicroTimeHexID()
			assert.False(t, idMap[id])

			idMap[id] = true
			fmt.Println(id, "len:", len(id))
			assert.NotEmpty(t, id)
			time.Sleep(time.Millisecond)
		}
	})

	t.Run("Base 36", func(t *testing.T) {
		idMap := make(map[string]bool)
		fmt.Println("MTimeBaseID: 36")
		for i := 0; i < 10; i++ {
			id := strutil.MTimeBaseID(36)
			assert.False(t, idMap[id])

			idMap[id] = true
			fmt.Println(id, "len:", len(id))
			assert.NotEmpty(t, id)
			time.Sleep(time.Millisecond)
		}
		assert.NotEmpty(t, strutil.MTimeBase36())
	})

	t.Run("Base 48", func(t *testing.T) {
		idMap := make(map[string]bool)
		fmt.Println("MTimeBaseID: 48")
		for i := 0; i < 10; i++ {
			id := strutil.MTimeBaseID(48)
			assert.False(t, idMap[id])

			idMap[id] = true
			fmt.Println(id, "len:", len(id))
			assert.NotEmpty(t, id)
			time.Sleep(time.Millisecond)
		}
		assert.NotEmpty(t, strutil.MTimeBase36())
	})
}

func TestDateSN(t *testing.T) {
	fmt.Println("DatetimeNo:")
	for i := 0; i < 100; i++ {
		no := strutil.DatetimeNo("test")
		fmt.Println(no, "len:", len(no))
		assert.NotEmpty(t, no)
	}
}

func TestDateSNv2(t *testing.T) {
	fmt.Println("DateSNv2:")

	t.Run("base36", func(t *testing.T) {
		idMap := make(map[string]bool)
		for i := 0; i < 100; i++ {
			no := strutil.DateSNv2("T")
			assert.False(t, idMap[no], "duplicate sn: %s", no)

			idMap[no] = true
			fmt.Println(no, "len:", len(no))
			assert.NotEmpty(t, no)
		}
	})

	t.Run("base16", func(t *testing.T) {
		idMap := make(map[string]bool)
		for i := 0; i < 100; i++ {
			no := strutil.DateSNv2("T", 16)
			assert.False(t, idMap[no], "duplicate sn: %s", no)

			idMap[no] = true
			fmt.Println(no, "len:", len(no))
			assert.NotEmpty(t, no)
		}
	})
}

func TestDateSNv3(t *testing.T) {
	fmt.Println("DateSNv3:")
	no := strutil.DateSNv3("", 8)
	fmt.Println(no, "len:", len(no))

	t.Run("dateLen8 base32", func(t *testing.T) {
		idMap := make(map[string]bool)
		for i := 0; i < 50; i++ {
			no = strutil.DateSNv3("T", 8)
			assert.NotEmpty(t, no)
			fmt.Println(no, "len:", len(no))

			assert.False(t, idMap[no], "duplicate sn: %s", no)
			idMap[no] = true
		}
	})

	t.Run("dateLen6 base36", func(t *testing.T) {
		idMap := make(map[string]bool)
		for i := 0; i < 60; i++ {
			no = strutil.DateSNv3("T", 6, 36)
			fmt.Println(no, "len:", len(no))
			assert.NotEmpty(t, no)
			assert.False(t, idMap[no], "duplicate sn: %s", no)
			idMap[no] = true
		}
	})

	t.Run("dateLen6 base48", func(t *testing.T) {
		idMap := make(map[string]bool)
		for i := 0; i < 60; i++ {
			no = strutil.DateSNv3("T", 6, 48)
			fmt.Println(no, "len:", len(no))
			assert.NotEmpty(t, no)
			assert.False(t, idMap[no], "duplicate sn: %s", no)
			idMap[no] = true
		}
	})

	t.Run("dateLen4 base62", func(t *testing.T) {
		idMap := make(map[string]bool)
		for i := 0; i < 60; i++ {
			no = strutil.DateSNv3("T", 4, 62)
			fmt.Println(no, "len:", len(no))
			assert.NotEmpty(t, no)
			assert.False(t, idMap[no], "duplicate sn: %s", no)
			idMap[no] = true
		}
	})
}

func TestDateSNOpt_GenSN(t *testing.T) {
	fmt.Println("DateSNOpt.GenSN:")

	t.Run("basic", func(t *testing.T) {
		opt := &strutil.DateSNOpt{}
		idMap := make(map[string]bool)
		for i := 0; i < 100; i++ {
			sn := opt.GenSN("T")
			assert.NotEmpty(t, sn)
			assert.False(t, idMap[sn], "duplicate sn: %s", sn)
			idMap[sn] = true
			fmt.Println(sn, "len:", len(sn))
		}
	})

	t.Run("with custom base", func(t *testing.T) {
		opt := &strutil.DateSNOpt{ConvBase: 62, DateLen: 6}
		for i := 0; i < 10; i++ {
			sn := opt.GenSN("P")
			fmt.Println(sn, "len:", len(sn))
			assert.NotEmpty(t, sn)
		}
	})

	t.Run("with sequence enabled", func(t *testing.T) {
		opt := &strutil.DateSNOpt{ConvBase: 36}
		for i := 0; i < 10; i++ {
			sn := opt.GenSN("")
			fmt.Println(sn, "len:", len(sn))
			assert.NotEmpty(t, sn)
		}
	})

	t.Run("concurrent safe", func(t *testing.T) {
		fmt.Println("concurrent generated 100 unique IDs")
		opt := &strutil.DateSNOpt{ConvBase: 36}
		idMap := make(map[string]bool)
		var mu sync.Mutex
		var wg sync.WaitGroup

		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				sn := opt.GenSN("")
				mu.Lock()
				assert.False(t, idMap[sn], "duplicate sn: %s", sn)
				idMap[sn] = true
				mu.Unlock()
			}()
		}
		wg.Wait()
	})

	t.Run("use random and base36", func(t *testing.T) {
		opt := &strutil.DateSNOpt{ConvBase: 36}
		idMap := make(map[string]bool)

		for i := 0; i < 100; i++ {
			sn := opt.GenSN("")
			assert.NotEmpty(t, sn)
			assert.False(t, idMap[sn], "duplicate sn: %s", sn)
			idMap[sn] = true
		}
	})
}
