package main

import (
	"log"
	"os"
)

func main() {
	// Package log implements simple struct, methods & functions to log program runtime info

	// Why not use package fmt?
	// 1) Package log is safe from concurrent goroutines while plain fmt isn't
	// 2) Log can attach information automatically, such as time, date, file path, etc.

	log.Println("Sup Ninjas!")
	log.SetFlags(log.Ldate | log.Lshortfile)
	log.Println("Sup Ninjas!")

	//log.Panic("Panicking...")
	//log.Fatal("Uh-oh...")

	// Output a file
	file, _ := os.Create("file.log")
	log.SetOutput(file)
	log.Println("Hello World!")
	file.Close()
	log.SetOutput(os.Stdout)
	log.Println("Printing into standard out again")

	// Common loggers
	flags := log.LstdFlags | log.Lshortfile
	infoLogger := log.New(os.Stdout, "INFO: ", flags)
	warnLogger := log.New(os.Stdout, "WARN: ", flags)
	errorLogger := log.New(os.Stdout, "ERROR: ", flags)
	infoLogger.Println("This is an info log")
	warnLogger.Println("This is a warning log")
	errorLogger.Println("This is an error log")

	// You can also aggregate all three into one
	al := aggregatedLogger{
		infoLogger:  infoLogger,
		warnLogger:  warnLogger,
		errorLogger: errorLogger,
	}
	al.info("This is an info log")
	al.warn("This is a warning log")
	al.error("This is an error log")
}

type aggregatedLogger struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
}

func (l *aggregatedLogger) info(v ...interface{}) {
	l.infoLogger.Println(v...)
}

func (l *aggregatedLogger) warn(v ...interface{}) {
	l.warnLogger.Println(v...)
}

func (l *aggregatedLogger) error(v ...interface{}) {
	l.errorLogger.Println(v...)
}
