package utils

import (
	"encoding/json"
	"github.com/daycat/aldershark/templates"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"net/http"
)

func ReadCFG() templates.CONFIG {
	var yamlcfg templates.CONFIG
	yamlFile, _ := ioutil.ReadFile("config.yaml")
	yaml.Unmarshal(yamlFile, &yamlcfg)
	return yamlcfg
}

func GetSurfsharkConfig(link string) templates.CFG {
	resp, err := http.Get("https://api.surfshark.com/v4/server/clusters/generic")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var (
		c templates.CFG
	)
	config, err := io.ReadAll(resp.Body)
	json.Unmarshal([]byte(config), &c)
	return c
}
