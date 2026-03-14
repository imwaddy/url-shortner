package logger

import (
	"log"
	"os"
)

var logg *log.Logger

func Init() {
	logg = log.New(
		os.Stdout,
		"[URL-SHORTNER] ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)
}

func Println(v ...any) {
	logg.Println(v...)
}

func Printf(format string, v ...any) {
	logg.Printf(format, v...)
}

func Error(v ...any) {
	logg.Println(v...)
}

func Errorf(format string, v ...any) {
	logg.Printf(format, v...)
}

func Fatal(v ...any) {
	logg.Fatal(v...)
}

func Fatalf(format string, v ...any) {
	logg.Fatalf(format, v...)
}

func Panic(v ...any) {
	logg.Panic(v...)
}

func Panicf(format string, v ...any) {
	logg.Panicf(format, v...)
}
