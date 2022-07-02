package yaml

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"strings"
)

// ConfigYaml app config location
const (
	ConfigYaml = "configs/application.yaml"
)

// application config
var (
	configData []byte
)

func init() {
	fileByte, err := ioutil.ReadFile(ConfigYaml)
	if err != nil {
		log.Fatalln("load application config file error!")
	}
	configData = fileByte
}

// D crawler obj
type D struct {
	DouBan struct {
		Tag struct {
			Movie string `yaml:"movie" json:"movie"`
			Tv    string `yaml:"tv" json:"tv"`
		} `yaml:"tag" json:"tag"`

		Movie struct {
			Url string `yaml:"url" json:"url"`
		} `yaml:"movie" json:"movie"`

		Tv struct {
			Url string `yaml:"url" json:"url"`
		} `yaml:"tv" json:"tv"`
	} `yaml:"DouBan" json:"DouBan"`
}

func GetDouBan() *D {
	retVal := D{}
	d := retVal.DouBan
	_ = yaml.Unmarshal(configData, &retVal)
	if strings.EqualFold(d.Tag.Movie, "") {
		d.Tag.Movie = "https://movie.douban.com/j/search_tags?type=movie"
	}
	if "" == d.Tag.Tv {
		d.Tag.Tv = "https://movie.douban.com/j/search_tags?type=tv"
	}
	if strings.EqualFold(d.Movie.Url, "") {
		d.Movie.Url = "https://movie.douban.com/j/search_subjects?type=movie&tag=%s&sort=%s&page_limit=%s&page_start=%s"
	}
	if strings.EqualFold(d.Tv.Url, "") {
		d.Tv.Url = "https://movie.douban.com/j/search_subjects?type=tv&tag=%s&sort=%s&page_limit=%s&page_start=%s"
	}
	log.Println(d.Movie.Url)
	return &retVal
}

// S server config
type S struct {
	Server struct {
		Port int `yaml:"port" json:"port"`
	} `yaml:"server" json:"server"`
}

// R redis config
type R struct {
	Redis struct {
		Addr     string `yaml:"addr" json:"addr"`
		Host     string `yaml:"host" json:"host"`
		Port     string `yaml:"port" json:"port"`
		Password string `yaml:"password" json:"password"`
		Db       int    `yaml:"db" json:"db"`
	} `json:"redis"`
}

// M mongodb config
type M struct {
	Mongodb struct {
		Uri      string `yaml:"uri" json:"uri"`
		Database string `yaml:"database" json:"database"`
		Host     string `yaml:"host" json:"host"`
		Port     string `yaml:"port" json:"port"`
		Username string `yaml:"username" json:"username"`
		Password string `yaml:"password" json:"password"`
	} `json:"mongodb"`
}

func GetServer() *S {
	s := S{}
	_ = yaml.Unmarshal(configData, &s)
	if s.Server.Port <= 0 || s.Server.Port >= 65535 {
		s.Server.Port = 8080
	}
	return &s
}

func GetRedis() *R {
	r := R{}
	_ = yaml.Unmarshal(configData, &r)
	if strings.EqualFold(r.Redis.Host, "") {
		r.Redis.Host = "127.0.0.1"
	}
	if strings.EqualFold(r.Redis.Port, "") {
		r.Redis.Port = "6379"
	}
	return &r
}

func GetMongodb() *M {
	m := M{}
	_ = yaml.Unmarshal(configData, &m)
	if strings.EqualFold("", m.Mongodb.Uri) {
		m.Mongodb.Username = "admin"
		m.Mongodb.Password = "admin"
		m.Mongodb.Host = "127.0.0.1"
		m.Mongodb.Port = "27017"
	}
	return &m
}

func Test() {
	s := S{}
	err := yaml.Unmarshal(configData, &s)
	fmt.Println(s, s.Server.Port)
	if err != nil {
		fmt.Println("err msg", err)
	}
}
