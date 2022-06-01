package console

import (
	"fmt"
	"infantry/engine"
	"log"
	"os"
)

func main() {
	Execute()
}

func Execute() {
	var planFile string
	var outputFile string
	err := SetupArgs(&planFile, &outputFile).Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Thank you for using Infantry!")
	fmt.Println("    .--._____,")
	fmt.Println(" .-='=='==-,  ")
	fmt.Println("(O_o_o_o_o_O) ")
	fmt.Println("====================================")
	fmt.Println(fmt.Sprintf("Plan: %s", planFile))
	fmt.Println(fmt.Sprintf("Output: %s", outputFile))

	SetupEventListeners()

	var plan = engine.LoadPlanSchemaFromPath(planFile)
	fmt.Printf("%+v\n", plan)
	engine.Run(plan, outputFile)
}
