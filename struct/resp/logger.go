package resp

import "fradyspace.com/go-server-practice/utils"

var logger *utils.Logger

func init() {
	logger = utils.NewLogger("make response")
}
