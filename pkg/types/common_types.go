package common_types

type DBCreds struct {
	DBHost     string `json:"database_host"`
	DBPort     int    `json:"database_port"`
	DBUsername string `json:"database_username"`
	DBPassword string `json:"database_password"`
}

type FileTypes struct {
	Json string
	Yaml string
	Xml  string
	Csv  string
	Pdf  string
	Text string
}
