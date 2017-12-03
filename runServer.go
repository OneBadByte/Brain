package main

import(
	"fmt"
	"./goResources/jsonResources"
	"./goResources/dataResources"
)

func CheckServerConfigFile() dataResources.ServerConfigFile{
	serverConfigFile := dataResources.ServerConfigFile{}
	err := jsonResources.ReadFromJsonFile(dataResources.ConfigFileLocation, &serverConfigFile)
	if err != nil{
		fmt.Println("failed to find config file.\n Creating new one!!!")
		newConfigFile := dataResources.ServerConfigFile{
			Name: "Brain",
			TcpPort: 3000,
		}
		jsonResources.WriteToJsonFile(dataResources.ConfigFileLocation, newConfigFile)
		jsonResources.ReadFromJsonFile(dataResources.ConfigFileLocation, &serverConfigFile)
	}
	return serverConfigFile
}


func main(){
	serverInfo := CheckServerConfigFile()
	fmt.Println("Hello I am", serverInfo.Name)
	for{
		break
	}
}
