package static
//
//import (
//	"manycopy/src/server"
//	"manycopy/src/settings"
//	"net/http"
//)
//
//func init() {
//	ser := server.Server {Name: "static" }
//	fserver  := http.FileServer(http.Dir(settings.Get("static_path").(string)))
//	ser.AddHandler("/", fserver)
//	ser.Init()
//}