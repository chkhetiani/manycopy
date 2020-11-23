package settings


var data = make(map[string]interface{})

func init() {
	data["api_port"] = ":1060"
	data["static_port"] = ":1061"
	data["content_port"] = ":1062"
	data["content_path"] = "/dev/ManyCopy/content"
	data["static_path"] = "/dev/ManyCopy/static"
}


func Get(key string) interface{} {
	return data[key]
}