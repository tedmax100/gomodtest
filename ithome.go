package ithome

import (
	"github.com/tedmax100/ithome/ironman"

	log "github.com/sirupsen/logrus"
)

func PrintItHome() {
	log.Info("hi ItHome V0.0.7")
	ironman.PrintIronMan()
}

func PrintItHomeV2() {
	log.Info("hi ItHome V2.0")
	ironman.PrintIronMan()
}
