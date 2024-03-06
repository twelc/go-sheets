package lib

type Config struct {
	credentials string
	sheetid     string
	table_name  string
	table_range string
}

func GetConfig(credentials string, sheetid string, table_name string) Config {
	return Config{
		credentials: credentials,
		sheetid:     sheetid,
		table_name:  table_name,
	}
}
