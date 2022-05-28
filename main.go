package infantry

import (
	"fmt"
	"infantry/cmd/console"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting Infantry!")

	var plan string
	var output string

	err := console.SetupArgs(plan, output).Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
