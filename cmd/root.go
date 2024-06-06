package cmd

import (
    "fmt"
    "os"
    "path/filepath"
    "text/template"

    "github.com/AlecAivazis/survey/v2"
    "github.com/spf13/cobra"
)

var year string
var name string

var rootCmd = &cobra.Command{
    Use:   "license-gen",
    Short: "A CLI tool to generate licenses",
    Run: func(cmd *cobra.Command, args []string) {
        // Get available templates
        templates, err := getTemplates("templates")
        if err != nil {
            fmt.Println("Error retrieving templates:", err)
            return
        }

        // If no templates found, return an error
        if len(templates) == 0 {
            fmt.Println("No templates found.")
            return
        }

        // Prompt user to select a template
        var selectedTemplate string
        prompt := &survey.Select{
            Message: "Choose a template:",
            Options: templates,
        }
        err = survey.AskOne(prompt, &selectedTemplate)
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

func getTemplates(dir string) ([]string, error) {
    var templates []string
    err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() && filepath.Ext(path) == ".tmpl" {
            templates = append(templates, filepath.Base(path))
        }
        return nil
    })
    return templates, err
}

func generateLicense(templateFile, year, name string) error {
    tmpl, err := template.ParseFiles(fmt.Sprintf("templates/%s", templateFile))
    if err != nil {
        return fmt.Errorf("loading template: %w", err)
    }

    data := struct {
        Year string
        Name string
    }{
        Year: year,
        Name: name,
    }

    file, err := os.Create("LICENSE")
    if err != nil {
        return fmt.Errorf("creating file: %w", err)
    }
    defer file.Close()

    if err := tmpl.Execute(file, data); err != nil {
        return fmt.Errorf("executing template: %w", err)
    }

    fmt.Println("License file generated successfully")
    return nil
}
