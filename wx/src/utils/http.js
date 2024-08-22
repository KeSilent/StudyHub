import {
  config
} from '../config.js'
const baseUrl = config.base_url

export const apiResquest = (options) => {
  const token = uni.getStorageSync('token'); // 获取token，可以根据实际情况修改
  return new Promise((resolve, reject) => {
    uni.request({
      url: baseUrl + options.url,
      method: options.method,
      data: options.data,
      header: {
        'content-type': 'application/json',
        'x-token': token // 添加token到请求头
      },
      success: (res) => {
        if (res.code != 0) {
          resolve(res.data);
        } else {
          reject(res);
        }
      },
      fail: (err) => {
        reject(err);
      }
    })
  }).catch((err) => {
    uni.showToast({ // 对错误进行提示
      title: err.errMsg || '请求失败',
      icon: 'none'
    });
    throw err; // 继续抛出错误，以便在调用时处理
  });
};