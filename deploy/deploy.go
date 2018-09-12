package deploy

import (
	"github.com/EdgeSmart/EdgeFairy/check"
	"github.com/EdgeSmart/EdgeFairy/user"
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
