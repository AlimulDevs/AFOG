package makefolder

import (
	"fmt"
	"os"
)

var mainFolderName = []string{"app", "controllers", "config", "helper", "models", "repositories", "services", "storage"}

func CreateMainFolder() {
	for _, data := range mainFolderName {

		// Membuat folder target jika belum ada
		err := os.MkdirAll(data, 0755)
		if err != nil {
			fmt.Println("Gagal membuat folder:", err)
			return
		}
		fmt.Println("success create folder", data)
	}
}
