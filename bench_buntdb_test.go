package contrast_benchmark

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/tidwall/buntdb"
)

var bdb *buntdb.DB
var bdbPath string

func setupBuntDB() {
	bdbPath = filepath.Join("benchmark", "buntdb", "data.db")
	cleanupBuntDB()

	bdb, err = buntdb.Open(bdbPath)
	if err != nil {
		log.Fatal(err)
	}
}

func cleanupBuntDB() {
	os.RemoveAll(bdbPath)
}

func initBuntDBData() {
	for i := 0; i < 10000; i++ {
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
	setupBuntDB()
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
	now := time.Now()
	b.Log("run Benchmark_GetValue_BuntDB")
	setupBuntDB()
	initBuntDBData()
	b.ResetTimer()
	b.ReportAllocs()
	b.Log("run Benchmark_GetValue_BuntDB testing", time.Since(now).Seconds())

	for i := 0; i < b.N; i++ {

		bdb.View(func(tx *buntdb.Tx) error {
			_, err = tx.Get(string(GetKey(1000)))
			if err != nil {
				panic(errors.New("not found from cache"))
			}
			return nil
		})
	}
}
