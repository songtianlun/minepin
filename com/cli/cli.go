package cli

import (
	"fmt"
	"github.com/spf13/pflag"
)

type CLI struct {
	abbr   string
	dft    interface{}
	desc   string
	value  *bool
	handle HandleCLI
}

var MapCLI = make(map[string]*CLI)

type HandleCLI func()

func RegisterCLI(k string, abbr string, desc string, handle HandleCLI) {
	_, ok := MapCLI[k]
	if ok {
		panic(fmt.Sprintf("%s is already registered", k))
	}
	MapCLI[k] = &CLI{
		abbr:   abbr,
		dft:    false,
		desc:   desc,
		value:  pflag.BoolP(k, abbr, false, desc),
		handle: handle,
	}
}

func CheckCLI() (isCli bool) {
	pflag.Parse()
	for _, v := range MapCLI {
		if *v.value {
			v.handle()
			isCli = true
			break
		}
	}
	return
}
