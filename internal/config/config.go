package config

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) Print() {
	println("DBUrl: " + c.DBUrl)
	println("CurrentUserName: " + c.CurrentUserName)
}
