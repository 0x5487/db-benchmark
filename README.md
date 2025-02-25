# Contrast Benchmark

Based on the results of a simple benchmark test of the open source kv database written in Golang.

## Test database:

- [flydb](https://github.com/ByteStorage/FlyDB)
- [bbolt](https://github.com/etcd-io/bbolt)
- [goleveldb](https://github.com/syndtr/goleveldb)
- [nutsdb](https://github.com/nutsdb/nutsdb)
- [rosedb](https://github.com/flower-corp/rosedb)
- [badger](https://github.com/dgraph-io/badger)
- [pebble](https://github.com/cockroachdb/pebble)

## Options:

- Value size: 512 bytes

## Results

```
goos: linux
goarch: amd64
pkg: contrast-benchmark
cpu: 11th Gen Intel(R) Core(TM) i7-11800H @ 2.30GHz

Benchmark_PutValue_FlyDB
Benchmark_PutValue_FlyDB-16        	   95023	     13763 ns/op	    2904 B/op	      16 allocs/op
Benchmark_GetValue_FlyDB
Benchmark_GetValue_FlyDB-16    	 	 2710143	     463.5 ns/op	     259 B/op	       5 allocs/op
Benchmark_PutValue_Badger
Benchmark_PutValue_Badger-16       	   59331	     22711 ns/op	    6006 B/op	      48 allocs/op
Benchmark_GetValue_Badger
Benchmark_GetValue_Badger-16       	  158686	      7686 ns/op	   10844 B/op	      42 allocs/op
Benchmark_PutValue_BoltDB
Benchmark_PutValue_BoltDB-16       	   32637	     56519 ns/op	   21009 B/op	     123 allocs/op
Benchmark_GetValue_BoltDB
Benchmark_GetValue_BoltDB-16       	  655971	     24327 ns/op	     723 B/op	      26 allocs/op 
Benchmark_PutValue_GoLevelDB
Benchmark_PutValue_GoLevelDB-16    	   71931	     14709 ns/op	    2226 B/op	      12 allocs/op
Benchmark_GetValue_GoLevelDB
Benchmark_GetValue_GoLevelDB-16    	  500736	      2520 ns/op	    1278 B/op	      15 allocs/op
Benchmark_PutValue_NutsDB
Benchmark_PutValue_NutsDB-16       	   78801	     13582 ns/op	    3242 B/op	      22 allocs/op
Benchmark_GetValue_NutsDB
Benchmark_GetValue_NutsDB-16       	  373124	      5702 ns/op	    1392 B/op	      14 allocs/op
Benchmark_PutValue_RoseDB
Benchmark_PutValue_RoseDB-16       	   69776	     19166 ns/op	    6242 B/op	      59 allocs/op
Benchmark_GetValue_RoseDB
Benchmark_GetValue_RoseDB-16       	 4155183	     298.0 ns/op	     167 B/op	       4 allocs/op
Benchmark_PutValue_Pebble
Benchmark_PutValue_Pebble-16       	   91304	     21877 ns/op	    2720 B/op	       8 allocs/op
Benchmark_GetValue_Pebble
Benchmark_GetValue_Pebble-16       	   66135	     15837 ns/op	   17193 B/op	      22 allocs/op
PASS
```

```
goos: linux
goarch: amd64
pkg: contrast-benchmark
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
Benchmark_PutValue_Badger
Benchmark_PutValue_Badger-16               45649             25159 ns/op            3451 B/op         48 allocs/op
Benchmark_GetValue_Badger
Benchmark_GetValue_Badger-16              163756             16241 ns/op            8415 B/op         41 allocs/op
Benchmark_PutValue_BoltDB
Benchmark_PutValue_BoltDB-16               10000            120696 ns/op           21038 B/op        123 allocs/op
Benchmark_GetValue_BoltDB
Benchmark_GetValue_BoltDB-16              686884             27959 ns/op             723 B/op         26 allocs/op
Benchmark_PutValue_FlyDB
Benchmark_PutValue_FlyDB-16               136382              8919 ns/op            2934 B/op         18 allocs/op
Benchmark_GetValue_FlyDB
Benchmark_GetValue_FlyDB-16              1832960               562.7 ns/op           319 B/op          6 allocs/op
Benchmark_PutValue_GoLevelDB
Benchmark_PutValue_GoLevelDB-16           102572             12261 ns/op            2198 B/op         12 allocs/op
Benchmark_GetValue_GoLevelDB
Benchmark_GetValue_GoLevelDB-16           446858              2947 ns/op            1277 B/op         15 allocs/op
Benchmark_PutValue_NutsDB
Benchmark_PutValue_NutsDB-16              104098             12266 ns/op            3305 B/op         24 allocs/op
Benchmark_GetValue_NutsDB
Benchmark_GetValue_NutsDB-16              317010              4813 ns/op            1424 B/op         15 allocs/op
Benchmark_PutValue_Pebble
Benchmark_PutValue_Pebble-16               93738             10675 ns/op            2105 B/op          8 allocs/op
Benchmark_GetValue_Pebble
Benchmark_GetValue_Pebble-16               55428             22125 ns/op           25700 B/op         33 allocs/op
Benchmark_PutValue_RoseDB
Benchmark_PutValue_RoseDB-16               82424             15541 ns/op            3750 B/op         22 allocs/op
Benchmark_GetValue_RoseDB
Benchmark_GetValue_RoseDB-16              404356              3144 ns/op            1301 B/op          9 allocs/op
PASS

```