/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"cli-app/data"
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"os"

	"github.com/spf13/cobra"
)
// cobra add new -p 'cmd'

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new note",
	Long: `Create a new note`,
	Run: func(cmd *cobra.Command, args []string) {
		createNewNote()
	},
}

func init() {
// new is subcommand of note
	noteCmd.AddCommand(newCmd)

}

type promptContent  struct {
	errorMsg string
	label string
}
/*
prompt input mode:
validate func
templates
prompt
*/
func promptGetInput(pc promptContent) string{
	validate := func(input string)error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		return  nil
	}
	prompt := promptui.Prompt{
		Label:pc.label,
        Validate:validate,
	}
 result , err := prompt.Run()
 if err!=nil{
 	fmt.Printf("prompt failed %v\n", err)
 	os.Exit(1)
 }
 return  result
}

func promptGetSelect(pc promptContent) string {
	items := []string{"food", "study", "Games", "object"}
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.label,
			Items:    items,
			AddLabel: "Other",
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}

func createNewNote() {
	wordPromptContent := promptContent{
		"Please provide a word.",
		"What word would you like to make a note of?",
	}
	// capture input from user: promptGetInput
	word := promptGetInput(wordPromptContent)

	definitionPromptContent := promptContent{
		"Please provide a definition.",
		fmt.Sprintf("What is the definition of %s?", word),
	}
	definition := promptGetInput(definitionPromptContent)
	fmt.Sprintf("What is the definition of %s?",definition)

	categoryPromptContent := promptContent{
		"Please provide a category.",
		fmt.Sprintf("What category does %s belong to?", word),
	}
	category := promptGetSelect(categoryPromptContent)

	data.InsertNote(word, definition, category)
}