
###Persistent, consistent sequence###

I was interested in creating a persistent, *consistent* sequence generator for a uint64 for Go.

I thought that using a (memory mapped)[https://en.wikipedia.org/wiki/Memory-mapped_file] file would be a good idea.
I have never used mmap in any language, on any system, but it seemed to be a reasonable way to go.

I do want a sequence that guarantees that it persists (*synced*) _before_ returned to the caller.
But I also wanted to explore different ways of doing this, and the impact of syncing or not on performance.

So I've looked at writing to a file (with binary.Write) and using mmap (with the (mmap)[github.com/edsrzf/mmap-go] package], each with and without syncing.

Of course, for my use case, syncing is necessary.
But I wanted to see what the performance differences were.

###Benchmark results###
n
```
$ go test -bench=.
testing: warning: no tests to run
PASS
BenchmarkMmapSync	       1	28622729995 ns/op
BenchmarkMmapNoSync	2000000000	         0.17 ns/op
BenchmarkFileWriteSync	       1	256338155036 ns/op
BenchmarkFileWriteNoSync	       1	227634544034 ns/op
ok  	github.com/gnewton/goseqbench	517.653s
$

```

Clearly mmap-no-sync is hugely fast, but not useful for me.

mmap-sync versus filewrite-sync is an order of magnatude difference.

Unless I made some mistake here, it looks like mmap is the way to go!

