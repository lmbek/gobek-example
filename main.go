//go:generate ./generate/GoPackager.exe
// go generate
// go build -ldflags -H=windowsgui -o NewProjectName.exe

// ** Generating and Building **
// The above is the generate commands to add icon and manifest to binary,
// and the build command
// NOTE: GoPackager.exe is a modified version of goversioninfo.
// You can find goversioninfo at https://github.com/josephspurrier/goversioninfo

// ** Reason Of Framework Explained **
// This application is meant to find a port and launch a frontend in chrome (windows) or chromium (linux).
// Meanwhile, it will open a http localhost backend with an api.
// The API package at ./backend/api/ can be replaced to fit another application,
// if the framework is to be used for other applications.
// It is important to note, that the system will require a frontend directory and a data directory,
// when the application is tested and released, these will also need to be managed (other applications can be put in)
// This framework is very open-source, since when released to users, the users can modify the frontend files.
// However the backend can be kept as binary, but sent to users if they later want to modify it,
// for example users must be able to modify the application backend, and replace the binary with the modified version.

package main

import (
	"fmt"
	"github.com/lmbek/gobek/launcher"
	"os"
)

// For windows we need a organisation name and project name
var organisationName = "NewOrganisationName" // put in organisation name
var projectName = "NewProjectName"           // put in project name

var frontendPath = "./frontend" // this should be set to where frontend files is (frontend folder: html, css, javascript...)

var chromeLauncher = launcher.ChromeLauncher{
	Location:                os.Getenv("programfiles") + "\\Google\\Chrome\\Application\\chrome.exe",
	FrontendInstallLocation: os.Getenv("localappdata") + "\\Google\\Chrome\\InstalledApps\\" + "DefaultOrganisationName" + "\\" + "DefaultProjectName",
}

var chromiumLauncher = launcher.DefaultChromiumLauncher // default chrome or chromium launcher settings can be used like this

/*
	// Otherwise they can also be customized like this

	var chromiumLauncher = launcher.ChromiumLauncher{
		Location:      "/var/lib/snapd/desktop/applications/chromium_chromium.desktop", // TODO: check if better location or can be customised
		Domain:        "localhost",
	}
*/

func main() {
	// api example is temporary not avaiable before release of gobek 0.7.0

	//api.Init()
	/*
		var once sync.Once
		once.Do(func() {
			http.HandleFunc("/", fileserver.ServeFileServer)
			http.HandleFunc("/api/", api.ServeAPIUseGZip)
		})
		err := launcher.Start(frontendPath, chromeLauncher, chromiumLauncher) // serves "/" as fileserver.ServeFileServer. If you want to manage "/", then use launcher.StartCustom() instead
		if err != nil {
			fmt.Println(err)
		}
	*/
	err := launcher.StartDefault(frontendPath, chromeLauncher, chromiumLauncher)
	if err != nil {
		fmt.Println(err)
	}
}
