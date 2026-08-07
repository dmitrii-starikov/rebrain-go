package config

func (c Config) StructToMap() map[string]interface{} {
	result := make(map[string]interface{})
	result["Name"] = c.Name
	result["app_host"] = c.Host
	result["app_port"] = c.Port
	result["environment"] = c.Environment
	result["volumes"] = c.Volumes

	return result
}
