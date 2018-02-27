package ginit

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os"
)

type GoInitConfig struct {
	goPackageDirectory string
}

func runInstallation(config GoInitConfig) error {
	fmt.Println("Welcome to ginit!\n We just need to set up a few configuration values to start")
	fmt.Println("What is the directory that you would like to create your packages (i.e. $GOPATH/github.com/username?")
	fmt.Scanf("%s\n", &config.goPackageDirectory)

	fp, err := os.Create("~/.ginit.json")

	if err != nil {
		fmt.Println("Sorry I encontered an error, please send this stack trace to my github repo")
		fmt.Println(err.Error())
		return err
	}
	
	defer fp.Close()

	b, err := json.Marshal(config)

	if err != nil {
		fmt.Println("Sorry I encontered an error, please send this stack trace to my github repo")
		fmt.Println(err.Error())
		return err
	}
	
	fp.Write(b)
	return nil
}

func main() {
	var config GoInitConfig

	if _, err := os.Stat("~/.ginit.json"); os.IsNotExist(err) {
		err = runInstallation(config)

		if err != nil {
			fmt.Print("Ending execution")
			return
		}
	}

	byteArray, err := ioutil.ReadFile("~/.ginit.json")

	if err != nil {
		fmt.Println("Sorry I encontered an error, please send this stack trace to my github repo")
		fmt.Println(err.Error())
	}

	err = json.Unmarshal(byteArray, &config)

	if err != nil {
		fmt.Println("Sorry I encontered an error, please send this stack trace to my github repo")
		fmt.Println(err.Error())
	}

	
}
