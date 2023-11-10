package about

import (
	"github.com/no-src/replacer/internal/version"
)

const logo = `
 ________  _______   ________  ___       ________  ________  _______   ________     
|\   __  \|\  ___ \ |\   __  \|\  \     |\   __  \|\   ____\|\  ___ \ |\   __  \    
\ \  \|\  \ \   __/|\ \  \|\  \ \  \    \ \  \|\  \ \  \___|\ \   __/|\ \  \|\  \   
 \ \   _  _\ \  \_|/_\ \   ____\ \  \    \ \   __  \ \  \    \ \  \_|/_\ \   _  _\  
  \ \  \\  \\ \  \_|\ \ \  \___|\ \  \____\ \  \ \  \ \  \____\ \  \_|\ \ \  \\  \| 
   \ \__\\ _\\ \_______\ \__\    \ \_______\ \__\ \__\ \_______\ \_______\ \__\\ _\ 
    \|__|\|__|\|_______|\|__|     \|_______|\|__|\|__|\|_______|\|_______|\|__|\|__|

`
const (
	openSourceUrl    = "https://github.com/no-src/replacer"
	documentationUrl = "https://pkg.go.dev/github.com/no-src/replacer@" + version.VERSION
	releaseUrl       = "https://github.com/no-src/replacer/releases"
	dockerImageUrl   = "https://hub.docker.com/r/nosrc/replacer"
)

// PrintAbout print the program logo and basic info
func PrintAbout(out func(format string, args ...any)) {
	out(logo)
	out("The replacer is a configuration-based file replace tool")
	out("Open source repository at: <%s>", openSourceUrl)
	out("Download the latest version at: <%s>", releaseUrl)
	out("The docker image repository address at: <%s>", dockerImageUrl)
	out("Full documentation at: <%s>", documentationUrl)
}
