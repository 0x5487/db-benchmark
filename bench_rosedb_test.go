package contrast_benchmark

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/nite-coder/blackbear/pkg/cache"
	"github.com/rosedblabs/rosedb/v2"
)

var roseDB *rosedb.DB
var dirPath string
var roseCache *cache.Cache

func setupRoseDB() {

	dirPath = filepath.Join("benchmark", "rosedb")

	opts := rosedb.Options{
		DirPath:     dirPath,
		SegmentSize: 50 * rosedb.MB,
		BlockCache:  10 * rosedb.MB,
		Sync:        false,
	}

	var err error
	roseDB, err = rosedb.Open(opts)
	if err != nil {
		panic(err)
	}

}

func cleanupRoseDB() {
	os.RemoveAll(dirPath)
}

func initRoseDBData() {
	for i := 0; i < 10000; i++ {
		key := GetKey(i)
		val := GetValue()
		err := roseDB.Put(key, val)
		if err != nil {
			panic(err)
		}
	}

}

func Benchmark_PutValue_RoseDB(b *testing.B) {
	setupRoseDB()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		err := roseDB.Put(GetKey(i), GetValue())
		if err != nil {
			panic(err)
		}
	}

	cleanupRoseDB()
}

func Benchmark_GetValue_RoseDB(b *testing.B) {
	cleanupRoseDB()
	setupRoseDB()
	initRoseDBData()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		b, err := roseDB.Get(GetKey(1000))
		if len(b) == 0 {
			panic(errors.New("not found from cache"))
		}
		if err != nil {
			panic(errors.New("not found from cache"))
		}
	}
}
