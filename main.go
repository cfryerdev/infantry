package infantry

import (
	"fmt"
	"infantry/cli/cmd"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting Infantry!")

	var plan string
	var output string

	err := cmd.SetupArgs(plan, output).Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
