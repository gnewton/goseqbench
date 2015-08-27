package main

import (
	"encoding/binary"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func BenchmarkFileWriteSync(b *testing.B) {
	writeSequence(true)
}

func BenchmarkFileWriteNoSync(b *testing.B) {
	writeSequence(false)
}

func writeSequence(flush bool) {
	file, err := ioutil.TempFile(os.TempDir(), "write")
	defer os.Remove(file.Name())
	if err != nil {
		log.Fatal("Unable to create temporary file")
	}
	var i uint64

	for i = 0; i < maxSequenceValue; i++ {
		_, err = file.Seek(0, 0)
		if err != nil {
			log.Fatal("Failed file seek")
		}
		binary.Write(file, binary.BigEndian, i)
		if flush {
			err = file.Sync()
			if err != nil {
				log.Fatal("Failed file sync")
			}
		}
	}

}
