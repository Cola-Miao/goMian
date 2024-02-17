package initialize

import (
	"offerBook/dao"
	"offerBook/pkg"
)

func Init() (err error) {
	if err = pkg.Init(); err != nil {
		return
	}
	if err = dao.Init(); err != nil {
		return
	}
	return
}
