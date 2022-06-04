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

	LoadDotEnv()

	var planFile string
	var outputFile string
	SetupArgs(&planFile, &outputFile)

	SetupEventListeners()

	var plan = engine.LoadPlanSchemaFromPath(planFile)
	var report = engine.Run(plan)
	fmt.Printf("DEBUG: %+v\n", report)
}
