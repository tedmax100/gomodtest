package ithome

import (
	"github.com/tedmax100/ithome/ironman"

	log "github.com/sirupsen/logrus"
)

func PrintItHome() {
	log.Info("hi ItHome")
	ironman.PrintIronMan()
}

func V2PrintItHome() {
	log.Debug("hi ItHome V2")
	ironman.PrintIronMan()
}
