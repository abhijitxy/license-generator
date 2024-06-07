package cmd

import (
    "fmt"

    "github.com/AlecAivazis/survey/v2"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "license-gen",
    Short: "A CLI tool to generate licenses",
    Run: func(cmd *cobra.Command, args []string) {
        var year string
        var name string

        // Define available templates
        templates := []string{"MIT", "Apache"}

        // Prompt user to select a template
        var selectedTemplate string
        templatePrompt := &survey.Select{
            Message: "Choose a template:",
            Options: templates,
        }
        err := survey.AskOne(templatePrompt, &selectedTemplate)
        if err != nil {
            fmt.Println("Error selecting template:", err)
            return
        }

        // Prompt user for the year
        yearPrompt := &survey.Input{
            Message: "Enter the year:",
            Default: "2024",
        }
        err = survey.AskOne(yearPrompt, &year)
        if err != nil {
            fmt.Println("Error entering year:", err)
            return
        }

        // Prompt user for the name
        namePrompt := &survey.Input{
            Message: "Enter your name:",
            Default: "Your Name",
        }
        err = survey.AskOne(namePrompt, &name)
        if err != nil {
            fmt.Println("Error entering name:", err)
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
    // No longer setting default values for year and name
}
