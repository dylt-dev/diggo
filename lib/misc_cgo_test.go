package lib

// #include <stdio.h>
import (
	"os"
	"testing"
)

func TestPutchar (t *testing.T) {
	Putchar()
}

func TestUnlink (t *testing.T) {
	path := "/tmp/hello.txt"
	os.Create(path)
	Unlink(path)
}