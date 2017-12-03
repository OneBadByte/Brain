package dataResources

// servers config file location.
var ConfigFileLocation string = "./json/config.json"

//Users info json file location.
var UserInfoFileName string = "./json/testing.json"

//users personal information
type PersonalInfo struct{
	Name string `json:"name"`
	Age int `json:"age"`
}

//Used as the servers config information
type ServerConfigFile struct{
	Name string `json:"name"`
	TcpPort int `json:"tcpPort"`
}
