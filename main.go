//go:generate statik -src=./public

package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
)

// Install : allows to install package
func Install(pkg string) {
	var cmd *exec.Cmd
	cmd = exec.Command("choco", "install", pkg, "-y")
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

		Install(chocoPackage)

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

}
