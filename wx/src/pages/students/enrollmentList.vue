<template>
  <view class="enrollment">
    <uni-nav-bar :fixed="true" :border="false" background-color="#F7F7F7" shadow status-bar left-icon="left"
      left-text="返回" title="编辑学生" @clickLeft="back" />
    <view class="content">
      <uni-table border stripe emptyText="暂无更多数据" @sort-change="sortChange()">
        <uni-tr>
          <uni-th width="100" align="center">科目</uni-th>
          <uni-th width="70" align="center">课时</uni-th>
          <uni-th width="100" align="center">类型</uni-th>
          <uni-th width="100" align="center">日期</uni-th>
        </uni-tr>

        <uni-tr v-for="(item, index) in list" :key="index">
          <uni-td align="center">
            <view>{{ item.courseName}}</view>
          </uni-td>
          <uni-td align="center">
            <view>{{ item.numSessions}}</view>
          </uni-td>
          <uni-td align="center">
            <view>{{ item.action == "subtract" ? "消耗" : "增加" }}</view>
          </uni-td>
          <uni-td align="center">
            <view>{{ formattedUseDate(item.useDate) }}</view>
          </uni-td>
        </uni-tr>
      </uni-table>
    </view>
  </view>
</template>

<script>
import { getEduClassSessionList } from '@/api/organization/enrollment'
import {formatDate} from '@/utils/formatDate'

export default {
  data () {
    return {
      userId: '',
      courseId: '',
      list:[]
    }
  },
  computed: {},
  methods: {
    formattedUseDate(t) {
      return formatDate(t);
    },

    init () {
      const params = {
        studentId: this.userId
      }
      getEduClassSessionList(params).then(res => {
        this.list = res.data.list
        console.log(this.list)
      })
    },
    back () {
      uni.navigateBack();
    },
  },
  watch: {},

  // 页面周期函数--监听页面加载
  onLoad (options) {
    this.userId = options.userId
    this.courseId = options.courseId
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
.enrollment {
  font-size: 15px;
  color: #606266;
}
</style>