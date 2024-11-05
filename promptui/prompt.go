package promptui

import (
	"log"

	"github.com/manifoldco/promptui"
)

func PromptUser(col1, col2 map[string]int) (string, string) {

	col1Slice := []string{}
	for v, _ := range col1 {
		col1Slice = append(col1Slice, v)
	}
	promptOne := promptui.Select{
		Label: "Column Name for csv 1",
		Items: col1Slice,
	}

	col2Slice := []string{}
	for v, _ := range col2 {
		col2Slice = append(col2Slice, v)
	}

	promptTwo := promptui.Select{
		Label: "Column Name for csv 2",
		Items: col2Slice,
	}

	_, resOne, err := promptOne.Run()
	_, resTwo, err := promptTwo.Run()

	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	return resOne, resTwo

}
