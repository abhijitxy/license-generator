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
    "text/template"

    "github.com/spf13/cobra"
)

var (
    licenseType string
    year        string
    name        string
)

var generateCmd = &cobra.Command{
    Use:   "generate",
    Short: "Generate a license file",
    Run: func(cmd *cobra.Command, args []string) {
        if err := generateLicense(licenseType, year, name); err != nil {
            fmt.Println("Error:", err)
        }
    },
}

func init() {
    rootCmd.AddCommand(generateCmd)
    generateCmd.Flags().StringVarP(&licenseType, "type", "t", "", "Type of license (e.g., MIT, Apache)")
    generateCmd.Flags().StringVarP(&year, "year", "y", "2024", "Year for the license")
    generateCmd.Flags().StringVarP(&name, "name", "n", "Your Name", "Name for the license")
    generateCmd.MarkFlagRequired("type")
}

func generateLicense(licenseType, year, name string) error {
    tmpl, err := template.ParseFiles(fmt.Sprintf("templates/%s.tmpl", licenseType))
    if (err != nil) {
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
    if (err != nil) {
        return fmt.Errorf("creating file: %w", err)
    }
    defer file.Close()

    if err := tmpl.Execute(file, data); err != nil {
        return fmt.Errorf("executing template: %w", err)
    }

    fmt.Println("License file generated successfully")
    return nil
}
