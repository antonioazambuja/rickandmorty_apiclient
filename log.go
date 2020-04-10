package main

import (
	"log"
	"os"
)

// LogCharacter is used for logs in calls of Character endpoint
var LogCharacter = log.New(os.Stderr, "Character API: ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
// LogLocation is used for logs in calls of Location endpoint
var LogLocation  = log.New(os.Stderr, "Character API: ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
// LogEpisode is used for logs in calls of Episode endpoint
var LogEpisode  = log.New(os.Stderr, "Character API: ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
// LogDefault is used for logs in any operation
var LogDefault  = log.New(os.Stderr, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
