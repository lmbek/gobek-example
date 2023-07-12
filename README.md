# gobek (Local GUI / Go Chrome Framework)

gobek is a simple framework made for developing localhosted software that can reuse chrome/chromium or embed chromium (in future releases). Both available in deployment for the applications.

This framework uses Chrome (Windows) or Chromium (Linux) as frontend by opening them with cmd/terminal and hosting a localhost webserver, while opening chrome/chromium with --app and --user-directory arguments. The frontend can be changed by the user in runtime, while the backend needs to be compiled/build. The API can be decoupled in future versions, so every part of the application is changeable - Sustainable development. Frontends is easy to change. Alternatives to this is embedding a chromium or webview framework into the project, which will require more space. I chose to depend on Chrome/Chromium, as they are market leaders and html/css/javascript technology frontrunners.

Feel free to use this piece of software, I will be happy to assist you <br>
I am currently working on this project, it will be updated and maintained. I consider it production ready. <br>
<br>
This project is used by Beksoft ApS for projects such as:
* BekCMS
* PingPong Game made in Three.js
* Several local webbased software projects
  <br>
Write to me at lars@beksoft.dk if you want to have your project listed<br>

## Contributors
Lars Morten Bek (https://github.com/lmbek) <br>
Ida Marcher Jensen (gobek-example) (https://github.com/notHooman996) <br>
<br><br>
Feel free to open an issue or pull request with feature requests or if you have any issues<br>

## Requirements to developers
Go 1.20+ <br>
Chrome (Windows) or Chromium (Linux) <br>
macOS/DARWIN NOT SUPPORTED YET <br>

## Requirements for users
Chrome (Windows) or Chromium (Linux) <br>
macOS/DARWIN NOT SUPPORTED YET <br>

## How to use (download example project)
The best way to start using the project is to download the example project at: <br>
https://github.com/lmbek/gobek-example

This example project uses this package and combines it with a local api. <br>
Then the Go api is being developed and customized by you together with the frontend (JavaScript, HTML, CSS) <br>

## How to use (with go get)
first run the following in CMD (with go installed) <br><br>
The Go package:

    go get github.com/lmbek/gobek
<br><br>
Example: how to add framework to main.go<br>

 	//go:generate ./bin/windows/go-packager/GoPackager.exe
	// go generate
	// go build -ldflags -H=windowsgui -o NewProjectName.exe
	
	// ** Generating and Building **
	// The above is the generate commands to add icon and manifest to binary,
	// and the build command
	// NOTE: GoPackager.exe is a modified version of goversioninfo (https://github.com/josephspurrier/goversioninfo).
	// You can compile your own GoPackager.exe at ./bin/windows/go-packager by go building the main.go file
	
	// ** Reason Of Framework Explained **
	// This application is meant to find a port and launch a frontend in chrome (windows) or chromium (linux).
	// Meanwhile, it will open a http localhost backend with an api.
	// The API package at ./backend/api/ can be replaced to fit another application,
	// if the framework is to be used for other applications.
	// It is important to note, that the system will require a frontend directory and a data directory,
	// when the application is tested and released, these will also need to be managed (other applications can be put in)
	// This framework is very open-source, since when released to users, the users can modify the frontend files.
	// However, the backend can be kept as binary, but sent to users if they later want to modify it,
	// for example users must be able to modify the application backend, and replace the binary with the modified version.
	
	package main
	
	import (
		"fmt"
		"github.com/lmbek/gobek"
		"os"
	)
	
	// For windows, we need an organisation name and project name
	var organisationName = "NewOrganisationName" // put in organisation name
	var projectName = "NewProjectName"           // put in project name
	var chromeLauncher = gobek.ChromeLauncher{
		Location:                os.Getenv("programfiles") + "\\Google\\Chrome\\Application\\chrome.exe",
		FrontendInstallLocation: os.Getenv("localappdata") + "\\Google\\Chrome\\InstalledApps\\" + organisationName + "\\" + projectName,
	}
	var chromiumLauncher = gobek.ChromiumLauncher{
		Location: "/var/lib/snapd/desktop/applications/chromium_chromium.desktop",
	}
	var frontendPath = "./frontend" // this should be set to where frontend files is (frontend folder: html, css, javascript...)
	func main() {
		err := gobek.StartDefault(frontendPath, chromeLauncher, chromiumLauncher)
		if err != nil {
			fmt.Println(err)
		}
	}

## How to test

	go test ./tests/...

## How to run
make your own main.go by following https://github.com/lmbek/gobek-example

## How to apply manifest and logo to executible
Use something like goversioninfo: https://github.com/josephspurrier/goversioninfo

## How to build

	go build -ldflags -H=windowsgui -o NewProjectName.exe
