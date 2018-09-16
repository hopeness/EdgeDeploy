package deploy

import (
	"github.com/EdgeSmart/EdgeDeploy/check"
	"github.com/EdgeSmart/EdgeDeploy/user"
)

// Run execute daemon
func Run() error {
	err := check.Run()
	if err != nil {
		return err
	}
	user.Login()
	return nil
}
