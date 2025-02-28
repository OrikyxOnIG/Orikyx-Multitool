package banner

import "fmt"

func PrintBanner() {
	banner := `
    █▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀
    █  O R I K Y X  M U L T I T O O L       █
    █   A powerful multitool for networking █
    █   SSH, and much more...               █
    █▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀ 
	`
	fmt.Println(banner)
}
