package main

import (
	"fmt"
	"github.com/edsrzf/mmap-go"
	"io/ioutil"
	"log"
	"os"
	"syscall"
	"testing"
	"unsafe"
)

const maxSequenceValue = 50000000

func BenchmarkMmapSync(b *testing.B) {
	flush := true
	mapSequence(flush)
}

func BenchmarkMmapNoSync(b *testing.B) {
	flush := false
	mapSequence(flush)

}

func mapSequence(flush bool) {
	mfile, err := ioutil.TempFile(os.TempDir(), "mmap")
	defer os.Remove(mfile.Name())
	if err != nil {
		log.Fatal("Unable to create temporary file")
	}
	// Write 8 bytes
	_, err = mfile.Write([]byte("01234567"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	mm, err := mmap.MapRegion(mfile, 4096, mmap.RDWR, syscall.PROT_WRITE, 0)

	v := (*uint64)(unsafe.Pointer(&mm[0]))

	var i uint64
	for i = 0; i < maxSequenceValue; i++ {
		*v = *v + 1
		if flush {
			mm.Flush()
		}
	}
	err = mm.Unmap()
	if err != nil {
		log.Fatal("Failed unmap")
	}
}
