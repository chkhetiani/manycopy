package settings

import "os"

var data = make(map[string]interface{})

func init() {
	data["api_port"] = ":1060"
	data["static_path"] = (map[bool]string{true:  "/dev/many-copy/static", false: "/app/static"})[os.Getenv("PORT") == ""]
}

func Get(key string) interface{} {
	return data[key]
}