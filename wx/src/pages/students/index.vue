<template>
  <view class="students">
    <!-- 如果这一层不是scroll-view标签要换成scroll-view或者添加一层scroll-view标签 -->
    <scroll-view style="height: 100%;" scroll-y="true">
      <view>

        <view class="bar"></view>
        <view class="search-bar">
          <uni-search-bar class="search" v-model="searchName" radius="100" placeholder="学生姓名" clearButton="none"
            cancelButton="none" @confirm="initUserList()" />
        </view>
        <view class="content">
          <uni-table border stripe emptyText="暂无更多数据" @sort-change="sortChange()">
            <uni-tr>
              <uni-th align="center">姓名</uni-th>
              <uni-th align="center">科目</uni-th>
              <uni-th align="center">课时</uni-th>
              <!-- <uni-th width="100" align="center">电话</uni-th> -->
            </uni-tr>
            <uni-tr v-for="(item, index) in userList" :key="index">
              <uni-td align="center">
                <view @click="showEditStudent(item)">{{ item.nickName }}</view>
              </uni-td>
              <uni-td align="center">
                <view @click="showCourseUpdateDialog(item)">{{ item.eduEnrollment[0].eduCourse.courseName }}</view>
              </uni-td>
              <uni-td align="center">
                <view @click="showHoursConsumed(item)">
                  <span class="hours">{{ item.eduEnrollment[0].remainingSessions }}</span>
                </view>
              </uni-td>
              <uni-td align="center">
                <view @click="showEditStudent(item)">{{ item.phone }}</view>
              </uni-td>
            </uni-tr>
          </uni-table>
        </view>
        <uni-fab horizontal="right" :pattern="pattern" popMenu="false" @fabClick="toAddStudent"></uni-fab>

        <view>
          <!-- 消耗课时 -->
          <uni-popup ref="hoursConsumed" type="dialog">
            <view class="dialog">
              <view class="uni-dialog-title">
                <text>消耗</text>
                <text style="color: #e43d33;">{{ selectStudent.nickName }}</text>
                <text>课程</text>
              </view>
              <view class="uni-dialog-content">
                <view style="display: contents;">
                  <view style="margin: 20px;font-size: 26px;" @click="subtraction()">-</view>
                  <uni-easyinput style="width: 30%;" :clearable="false" v-model="classHoure" type="text"
                    placeholder="" />
                  <view style="margin: 20px;font-size: 26px;" @click="addition()">+</view>
                </view>
              </view>
              <view style="padding:0 20px 20px 20px;display: flex;align-items: center;">
                <uni-datetime-picker type="date" :clear-icon="false" v-model="useDate" />
              </view>
              <view class="uni-dialog-button-group">
                <view type="primary" class="uni-dialog-button" @click="this.$refs.hoursConsumed.close()">取消</view>
                <view type="primary" class="uni-dialog-button uni-border-left" @click="hoursConsumed()">确定</view>
              </view>
            </view>

          </uni-popup>

          <uni-popup ref="courseUpdateDialog" type="dialog">
            <view class="dialog">
              <text class="uni-dialog-title">切换课程</text>
              <view class="uni-dialog-content" style="display:block;">
                <uni-data-select v-model="selectStudent.eduEnrollment[0].courseId" :clear="false"
                  :localdata="courseList"></uni-data-select>
              </view>
              <view class="uni-dialog-button-group">
                <view type="primary" class="uni-dialog-button" @click="this.$refs.courseUpdateDialog.close()">取消</view>
                <view type="primary" class="uni-dialog-button uni-border-left" @click="updateCourse()">确定</view>
              </view>
            </view>
          </uni-popup>


          <!-- 提示信息弹窗 -->
          <uni-popup ref="message" type="message">
            <uni-popup-message :type="msgType" :message="messageText" :duration="2000"></uni-popup-message>
          </uni-popup>
        </view>
      </view>
    </scroll-view>
  </view>
</template>

<script>
import { getUserListByRoleIdAndOrgId } from '@/api/user/user.js'
import { getEduCourseList } from '@/api/organization/course'
import {formatDate} from '@/utils/formatDate'

import { updateEduEnrollment, consumptionClass } from '@/api/organization/enrollment'

export default {
  components: {},
  data () {
    return {
      page: 1,
      pageSize: 1000,
      messageText: '',
      msgType: 'success',
      pattern: {
        color: '#7A7E83',
        backgroundColor: '#fff',
        selectedColor: '#18bc37',
        buttonColor: '#18bc37',
        iconColor: '#fff'
      },
      searchName: '',
      userList: [],
      classHoure: 1,
      useDate:formatDate(Date.now()),
      courseList: [],
      selectStudent: {},
      rules: {
        nickName: {
          rules: [{
            required: true,
            errorMessage: '请输入姓名',
          }
          ]
        }
      }
    }
  },
  created () {
    this.init()
  },
  methods: {
    init () {
      this.initUserList()
      this.initEduCourseList()
    },
    //切换课程
    updateCourse () {
      this.selectStudent.eduEnrollment[0]
      const param = {
        courseId: this.selectStudent.eduEnrollment[0].courseId,
        ID: this.selectStudent.eduEnrollment[0].ID
      }
      updateEduEnrollment(param).then(res => {
        this.$refs.courseUpdateDialog.close()
        this.messageText = '修改成功'
        this.$refs.message.open()
        this.initUserList()
      })
    },
    //消费课时
    async hoursConsumed () {
      const param = {
        userId: this.selectStudent.ID,
        courseId: this.selectStudent.eduEnrollment[0].courseId,
        sessionsToConsume: this.classHoure,
        useDate:this.useDate
      }
      consumptionClass(param).then(res => {
        if(res.code==0){
          this.$refs.hoursConsumed.close()
          this.messageText = '消费成功'
          this.$refs.message.open()
          this.initUserList()
        }else{
          this.msgType = 'error'
          this.messageText = res.msg
          this.$refs.message.open()
        }
      })
    },

    //显示课程修改弹窗
    showCourseUpdateDialog (student) {
      this.selectStudent = student
      this.$refs.courseUpdateDialog.open()
    },
    //显示消耗课时弹窗
    showHoursConsumed (student) {
      this.selectStudent = student
      this.$refs.hoursConsumed.open()
    },
    //显示学生修改弹窗
    showEditStudent (student) {
      uni.navigateTo({
        url: '/pages/students/info?userId=' + student.ID
      });
    },
    addition () {
      if (this.classHoure <= 5) {
        this.classHoure++;
      }
    },
    subtraction () {
      if (this.classHoure > 1) {
        this.classHoure--;
      }
    },
    //排序
    sortChange () {
      userList.sort(function (a, b) {
        return Date.parse(b.eduEnrollment[0].eduCourse.courseName) - Date.parse(a.eduEnrollment[0].eduCourse.courseName)
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
        userName: this.searchName
      }
      getUserListByRoleIdAndOrgId(params).then(res => {
        this.userList = res.data.list
      })
    },
    initEduCourseList () {
      var that = this
      getEduCourseList().then(res => {
        if (res.code === 0) {
          const courseList = []
          var list = res.data.list
          list.forEach(item => {
            courseList.push({
              text: item.courseName,
              value: item.ID
            })
          })
          that.courseList = courseList
        }
      })
    },
    //添加学生
    toAddStudent () {
      uni.navigateTo({
        url: '/pages/students/addStudent'
      });
    },
  },
  watch: {},

  // 页面周期函数--监听页面加载
  onLoad () { },
  // 页面周期函数--监听页面初次渲染完成
  onReady () { },
  // 页面周期函数--监听页面显示(not-nvue)
  onShow () {
    var pages = getCurrentPages();	// 获取当前页面栈
    var curPage = pages[pages.length - 1];	// 当前页面
    if (curPage.data.init) {
      curPage.data.init = false;
      this.init()    // A页面 method中的方法，用来刷新页面A
    }
  },
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
.students {
  .bar {
    height: var(--status-bar-height);
  }

  .search-bar {
    display: flex;
    align-items: center;
    padding-top: var(--status-bar-height);
    width: 75%;
    margin: 0 14rpx;

    .search {
      width: 100%;
    }
  }

  .content {
    .hours {
      color: #f29e99;
    }

    .uni-button {
      margin-right: 10px;
    }
  }

  .dialog {
    width: 300px;
    border-radius: 11px;
    background-color: #fff;

    .uni-dialog-title {
      display: flex;
      flex-direction: row;
      justify-content: center;
      padding-top: 25px;
    }

    .uni-dialog-content {
      display: flex;
      flex-direction: row;
      justify-content: center;
      align-items: center;
      padding: 20px;
    }

    .uni-dialog-button-group {
      display: flex;
      flex-direction: row;
      border-top-color: #f5f5f5;
      border-top-style: solid;
      border-top-width: 1px;

      .uni-dialog-button {
        display: flex;
        flex: 1;
        flex-direction: row;
        justify-content: center;
        align-items: center;
        height: 45px;
      }

      .uni-border-right {
        border-bottom-left-radius: 11px;
      }

      .uni-border-left {
        border-left-color: #f0f0f0;
        border-left-style: solid;
        border-left-width: 1px;
        background-color: #18bc37;
        color: #fff;
        border-bottom-right-radius: 11px;
      }
    }
  }
}
</style>