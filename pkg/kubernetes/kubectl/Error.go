package kubectl

import "errors"

var (
	BinaryPathCanNotBeFoundOnKubectl = errors.New("binary path can not be found")
	OptionCanNotBeFoundOnKubectl     = errors.New("option can not be found")
)
