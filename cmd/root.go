package cmd

import (
    "fmt"

    "github.com/AlecAivazis/survey/v2"
    "github.com/spf13/cobra"
)

var year string
var name string

var rootCmd = &cobra.Command{
    Use:   "license-gen",
    Short: "A CLI tool to generate licenses",
    Run: func(cmd *cobra.Command, args []string) {
        // Define available templates
        templates := []string{"MIT", "Apache"}

        // Prompt user to select a template
        var selectedTemplate string
        prompt := &survey.Select{
            Message: "Choose a template:",
            Options: templates,
        }
        err := survey.AskOne(prompt, &selectedTemplate)
        if err != nil {
            fmt.Println("Error selecting template:", err)
            return
        }

        // Generate the license file with the selected template
        if err := generateLicense(selectedTemplate, year, name); err != nil {
            fmt.Println("Error:", err)
        }
    },
}

func Execute() error {
    return rootCmd.Execute()
}

func init() {
    rootCmd.Flags().StringVarP(&year, "year", "y", "2024", "Year for the license")
    rootCmd.Flags().StringVarP(&name, "name", "n", "Your Name", "Name for the license")
}
