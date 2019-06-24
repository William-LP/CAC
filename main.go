package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
)


func shell(command string) {

	cmd := exec.Command(command)

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
		chocoLine := "choco install -y" + chocoPackage

		shell(chocoLine)

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	shell("npm install -g ungit")
}
