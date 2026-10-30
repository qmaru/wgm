package utils

import (
	"fmt"
)

const (
	marjorVer string = "3"
	minorVer  string = "0"
	dateVer   string = "20261029"
)

var AppVersion string = fmt.Sprintf("%s.%s-%s", marjorVer, minorVer, dateVer)
