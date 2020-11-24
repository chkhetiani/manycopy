package cdn

import (
	"manycopy/src/server"
	"manycopy/src/settings"
	"net/http"
)



func init() {
	ser := server.Server {Name: "cdn" }
	fserver := http.StripPrefix("/cdn/", http.FileServer(http.Dir(settings.Get("content_path").(string))))
	ser.AddHandler("/cdn/", fserver)
	ser.Init()
}