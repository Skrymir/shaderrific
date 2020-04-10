package main

import (
	"io/ioutil"
	"net/http"
)

const frontOpenSceneId = 1
const frontCloseSceneId = 2
const backOpenSceneId = 3
const backCloseSceneId = 4

func allSadeData(hubIp string) string {
	return shadeGet(hubIp, "/shades")
}

func allSceneData(hubIp string) string {
	return shadeGet(hubIp, "/scenes")
}

func shadeGet(hubIp string, u string) string {
	resp, _ := http.Get(baseUrl(hubIp) + u)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func baseUrl(hubIp string) string {
	return "http://" + hubIp + "/api"
}
