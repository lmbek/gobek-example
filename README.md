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

## How to use
The best way to start using the project is to download the example project at: <br>
https://github.com/lmbek/gobek-example
<br><br>
Then install the Go and try to run the project with go run <br><br>
If it asks about github.com/lmbek/gobek then use go get github.com/lmbek/gobek to install it<br><br>
Then it should be ready and you can use go run on the project.

This example project uses this package and combines it with a local api. <br>
Then the Go api is being developed and customized by you together with the frontend (JavaScript, HTML, CSS) <br><br>

If you want to use this project without the builtin api and instead want to start from scratch, 
then please go to https://github.com/lmbek/gobek where you can find an example code of how to use the project from scratch

## How to test

	go test ./tests/...

## How to run
make your own main.go by following https://github.com/lmbek/gobek-example

## How to apply manifest and logo to executible
Use something like goversioninfo: https://github.com/josephspurrier/goversioninfo

## How to build

	go build -ldflags -H=windowsgui -o NewProjectName.exe

## For each project only one instance
The project is made, so you can only have one instance of the same organisation name and project name open, so if you have multiple project you are developing, please change their project names.

## For advanced users (Databases)
I have made an example project for more advanced users, where i demonstrated use of sqlite together with gobek <br>
The project can be found at: https://github.com/lmbek/gobek-sqlite-example