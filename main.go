package main

import (
	"github.com/daycat/aldershark/templates"
	"github.com/daycat/aldershark/utils"
	"github.com/k0kubun/pp/v3"
	"sync"
)

var (
	Sfc templates.CFG
	wg  sync.WaitGroup
)

func main() {
	cfg := utils.ReadCFG()
	Sfc = utils.GetSurfsharkConfig(cfg.URL)
	for i := range Sfc {
		wg.Add(1)
		go utils.Resolve(Sfc[i], cfg.SecureDNS, wg)
	}
	wg.Wait()

	pp.Print(Sfc)
}
