// package cmd

// import (
//     "fmt"
//     "os"
//     "text/template"

//     "github.com/spf13/cobra"
// )

// var licenseType string

// var generateCmd = &cobra.Command{
//     Use:   "generate",
//     Short: "Generate a license file",
//     Run: func(cmd *cobra.Command, args []string) {
//         generateLicense(licenseType)
//     },
// }

// func init() {
//     rootCmd.AddCommand(generateCmd)
//     generateCmd.Flags().StringVarP(&licenseType, "type", "t", "", "Type of license (e.g., MIT, Apache)")
//     generateCmd.MarkFlagRequired("type")
// }

// func generateLicense(licenseType string) {
//     tmpl, err := template.ParseFiles(fmt.Sprintf("templates/%s.tmpl", licenseType))
//     if err != nil {
//         fmt.Println("Error loading template:", err)
//         return
//     }

//     data := struct {
//         Year string
//         Name string
//     }{
//         Year: "2024",
//         Name: "Your Name",
//     }

//     file, err := os.Create("LICENSE")
//     if err != nil {
//         fmt.Println("Error creating file:", err)
//         return
//     }
//     defer file.Close()

//     tmpl.Execute(file, data)
//     fmt.Println("License file generated successfully")
// }

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

var generateCmd = &cobra.Command{
    Use:   "generate",
    Short: "Generate a license file",
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

func init() {
    rootCmd.AddCommand(generateCmd)
    generateCmd.Flags().StringVarP(&year, "year", "y", "2024", "Year for the license")
    generateCmd.Flags().StringVarP(&name, "name", "n", "Your Name", "Name for the license")
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

