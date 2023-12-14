package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {
	log.Println("standard logger")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	log.SetFlags(log.LstdFlags & 0)
	log.Println("no flags")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	mylog := log.New(os.Stdout, "my: ", 0)
	mylog.Println("from mylog")

	mylog.SetPrefix("oh my: ")
	mylog.Println("log")

	log.Default().SetPrefix("DEFAULT: ")
	log.Println("?")

	var bb bytes.Buffer
	buflog := log.New(&bb, "bb:", log.LstdFlags)
	buflog.Println("hello")
	fmt.Println("bb's guts:", bb.String())

	slogger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	slogger.Info("hi there")
	slogger.Info("hello again", "key", "val", "age", 25)
}
