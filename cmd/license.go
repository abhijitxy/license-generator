package cmd

import (
    "fmt"
    "os"
)

func generateLicense(template, year, name string) error {
    var licenseText string

    switch template {
    case "MIT":
        licenseText = fmt.Sprintf(mitTemplate, year, name)
    case "Apache":
        licenseText = fmt.Sprintf(apacheTemplate, year, name)
    default:
        return fmt.Errorf("unknown template: %s", template)
    }

    file, err := os.Create("LICENSE")
    if err != nil {
        return fmt.Errorf("creating file: %w", err)
    }
    defer file.Close()

    if _, err := file.WriteString(licenseText); err != nil {
        return fmt.Errorf("writing to file: %w", err)
    }

    fmt.Println("License file generated successfully")
    return nil
}

const mitTemplate = `MIT License

Copyright (c) %s %s

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.`

const apacheTemplate = `Apache License 2.0

Copyright (c) %s %s

Licensed under the Apache License, Version 2.0 (the "License"); you may not use
this file except in compliance with the License. You may obtain a copy of the
License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed
under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
CONDITIONS OF ANY KIND, either express or implied. See the License for the
specific language governing permissions and limitations under the License.`
