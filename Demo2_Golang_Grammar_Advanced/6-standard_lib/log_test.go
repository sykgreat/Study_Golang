package __standard_lib

import (
	"log"
	"os"
	"testing"
)

func init() {
	log.SetFlags(log.Llongfile | log.LstdFlags)
	log.SetOutput(os.Stdout)
}

func Test_log(t *testing.T) {
	log.Println("hello world")
	log.Printf("hello world, %s\n", "nick1")
	log.Fatalf("hello world, %s\n", "nick2")
	log.Panicf("hello world, %s\n", "nick3")
}
