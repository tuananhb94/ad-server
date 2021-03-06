package core

import (
	"github.com/valyala/fasthttp"
	"strconv"
    "encoding/base64"
	"fmt"
	"net/url"
)

func ClickHandler(ctx *fasthttp.RequestCtx) {
	args := ctx.QueryArgs()
	if !args.Has("i") {
		ctx.SetBody([]byte("{\"status\": 1}"))
		return
	}
	argsVlueBytes := args.Peek("i")
	queryStringBytes, err := base64.URLEncoding.DecodeString(string(argsVlueBytes))
	paramMap, err := url.ParseQuery(string(queryStringBytes))
	if err != nil {
		ctx.SetBody([]byte("{\"status\": 1," + "\"error\":" + err.Error() + "}"))
		return 
	}

	// search_id
	var searchId string
	if searchIds, exist := paramMap["search_id"]; exist {
		searchId = searchIds[0]
	}

	// slot_id
	var slotId uint32
	if slotIds, exist := paramMap["slot_id"]; exist {
		tmpInt, err := strconv.ParseUint(slotIds[0], 10, 32)
		if err != nil {
			ctx.SetBody([]byte("{\"status\": 1," + "\"error\":" + err.Error() + "}"))
		return 
	}
		slotId = uint32(tmpInt)
	}

	// ip
	var ip string
	if ips, exist := paramMap["ip"]; exist {
		ip = ips[0]
	}

	// device_id
	var deviceId string
	if deviceIds, exist := paramMap["device_id"]; exist {
		deviceId = deviceIds[0]
	}

	// os
	var os uint32
	if osString, exist := paramMap["os"]; exist {
		tmpInt, _ := strconv.ParseUint(osString[0], 10, 32)
		os = uint32(tmpInt)
	}

	// os_version
	var osVersion string
	if osVersions, exist := paramMap["os_version"]; exist {
		osVersion = osVersions[0]
	}

	// unit_id
	var unitId uint32
	if unitIdString, exist := paramMap["unit_id"]; exist {
		tmpInt, _ := strconv.ParseUint(unitIdString[0], 10, 32)
		unitId = uint32(tmpInt)
	}

	// creative_id
	var creativeId uint32
	if creativeIdString, exist := paramMap["creative_id"]; exist {
		tmp, _ := strconv.ParseUint(creativeIdString[0], 10, 32)
		creativeId = uint32(tmp)
	}

	// click_url
	var clickUrl string
	if clickUrls, exist := paramMap["click_url"]; exist {
		clickUrl = clickUrls[0]
	}
	AdServerLog.Debug(fmt.Sprintf("ClickHandler click_url=%s", clickUrl))

	ClickLog.Info(fmt.Sprintf(
		"click=1 searchId=%s slotId=%d ip=%s deviceId=%s os=%d osVersion=%s unit_id=%d creativeId=%d",
		searchId, slotId, ip, deviceId, os, osVersion, unitId, creativeId))
	ctx.Redirect(clickUrl, fasthttp.StatusFound)
}