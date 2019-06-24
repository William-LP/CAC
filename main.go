package main

import (

"os/exec"
"bufio"
"log"
"os"

)

func main() {
	choco := "choco"
	option := "install"
	param := "-y"


    listePackage, err := os.Open("./package.txt")
	if err != nil {
		log.Fatal(err)
	}
    defer listePackage.Close()

    scanner := bufio.NewScanner(listePackage)
	for scanner.Scan() {
		chocoPackage := scanner.Text()
		cmd := exec.Command(choco, option, chocoPackage ,param)
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	print(string(stdout))
}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}


