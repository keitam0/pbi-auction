package main

import "github.com/keitam0/pbi-auction/di"

func main() {
	dc := di.Container{}
	if err := dc.Router().Run(":80"); err != nil {
		panic(err)
	}
}
