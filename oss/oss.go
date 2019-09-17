package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var DistList []string // list of dist

type (
	Plugin struct {
		Config Config // Git clone configuration
	}
	// OSS config
	Config struct {
		Dist            string // local package
		Path            string // oss path
		EndPoint        string // oss path
		AccessKeyID     string // access key id
		AccessKeySecret string // access key sectret
		ModName         string // git module name
	}
	// Envfile
	Envfile struct {
		ConfigPkg string   `yaml:"configPkg"`
		CheckList []string `yaml:"checkList"`
	}
)

func (p Plugin) Exec() error {
	if p.Config.ModName != "" {
		envfile := Envfile{}
		envfile.ReadYaml("./env.yaml")
		modname := envfile.CheckList
		var exist bool
		for _, mod := range modname {
			if mod == p.Config.ModName {
				exist = true
				break
			}
		}
		if exist {
			fmt.Printf("+ Name matching succeeded, 「%s」 continue !\n", p.Config.ModName)
			p.Upload()
		} else {
			fmt.Println("+ No matching name,jump step")
		}
	} else {
		fmt.Println("+ skip module package check")
		p.Upload()
	}
	return nil
}

func (p Plugin) Upload() {
	// 创建OSSClient实例。
	client, err := oss.New(p.Config.EndPoint, p.Config.AccessKeyID, p.Config.AccessKeySecret)
	if err != nil {
		HandleError(err)
	}

	bucketPath := strings.Split(p.Config.Path, "/")
	bucketName := bucketPath[0]
	objectName := bucketPath[1]

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	var maekerList []string

	marker := oss.Marker(objectName)
	for {
		lsRes, err := bucket.ListObjects(oss.MaxKeys(200), marker)
		if err != nil {
			HandleError(err)
		}
		marker = oss.Marker(lsRes.NextMarker)
		for _, path := range lsRes.Objects {
			obj := strings.Split(path.Key, "/")[0]
			if obj == objectName {
				maekerList = append(maekerList, path.Key)
			}
		}
		if !lsRes.IsTruncated {
			break
		}
	}

	fmt.Printf("+ %d files in total", len(maekerList))

	delRes, err := bucket.DeleteObjects(maekerList)
	if err != nil {
		HandleError(err)
	}
	fmt.Println("\n+ Deleted Objects:")
	for _, obj := range delRes.DeletedObjects {
		fmt.Println(obj)
	}
	fmt.Println("\n+ Clean Up!")

	listFile(p.Config.Dist)

	for _, file := range DistList {
		objectPath := objectName + "/" + file[len(p.Config.Dist)+1:]
		err = bucket.PutObjectFromFile(objectPath, file)
		if err != nil {
			HandleError(err)
		}
	}
}

func listFile(folder string) {
	files, _ := ioutil.ReadDir(folder) //specify the current dir
	for _, file := range files {
		if file.IsDir() {
			listFile(folder + "/" + file.Name())
		} else {
			fmt.Println(folder + "/" + file.Name())
			DistList = append(DistList, folder+"/"+file.Name())
		}
	}

}

func (c *Envfile) ReadYaml(f string) {
	buffer, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = yaml.Unmarshal(buffer, &c)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func HandleError(err error) {
	fmt.Println("+ Error:", err)
	os.Exit(-1)
}
