<template>
  <view class="init">
    <view class="card-title">
      <view class="title">🐱学时印记🐱</view>
      <view class="content">欢迎使用🐱学时印记🐱，请选择您的身份。</view>
    </view>
    <view class="card-button company" @click="toCompany()">
      <uni-icons type="shop" color="#fff" size="80"></uni-icons>
      <text class="text">我是机构</text>
    </view>
    <view class="card-button individual" @click="">
      <uni-icons type="person" color="#fff" size="80"></uni-icons>
      <text class="text">我是学生</text>
    </view>
  </view>
</template>

<script>
export default {
  components: {},
  data () {
    return {}
  },
  computed: {},
  methods: {
    toCompany () {
      uni.navigateTo({
        url: '/pages/company/index'
      });
    },
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
.init {
  display: flex;
  flex-direction: column;
  justify-content: center;

  .card-title {
    text-align: center;
    border: 1px solid #EBEEF5;
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
  }

  .card-button {
    border: 1px solid #EBEEF5;
    border-radius: 15px;
    margin: 60rpx 100rpx 20rpx 100rpx;
    padding: 10px;
    display: flex;
    justify-content: center;

    .text {
      font-size: 30px;
      margin-top: 30px;
      margin-left: 10px;
    }
  }

  .company {
    background-color: #3D7EFF;
    color: #fff;
  }

  .individual {
    background-color: #FFB7A1;
    color: #fff;
  }
}
</style>