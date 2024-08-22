<!--
 * @Author: Yang
 * @Date: 2019-08-22 19:41:20
 * @Description: 学生家长端
-->
<template>
  <view class="content">
    <scroll-view style="height: 300px;" scroll-y="true" refresher-enabled="true" :refresher-triggered="triggered"
      :refresher-threshold="100" refresher-background="lightgreen" @refresherrefresh="onRefresh"
      @refresherpulling="onPulling" @refresherrestore="onRestore" @refresherabort="onAbort">
      <view class="card-title">
        <view class="title">🐱学时印记🐱</view>
      </view>
      <view class="text-area">
        <uni-table border stripe emptyText="暂无更多数据" @sort-change="sortChange()">
          <uni-tr>
            <uni-th align="center">学生姓名</uni-th>
            <uni-th align="center">剩余课时</uni-th>
          </uni-tr>
          <uni-tr v-for="(item, index) in list" :key="index">
            <uni-td align="center">
              <view>{{ item.nickName }}</view>
            </uni-td>
            <uni-td align="center">
              <view>{{ item.eduEnrollment[0].remainingSessions }}</view>
            </uni-td>
          </uni-tr>
        </uni-table>
      </view>
    </scroll-view>
    <view class="bottom">
      <button type="primary" @click="out()">退出</button>
    </view>
  </view>
</template>

<script>
import { getUserListByRoleIdAndOrgId } from '@/api/user/user.js'
export default {
  components: {},
  data () {
    return {
      list: [],
      page: 1,
      pageSize: 1000,
      triggered: false
    }
  },
  computed: {},
  methods: {
    out () {
      uni.clearStorageSync();
      uni.reLaunch({
        url: '/pages/login/index'
      });
    },
    //获取学生列表
    initUserList () {
      var user = uni.getStorageSync('user')
      let params = {
        page: this.page,
        pageSize: this.pageSize,
        roleId: this.studentRoleId,
        orgId: user.eduOrganizationID,
        phone: user.phone
      }
      getUserListByRoleIdAndOrgId(params).then(res => {
        this.list = res.data.list
      })
    },
    // 自定义下拉刷新控件被下拉
    onPulling (e) {
      console.log("onpulling", e);
      if (e.detail.deltaY < 0) return  // 防止上滑页面也触发下拉
      this.triggered = true;
    },
    onRefresh () {
      if (this._freshing) return;
      this._freshing = true;
      setTimeout(() => {
        this.triggered = false;
        this._freshing = false;
        this.initUserList();
      }, 500);
    },
    onRestore () {
      this.triggered = 'restore'; // 需要重置
      console.log("onRestore");
    },
    onAbort () {
      console.log("onAbort");
    }
  },
  watch: {},

  // 页面周期函数--监听页面加载
  onLoad () {
    this._freshing = false;
    this.initUserList()
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
.content {
  position: relative;
  margin-top: 100px;

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
  }
  .bottom{
    position: fixed;
    margin: 60rpx;
    bottom: 80rpx;
    left: 0;
    right: 0;
  }
}
</style>