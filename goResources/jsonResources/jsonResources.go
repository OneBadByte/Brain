package jsonResources

import(
	"fmt"
	"encoding/json"
	"io/ioutil"
)

func TestingJsonWorks(){
	fmt.Println("json works")
}

//Writes json to a file.
func WriteToJsonFile(fileName string, data interface{}) error{
	marshaledData, err := json.Marshal(data)
	err = ioutil.WriteFile(fileName, marshaledData, 0775)
	if err != nil{
		fmt.Println("couldn't read write to file")
	}
	return err
}

//Reads json from a file.
func ReadFromJsonFile(fileName string, data interface{}) error{
	marshaledData, err := ioutil.ReadFile(fileName)
	err = json.Unmarshal(marshaledData, &data)
	if err != nil{
		fmt.Println("couldn't read write to file")
	}
	return err
}

