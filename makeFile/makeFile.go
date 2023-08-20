package makefile

import (
	"fmt"
	templetefile "golang/templeteFile"
	"os"
	"path/filepath"
	"text/template"
)

var mainFileName []string = []string{"go.mod", ".env"}
var mainFileContent []string = []string{templetefile.GoModTemplete, templetefile.EnvTemplete}
var helperFileName []string = []string{"hashPassword.go", "makeUuid.go", "util.go"}
var helperFileContent []string = []string{templetefile.HashPasswordTemplete, templetefile.MakeUuidTemplete, templetefile.UtilTemplete}

func CreateMainFile() {
	for i := 0; i < len(mainFileContent); i++ {
		// Buat file baru
		file, err := os.Create(mainFileName[i])
		if err != nil {
			panic(err)
		}
		defer file.Close()
		// Buat template
		t := template.Must(template.New("mytemplate").Parse(mainFileContent[i]))

		// Isi template dengan data dan tulis ke file
		err = t.Execute(file, "data")
		if err != nil {
			panic(err)
		}
		fmt.Println("success create file", mainFileName[i])
	}
}

func CreateHelperFile() {
	for i := 0; i < len(helperFileName); i++ {
		// Buat file baru
		filePath := filepath.Join("helper", helperFileName[i])
		file, err := os.Create(filePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		// Buat template
		t := template.Must(template.New("mytemplate").Parse(helperFileContent[i]))

		// Isi template dengan data dan tulis ke file
		err = t.Execute(file, "data")
		if err != nil {
			panic(err)
		}
		fmt.Println("success create file", helperFileName[i])
	}
}
