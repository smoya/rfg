package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var numFiles, fileSizeBytes int
var outputDir string

func init() {
	flag.IntVar(&numFiles, "n", 10, "-n=<number-of-files> Number of files to generate. 10 by default.")
	flag.IntVar(&fileSizeBytes, "s", 100, "-s=<size-of-files-in-bytes> Size of each generated files in Bytes. 100 by default.")
	flag.StringVar(&outputDir, "o", "", "-o=<size-of-files-in-bytes> Output directory. Mandatory")
	flag.Parse()

	if outputDir == "" {
		fmt.Println("-o argument is mandatory.")
		os.Exit(1)
	}

	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err := os.Mkdir(outputDir, 0777); err != nil {
			log.Fatal(err.Error())
		}
	} else {
		log.Println("Warning: Output dir already exists. Content may be overwritten.")
	}
}

func main() {
	for i := 0; i < numFiles; i++ {
		err := generate(i + 1)
		if err != nil {
			log.Fatal(err.Error())
		}

	}
}

func generate(fileID int) error {
	f, err := os.Create(filepath.Join(outputDir, fmt.Sprintf("file%d", fileID)))
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write(randBytes(fileSizeBytes))
	return err
}

// From https://stackoverflow.com/a/31832326
func randBytes(n int) []byte {
	src := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return b
}
