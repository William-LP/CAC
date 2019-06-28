//go:generate statik -src=./public

package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
)

func DownloadFile(filepath string, url string) error {

    // Get the data
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // Create the file
    out, err := os.Create(filepath)
    if err != nil {
        return err
    }
    defer out.Close()

    // Write the body to file
    _, err = io.Copy(out, resp.Body)
    return err
}


// Install : allows to install package
func Install(mode, pkg string) {
	var cmd *exec.Cmd
	switch(mode) {
	case "choco" :
		cmd = exec.Command("choco", "install", pkg, "-y")
	case "npm" :
		cmd = exec.Command("npm", "install", "-g", pkg)
	}

	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	print(string(stdout))
}

func main() {

	listeChocoPackage, _ := os.Open("./choco_packages.txt")

	scanner := bufio.NewScanner(listeChocoPackage)
	for scanner.Scan() {
		chocoPackage := scanner.Text()
		log.Println("install :",chocoPackage)

		Install("choco", chocoPackage)

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	Install("npm", "ungit")

	fileURL := "https://golangcode.com/images/avatar.jpg"

    if err := DownloadFile("avatar.jpg", fileURL); err != nil {
        panic(err)
    }
}
