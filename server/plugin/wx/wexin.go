/*
 * @Author: Yang
 * @Date: 2023-04-11 15:41:59
 * @Description: 请填写简介
 */
/*
 * @Author: Yang
 * @Date: 2023-03-20 20:45:21
 * @Description: 微信接口
 */
package wx

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/KeSilent/study-hub/server/global"
	"github.com/KeSilent/study-hub/server/model/system/response"
)

/**
 * @description: 获取微信OpenId
 * @param {string} code
 * @return {*}
 */
func GetWXSession(code string) (*response.WXLoginResp, error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

	// 合成url, 这里的appId和secret是在微信公众平台上获取的
	url = fmt.Sprintf(url, global.WXAppId, global.WXSecret, code)

	// 创建http get请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 解析http请求中body 数据到我们定义的结构体中
	wxResp := response.WXLoginResp{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&wxResp); err != nil {
		return nil, err
	}

	// 判断微信接口返回的是否是一个异常情况
	if wxResp.ErrCode != 0 {
		errCodeStr := strconv.Itoa(wxResp.ErrCode)
		return nil, fmt.Errorf("ErrCode:%s  ErrMsg:%s", errCodeStr, wxResp.ErrMsg)
	}

	return &wxResp, nil
}
