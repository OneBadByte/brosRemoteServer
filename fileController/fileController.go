package fileController

import(
	"io/ioutil"
	"os"
)

func ReadFile(fileName string)(string){
	fileInfo, _ := ioutil.ReadFile(fileName)
	return string(fileInfo)
}

func WriteToFile(fileName string, text string){
	CreateFile(fileName)
	ioutil.WriteFile(fileName, []byte(text), 0777)
}

func RemoveFile(fileName string){
	os.Remove(fileName)
}

func CreateFile(fileName string){
	_, err := os.Open(fileName)
	if err != nil{
		os.Create(fileName)
	}
}

