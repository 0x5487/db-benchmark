package contrast_benchmark

import (
	"os"
	"testing"

	"github.com/nutsdb/nutsdb"
)

var nutsDB *nutsdb.DB

func setupNutsDB() {
	dirPath := "benchmark/nutsdb"
	os.RemoveAll(dirPath)

	opts := nutsdb.DefaultOptions
	opts.Dir = dirPath
	opts.SyncEnable = false
	opts.SegmentSize = 1 * nutsdb.GB
	opts.RWMode = nutsdb.MMap
	opts.EntryIdxMode = nutsdb.HintKeyValAndRAMIdxMode
	var err error
	nutsDB, err = nutsdb.Open(opts)
	if err != nil {
		panic(err)
	}

}

func initNutsDBData() {
	for i := 0; i < 10000; i++ {
		nutsDB.Update(func(tx *nutsdb.Tx) error {
			err := tx.Put("test-bucket", GetKey(i), GetValue(), 0)
			if err != nil {
				panic(err)
			}
			return nil
		})
	}
}

func Benchmark_PutValue_NutsDB(b *testing.B) {
	setupNutsDB()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		nutsDB.Update(func(tx *nutsdb.Tx) error {
			err := tx.Put("test-bucket", GetKey(i), GetValue(), 0)
			if err != nil {
				panic(err)
			}
			return nil
		})
	}
}

func Benchmark_GetValue_NutsDB(b *testing.B) {
	setupNutsDB()
	initNutsDBData()
	b.ResetTimer()
	b.ReportAllocs()

	nutsDB.Merge()

	for i := 0; i < b.N; i++ {
		nutsDB.View(func(tx *nutsdb.Tx) error {
			_, err := tx.Get("test-bucket", GetKey(1000))
			if err != nil && err != nutsdb.ErrKeyNotFound {
				panic(err)
			}
			return nil
		})
	}
}
