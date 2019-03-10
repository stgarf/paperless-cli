package paperless

import (
	"fmt"
)

func (p Paperless) String() string {
	if !p.UseHTTPS {
		return fmt.Sprintf("http://%v:%v%v", p.Hostname, p.Port, p.Root)
	}
	return fmt.Sprintf("https://%v:%v%v", p.Hostname, p.Port, p.Root)
}

func (p Paperless) showInstanceInformation() string {
	return fmt.Sprintf("Username: %v, Hostname: %v, Port: %v, API root: %v, HTTPS: %v", p.Username, p.Hostname, p.Port, p.Root, p.UseHTTPS)
}

// ShowInstanceInformation shows the currently loaded Paperless instance configuration
func (p Paperless) ShowInstanceInformation() {
	fmt.Println(p.showInstanceInformation())
}
