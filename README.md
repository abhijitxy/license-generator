<h1 align="center">License Generator</h1>

### Overview

This is a blazing fast ⚡, command line license generator for your open source projects written in Go.

I know that GitHub has a great GUI to add licenses to projects but I always found myself doing too much work. First, you have to go to GitHub, create a file, type 'LICENSE', pick a license, push it, and then pull it locally. With this, you can just generate the license locally and push it to GitHub.

### Usage

```bash
license-gen
```

### Contributing

- Fork the repository
- Create a branch
- Install dependencies
```bash
go mod tidy
```
- Build
```bash
 go build -o license-generator
 ```
- Commit your changes and push to your branch
```bash
git commit -m "made an awesomeFix"
git push origin fix/awesomeFix
```
- Open a pull request
