package contrast_benchmark

import (
	"log"
	"path/filepath"
	"testing"

	"github.com/tidwall/buntdb"
)

var bdb *buntdb.DB

func init() {
	// Open the data.db file. It will be created if it doesn't exist.
	path := filepath.Join("benchmark", "buntdb", "data.db")
	bdb, err = buntdb.Open(path)
	if err != nil {
		log.Fatal(err)
	}
}

func initBuntDBData() {
	for i := 0; i < 500000; i++ {
		bdb.Update(func(tx *buntdb.Tx) error {
			_, _, err := tx.Set(string(GetKey(i)), string(GetValue()), nil)
			if err != nil {
				panic(err)
			}
			return nil
		})

	}
}

func Benchmark_PutValue_BuntDB(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		bdb.Update(func(tx *buntdb.Tx) error {
			_, _, err := tx.Set(string(GetKey(i)), string(GetValue()), nil)
			if err != nil {
				panic(err)
			}
			return nil
		})
	}
}

func Benchmark_GetValue_BuntDB(b *testing.B) {
	initBuntDBData()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {

		bdb.View(func(tx *buntdb.Tx) error {
			_, _ = tx.Get(string(GetKey(i)))
			return nil
		})
	}
}
