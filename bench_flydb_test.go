package contrast_benchmark

import (
	"os"
	"path/filepath"

	"github.com/ByteStorage/FlyDB/config"
	"github.com/ByteStorage/FlyDB/engine"
	"github.com/ByteStorage/FlyDB/flydb"
)

var FlyDB *engine.DB
var err error

func setupFlyDB() {
	opts := config.DefaultOptions
	opts.DirPath = filepath.Join("benchmark", "flydb")

	FlyDB, err = flydb.NewFlyDB(opts)
	if err != nil {
		panic(err)
	}
}

func cleanupFlyDB() {
	os.RemoveAll(dirPath)
}

// func Benchmark_PutValue_FlyDB(b *testing.B) {
// 	setupFlyDB()

// 	b.ResetTimer()
// 	b.ReportAllocs()

// 	for n := 0; n < b.N; n++ {
// 		err = FlyDB.Put(GetKey(1), GetValue())
// 		assert.Nil(b, err)
// 	}

// 	cleanupFlyDB()
// }

// func Benchmark_GetValue_FlyDB(b *testing.B) {
// 	setupFlyDB()

// 	for i := 0; i < 500000; i++ {
// 		err = FlyDB.Put(GetKey(i), GetValue())
// 		assert.Nil(b, err)
// 	}

// 	b.ResetTimer()
// 	b.ReportAllocs()

// 	for n := 0; n < b.N; n++ {
// 		_, err = FlyDB.Get(GetKey(n))
// 		if err != nil && err != _const.ErrKeyNotFound {
// 			panic(err)
// 		}
// 	}

// 	cleanupFlyDB()
// }
