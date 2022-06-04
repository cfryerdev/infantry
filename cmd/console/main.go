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

	var report, err = engine.Run(plan)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	engine.SaveReportFile(report, outputFile)
}
