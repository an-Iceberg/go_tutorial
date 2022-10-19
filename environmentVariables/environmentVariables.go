package environmentVariables

import (
	"fmt"

	"github.com/joho/godotenv"
)

func Go() {
	fmt.Println("  Environment Variables")

	var local_environment_variables map[string]string
	local_environment_variables, error := godotenv.Read("environmentVariables/.env")

	if error != nil {
		fmt.Println("Could not load .env file")
	}

	for key, value := range local_environment_variables {
		fmt.Printf("%s: %s\n", key, value)
	}
}
