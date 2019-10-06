
package main

import (
	"fmt"
	"net/http"
	"io"
	"os"
	"os/user"
	"os/exec"
)

func main() {
	// Check if root first
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	if u.Uid != "0" {
		panic("operation requires execution as root")
	}

	fmt.Println("downloading squash-firefox...");
	downloadToFilesystem("http://i.ronthecookie.me/squash-firefox", "firefox.squashfs")
	fmt.Println("mounting squashfs image...")
	err = os.MkdirAll("overlay/firefox", 744)
	if err != nil {
		panic(err)
	}
	err = exec.Command("mount", "").Run()
	if err != nil {
		panic(err)
	}
	
}


func downloadToFilesystem(url string, path string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	io.Copy(f, resp.Body)
}
