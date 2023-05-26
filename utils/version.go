package utils

import (
	"fmt"
)

const (
	marjorVer string = "2"
	minorVer  string = "0"
	dateVer   string = "20230526"
)

var BackendVersion string = fmt.Sprintf("%s.%s-%s (backend)", marjorVer, minorVer, dateVer)
