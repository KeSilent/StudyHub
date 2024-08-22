/*
 * @Author: Yang
 * @Date: 2023-04-11 15:38:34
 * @Description: 请填写简介
 */
/*
 * @Author: Yang
 * @Date: 2023-03-20 21:05:00
 * @Description: 微信返回
 */
package response

type WXLoginResp struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}
