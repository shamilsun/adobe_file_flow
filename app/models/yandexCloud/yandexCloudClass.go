package yandexCloud

import (
	"context"
	"github.com/shamilsun/yandex-disk-sdk-go"
	"log"
)

func getInstance() yadisk.YaDisk {

	yaDisk,err := yadisk.NewYaDisk(context.Background(),nil, &yadisk.Token{AccessToken: YandexCloudTokenSetting.GetValue()})
	if err != nil {
		panic(err.Error())
	}
	disk,err := yaDisk.GetDisk([]string{})
	if err != nil {
		// If response get error
		e, ok := err.(*yadisk.Error)
		if !ok {
			log.Println(e)
			panic(err.Error())
		}
		// e.ErrorID
		// e.Message
	}
	log.Println("connected to yadisk ", disk)
	return yaDisk
}

func ShareAndGetPublicURL(pathUnderCloudDir string) string {

	yaDisk := getInstance()

	link, err := yaDisk.CreateResource(pathUnderCloudDir, nil)
	if err != nil {
	  log.Println("err: ", err)
	}
	log.Println("link: ", link, link.Href)

	resourse, err := yaDisk.PublishResource(pathUnderCloudDir, nil)
	if err != nil {
		log.Println("err: ", err)
	}
	log.Println("link: ",resourse, resourse.Href)

	resourse2, err := yaDisk.GetResource(pathUnderCloudDir, nil,20,0,false ,"S","name")
	if err != nil {
		log.Println("err: ", err)
	}
	log.Println("resource: ", resourse2.PublicURL)

	return resourse2.PublicURL
}