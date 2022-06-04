package console

import (
	"fmt"
	"infantry/engine"
)

func main() {
	Execute()
}

func Execute() {
	fmt.Println("Thank you for using Infantry!")
	fmt.Println("    .--._____,")
	fmt.Println(" .-='=='==-,  ")
	fmt.Println("(O_o_o_o_o_O) ")
	fmt.Println("-----------------------------")

	var planFile string
	var outputFile string

	SetupArgs(&planFile, &outputFile)

	LoadDotEnv()
	
	SetupEventListeners()

	var plan = engine.LoadPlanSchemaFromPath(planFile)
	var report = engine.Run(plan)

	fmt.Printf("DEBUG: %+v\n", report)
}
