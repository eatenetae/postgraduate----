package core

import (
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"server/global"
)

func EsConnect() *elastic.Client {
	var err error
	sniffOpt := elastic.SetSniff(false)
	host := global.Config.ES.URL()
	c, err := elastic.NewClient(
		elastic.SetURL(host),
		sniffOpt,
		elastic.SetBasicAuth(global.Config.ES.User, global.Config.ES.Password),
	)
	if err != nil {
		logrus.Fatalf("es连接失败 %s", err.Error())
	}
	return c
}
