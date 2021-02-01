package makedbdir

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
)

func MakeDir() (string, error) {

	home, _ := homedir.Dir()

	homeDir, err := ioutil.ReadDir(home)
	if err != nil {
		log.Fatal("Cant locate home dir: ", err)
		return "", err
	}

	pathToFolder := fmt.Sprintf("%s/%s", home, "German")

	for _, f := range homeDir {
		if "German" != f.Name() {
			err := os.Mkdir(pathToFolder, 0755)
			if err != nil {
				log.Println(pathToFolder, err)
				return pathToFolder, nil
			}

			return pathToFolder, nil
		}
	}

	return pathToFolder, nil
}
