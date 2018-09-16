package install

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

const (
	scriptURL  = "https://get.docker.com/"
	scriptPath = "/tmp/docker_install_%s.sh"
)

func Run() error {
	// download()
	cmd := exec.Command("wget", "-qO- https://get.docker.com/", " | sh")
	output, err := cmd.Output()
	fmt.Println(string(output))
	if err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}

func download() (string, error) {
	res, err := http.Get(scriptURL)
	if err != nil {
		panic(err)
	}

	path := fmt.Sprintf(scriptPath, "a")
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(f, res.Body)
	return path, err
}
