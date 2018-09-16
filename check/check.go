package check

import (
	"fmt"
	"os/exec"
)

func Run() error {
	var err error
	// 检查是否有sudo权限
	err = checkSudo()
	if err != nil {
		return err
	}
	// 检查是否存在docker
	err = checkDocker()
	if err != nil {
		return err
	}

	return nil
}

func checkSudo() error {
	fmt.Print("Check sudo: ")
	cmd := exec.Command("sudo", "ls", "/tmp")
	_, err := cmd.Output()
	if err != nil {
		fmt.Print(err)
		return err
	}
	fmt.Println("ok")
	return nil
}

func checkDocker() error {
	fmt.Print("Check docker engine: ")
	cmd := exec.Command("sudo", "docker", "version")
	_, err := cmd.Output()
	if err != nil {
		fmt.Print(err)
		return err
	}
	fmt.Println("ok")
	return nil
}
