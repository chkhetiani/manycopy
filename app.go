package main

import (
	_ "manycopy/src/server/api"
	_ "manycopy/src/server/cdn"
	_ "manycopy/src/server/static"
)


func main() {
	select {}
}