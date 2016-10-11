package utils

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"io"
	"os"
)

// Copy and Write big file
func TestFiles() {
	dir, _ := os.Getwd()
	fmt.Printf("Current Dir: %s\n", dir)
	fileRead := "some_big_file"
	fileSave := "some_new_big_file"
	if _, err := os.Stat(fileRead); err == nil {
		fr, err := os.Open(fileRead)
		if err != nil {
			log.Fatalf("Cant Open file to read: %s", fileRead)
		}
		defer fr.Close()

		fw, err := os.Create(fileSave)
		if err != nil {
			log.Fatalf("Cant Open file to save: %s", fileSave)
		}
		defer fw.Chdir()

		rb := make([]byte, 4096)
		for {
			cnt, err := fr.Read(rb)
			if err != nil && err != io.EOF {
				log.Fatalf("Failig until read: %#v", err)
			}
			if cnt == 0 {
				break
			}
			if _, err := fw.Write(rb[:cnt]); err != nil {
				log.Fatalf("Failig until write: %#v", err)
			}
		}
	} else {
		log.Fatalf("File not exist: %s %v", fileRead, err)
	}
}
