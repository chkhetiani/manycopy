package settings

import "os"

var data = make(map[string]interface{})

func init() {
	data["api_port"] = ":1060"
	data["cdn_port"] = ":1061"
	data["static_port"] = ":1062"
	data["content_path"] = "/dev/many-copy/cdn"
	data["static_path"] = (map[bool]string{true:  "/dev/many-copy/static", false: "/app/static"})[os.Getenv("PORT") == ""]
}

func Get(key string) interface{} {
	return data[key]
}