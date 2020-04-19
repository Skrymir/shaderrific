package main

import (
	"io/ioutil"
	"net/http"
)

const (
	frontOpenSceneId  = "36493"
	frontCloseSceneId = "41163"
	backOpenSceneId   = "20962"
	backCloseSceneId  = "33825"
	//frontFullCloseSceneId = "35073"
	//backFullCloseSceneId  = "3407"
)

//36493 front open
//41163 front close kid
//35073 front full close

//20962 back open
//33825 back close kid
//3407 back full close

//func allSadeData(hubIp string) string {
//	return shadeGet(hubIp, "/shades")
//}

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

func activateScene(hubIp string, sceneId string) string {
	return shadeGet(hubIp, "/scenes?sceneId="+sceneId)
}

func frontClose() {
	activateScene(Config.HubIp, frontCloseSceneId)
}

func frontOpen() {
	activateScene(Config.HubIp, frontOpenSceneId)
}

func backClose() {
	activateScene(Config.HubIp, backCloseSceneId)
}

func backOpen() {
	activateScene(Config.HubIp, backOpenSceneId)
}
