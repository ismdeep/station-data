package main

import (
	"errors"
	"os"

	"github.com/ismdeep/station-data/archiver"
	"github.com/ismdeep/station-data/pkg/blog"
)

var allInstances []blog.BloggerInterface

func init() {
	allInstances = []blog.BloggerInterface{
		&blog.ShellHacks{},
		&blog.BogomolovTech{},
		&blog.CoolShell{},
		&blog.Cybartist{},
		&blog.GoDev{},
		&blog.LeovanMe{},
		&blog.LogRocket{},
		&blog.Solidot{},
		&blog.WhyDegree{},
		&blog.MorningSpace{},
		&blog.DunWu{},
		&blog.DonaldFeury{},
		&blog.EltonMinetto{},
		&blog.LiamPage{},
		&blog.Akarin{},
		&blog.Qiwsir{},
		&blog.Izsk{},
		&blog.Zzbd{},
		&blog.Johnj{},
		&blog.RaphLinus{},
		&blog.BetterProgrammingPub{},
		&blog.Kaixin{},
		&blog.ArthurChiaoArt{},
		&blog.Imbajin{},
		&blog.MikelEvins{},
		&blog.MeituanTech{},
		&blog.B303248153{},
		&blog.IOMeter{},
		&blog.Antonz{},
		&blog.BouKe{},
		&blog.KeldosMe{},
		&blog.Evilsocket{},
	}
}

func Instances(sources []string) []blog.BloggerInterface {
	if len(sources) <= 0 {
		return allInstances
	}

	m := make(map[string]bool)
	for _, source := range sources {
		m[source] = true
	}

	var lst []blog.BloggerInterface
	for _, instance := range allInstances {
		if _, ok := m[instance.GetBloggerName()]; ok {
			lst = append(lst, instance)
		}
	}

	return lst
}

func main() {
	instances := Instances(os.Args[1:])
	var errLst []error
	for _, instance := range instances {
		if err := archiver.Archive(instance); err != nil {
			errLst = append(errLst, err)
		}
	}
	if err := errors.Join(errLst...); err != nil {
		panic(err)
	}
}
