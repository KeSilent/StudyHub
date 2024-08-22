<template>
  <view class="userInfor">
    <view class="user">
      <image :src="userData.headerImg" mode="aspectFill" />
      <view>
        <view class="name">{{ userData.nickName }}</view>
        <view class="organization">{{ userData.eduOrganization.name }}</view>
      </view>
    </view>
    <uni-list>
      <uni-list-item title="课程管理" showArrow></uni-list-item>
    </uni-list>

    <view class="button">
      <button type="primary" @click="out()">退出</button>
    </view>
  </view>
</template>

<script>
export default {
  data () {
    return {
      userData: {}
    }
  },
  methods: {
    init () {
      this.userData = uni.getStorageSync('user')
      console.log(this.userData);
    },
    out () {
      uni.clearStorageSync();
      uni.reLaunch({
        url: '/pages/login/index'
      });
    },
  },
  watch: {},

  // 页面周期函数--监听页面加载
  onLoad () {
    this.init()
  },
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
.userInfor {
  .user {
    display: flex;
    align-items: center;
    margin-bottom: 20upx;
    padding: 20px;
    font-size: 14px;

    image {
      width: 100upx;
      height: 100upx;
      border-radius: 50%;
      margin-right: 20px;
    }
  }

  .button {
    margin-top: 20px;
    padding: 20px;

    button {
      width: 100%;
    }
  }
}
</style>