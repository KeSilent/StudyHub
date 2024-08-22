<template>
  <div class="login">
    <view class="card-title">
      <view class="title">🐱学时印记🐱</view>
      <view class="content">当前微信没有绑定账户，请输入手机号来绑定账户。</view>
      <view class="phone">
        <uni-forms :modelValue="formData">
          <uni-forms-item label="手机号" name="phone">
            <uni-easyinput type="number" confirmType="done" v-model="phone" placeholder="请输入手机号" @confirm="bind()" />
          </uni-forms-item>
        </uni-forms>
      </view>
    </view>

    <view class="card-button">
      <button type="primary" :loading="loading" @click="bind()">登录</button>
    </view>
  </div>
</template>

<script>
import { wxLoginForPhone } from '@/api/system/login.js'

import Vue from 'vue';

export default {
  components: {},
  data () {
    return {
      loading: false,
      phone: ""
    }
  },
  computed: {},
  methods: {
    //绑定
    bind () {
      let that = this
      if (that.phone === "") {
        return
      }
      that.loading = true
      wx.login({
        success (res) {
          if (res.code) {
            wxLoginForPhone(res.code, that.phone).then((resp) => {
              uni.setStorageSync('token', resp.data.token)
              uni.setStorageSync('user', resp.data.user)
              if (resp.data.user.authorityId == Vue.prototype.administratorRoleId) {
                //管理员角色
                uni.switchTab({
                  url: '/pages/index/index'
                });
              } else {
                //学生家长角色
                uni.redirectTo({
                  url: '/pages/clientInfo/index'
                });
              }
            }).catch(err => {
              console.log(err)
            })
          } else {
            console.log('登录失败！' + res.errMsg)
          }
          that.loading = false
        }
      })
    }
  },
  watch: {},

  // 页面周期函数--监听页面加载
  onLoad () { },
  // 页面周期函数--监听页面初次渲染完成
  onReady () { },
  // 页面周期函数--监听页面显示(not-nvue)
  onShow () { },
  // 页面周期函数--监听页面隐藏
  onHide () { },
  // 页面周期函数--监听页面卸载
  onUnload () { },
  // 页面处理函数--监听用户下拉动作
  // onPullDownRefresh() { uni.stopPullDownRefresh(); },
  // 页面处理函数--监听用户上拉触底
  // onReachBottom() {},
  // 页面处理函数--监听页面滚动(not-nvue)
  // onPageScroll(event) {},
  // 页面处理函数--用户点击右上角分享
  // onShareAppMessage(options) {},
} 
</script>

<style scoped lang="scss">
.login {
  position: relative;

  .card-title {
    text-align: center;
    margin: 40px 15px 15px 15px;
    padding: 10px;
    border-radius: 15px;

    .title {
      font-size: 25px;
    }

    .content {
      font-size: 15px;
      margin-top: 10px;
    }

    .phone {
      margin-top: 30px;
    }
  }

  .card-button {
    position: fixed;
    bottom: 100px;
    left: 0;
    right: 0;
    width: 80%;
    border: 1px solid #EBEEF5;
    border-radius: 15px;
    margin: 0 auto;

    .text {
      font-size: 30px;
      margin-left: 10px;
    }
  }

}
</style>