package cmd

import (
    "fmt"
    "os"
    "text/template"

    "github.com/spf13/cobra"
)

var licenseType string

var generateCmd = &cobra.Command{
    Use:   "generate",
    Short: "Generate a license file",
    Run: func(cmd *cobra.Command, args []string) {
        generateLicense(licenseType)
    },
}

func init() {
    rootCmd.AddCommand(generateCmd)
    generateCmd.Flags().StringVarP(&licenseType, "type", "t", "", "Type of license (e.g., MIT, Apache)")
    generateCmd.MarkFlagRequired("type")
}

func generateLicense(licenseType string) {
    tmpl, err := template.ParseFiles(fmt.Sprintf("templates/%s.tmpl", licenseType))
    if err != nil {
        fmt.Println("Error loading template:", err)
        return
    }

    data := struct {
        Year string
        Name string
    }{
        Year: "2024",
        Name: "Your Name",
    }

    file, err := os.Create("LICENSE")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    tmpl.Execute(file, data)
    fmt.Println("License file generated successfully")
}
