package main

import (
	"fmt"
	"os"
)

func LoadApps() {
	fmt.Println(os.ReadDir("../apps")) // todo: read apps/*/manifest.xml or similar
}
