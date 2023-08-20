package main

import (
	makefile "golang/makeFile"
	makefolder "golang/makeFolder"
)

func main() {

	makefolder.CreateMainFolder()
	makefile.CreateMainFile()
	makefile.CreateHelperFile()

}
