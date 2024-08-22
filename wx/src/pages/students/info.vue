<template>
  <view class="info">
    <view class="content">
      <view class="item">
        <view class="title">姓名：</view>
        <view>{{ formData.nickName }}</view>
      </view>
      <view class="item">
        <view class="title">科目：</view>
        <view>{{ formData.eduEnrollment[0].eduCourse.courseName }}</view>
      </view>
      <!-- <view class="item">
        <view class="title">总课时：</view>
        <view>{{ formData.eduEnrollment[0].totalSessions }}</view>
      </view> -->
      <view class="item">
        <view class="title">剩余课时：</view>
        <view>{{ formData.eduEnrollment[0].remainingSessions }}</view>
      </view>
      <view>
        <view class="info-button">
          <button type="default" style="background-color: cornflowerblue;color: #fff;"
            @click="this.$refs.studentDialog.open()">修改姓名</button>
        </view>
        <view class="info-button">
          <button type="warn" @click="del">删除学生</button>
        </view>
        <view class="info-button">
          <button type="default" @click="goLog()">查看课时记录</button>
        </view>
        <view class="info-button">
          <button type="primary" @click="this.$refs.addSessionDialog.open()">增加课时</button>
        </view>
      </view>
    </view>
    <view>
      <uni-popup ref="studentDialog" type="dialog">
        <view class="dialog">
          <text class="uni-dialog-title">修改学生</text>
          <view class="uni-dialog-content" style="display:block;">
            <uni-forms label-width="80px" ref="form" :modelValue="formData" :rules="rules">
              <uni-forms-item required label="姓名" name="nickName">
                <uni-easyinput type="text" v-model="formData.nickName" placeholder="请输入姓名" />
              </uni-forms-item>
            </uni-forms>
          </view>
          <view class="uni-dialog-button-group">
            <view type="primary" class="uni-dialog-button" @click="this.$refs.studentDialog.close()">取消</view>
            <view type="primary" class="uni-dialog-button uni-border-left" @click="updateUserInfo()">确定</view>
          </view>
        </view>
      </uni-popup>

      <uni-popup ref="addSessionDialog" type="dialog">
        <view class="dialog">
          <view class="uni-dialog-title">
            <text>增加</text>
            <text style="color: #e43d33;">{{ selectStudent.nickName }}</text>
            <text>课时</text>
          </view>
          <view class="uni-dialog-content">
            <view style="display: contents;">
              <view style="margin: 20px;font-size: 26px;" @click="subtraction()">-</view>
              <uni-easyinput style="width: 30%;" :clearable="false" v-model="classHoure" type="number" placeholder="" />
              <view style="margin: 20px;font-size: 26px;" @click="addition()">+</view>
            </view>
          </view>
          <view style="padding:0 20px 20px 20px;display: flex;align-items: center;">
            <uni-datetime-picker type="date" :clear-icon="false" v-model="useDate" />
          </view>
          <view class="uni-dialog-button-group">
            <view type="primary" class="uni-dialog-button" @click="this.$refs.addSessionDialog.close()">取消</view>
            <view type="primary" class="uni-dialog-button uni-border-left" @click="addStudentSession()">确定</view>
          </view>
        </view>

      </uni-popup>


      <!-- 提示信息弹窗 -->
      <uni-popup ref="message" type="message">
        <uni-popup-message :type="msgType" :message="messageText" :duration="2000"></uni-popup-message>
      </uni-popup>

      <uni-popup ref="alertDialog" type="dialog">
        <uni-popup-dialog type="error" cancelText="关闭" confirmText="删除" title="通知" content="是否要删除当前学生？"
          @confirm="delConfirm"></uni-popup-dialog>
      </uni-popup>
    </view>
  </view>
</template>

<script>
import { getUserInfoById, setUserInfo, deleteUser } from '@/api/user/user'
import { addSession } from '@/api/organization/enrollment'
import {formatDate} from '@/utils/formatDate'
export default {
  data () {
    return {
      formData: {},
      classHoure: 10,
      userId: 0,
      msgType: "success",
      messageText:"",
      useDate:formatDate(Date.now()),
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
  onLoad (options) {
    this.userId = options.userId
    this.initStudentInfo()
  },

  methods: {
    //删除学生
    del () {
      this.$refs.alertDialog.open()
    },
    delConfirm () {
      const param = {
        id: Number(this.userId)
      }
      deleteUser(param).then(res => {
        if (res.code === 0) {
          this.msgType = 'success'
          this.messageText = '删除成功'
          this.$refs.message.open()
          this.back()
        } else {
          this.msgType = 'error'
          this.messageText = res.msg
          this.$refs.message.open()
        }
      })
    },
    //增加课时
    addStudentSession () {
      const param = {
        userId: Number(this.userId),
        courseId: this.formData.eduEnrollment[0].courseId,
        sessionsToAdd: Number(this.classHoure),
        useDate:this.useDate
      }
      addSession(param).then(res => {
        if(res.code==0){
          this.$refs.addSessionDialog.close()
          this.msgType = 'success'
          this.messageText = '增加成功'
          this.$refs.message.open()
          this.initStudentInfo()
        }else{
          this.msgType = 'error'
          this.messageText = res.msg
          this.$refs.message.open()
        }
      })
    },
    //更新用户信息
    updateUserInfo () {
      this.$refs.form.validate().then(res => {
        if (res) {
          const param = this.formData
          setUserInfo(param).then(res => {
            this.$refs.studentDialog.close()
            this.msgType = 'success'
            this.messageText = '修改成功'
            this.$refs.message.open()
            this.initStudentInfo()
          })
        }
      })
    },

    //跳转到学生日志详情
    goLog () {
      uni.navigateTo({
        url: `/pages/students/enrollmentList?userId=${this.userId}&courseId=${this.formData.eduEnrollment[0].courseId}`
      });
    },

    initStudentInfo () {
      const param = {
        userId: this.userId
      }
      getUserInfoById(param).then(res => {
        this.formData = res.data.userInfo
      })
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

    back () {
      var pages = getCurrentPages();	// 获取当前页面栈
      var prePage = pages[pages.length - 2];	// 上一个页面
      prePage.data.init = true    // A 页面 init方法 为true
      setTimeout(() => {
        uni.navigateBack({ delta: 1 });    // 返回上一页
      }, 1000)
    },
  },
  watch: {},
}
</script>

<style scoped lang="scss">
.info {
  font-size: 15px;
  color: #606266;

  .content {
    margin: 40rpx;

    .item {
      display: flex;
      padding: 20px 0;
      border-bottom: 1px solid #f0f0f0;

      .title {
        width: 120px;
      }
    }

    .info-button {
      margin: 20px 0;
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