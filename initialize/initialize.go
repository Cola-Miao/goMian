package initialize

import (
	"goMian/dao"
	"goMian/generator"
	"goMian/pkg"
)

func Init() (err error) {
	if err = pkg.Init(); err != nil {
		return
	}
	if err = dao.Init(); err != nil {
		return
	}
	if err = generator.Gtr.Init(); err != nil {
		return err
	}
	return
}
