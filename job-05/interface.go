package main

import "learn.go/pkg/apis"

type Ranker interface {
	GetRank(fatRate float64) (rank int)
}

type Writer interface {
	Write(pi *apis.PersonalInformation)
}
