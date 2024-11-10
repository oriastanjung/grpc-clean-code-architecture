package utils

import "github.com/segmentio/ksuid"

func GenerateIDbyKSUID() ksuid.KSUID {
	return ksuid.New()
}