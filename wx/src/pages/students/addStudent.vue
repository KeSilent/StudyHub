<template>
  <view class="addStudent">
    <uni-forms label-width="80px" ref="form" :modelValue="formData" :rules="rules">
      <uni-forms-item required label="姓名" name="name">
        <uni-easyinput type="text" v-model="formData.name" placeholder="请输入姓名" />
      </uni-forms-item>
      <uni-forms-item required name="course" label="科目">
        <uni-data-select required v-model="formData.course" :localdata="courseList" :clear="false"></uni-data-select>
      </uni-forms-item>
      <uni-forms-item required label="总课时" name="totalSessions">
        <uni-easyinput type="text" v-model="formData.totalSessions" placeholder="请输入课时"
          @input="formData.remainingSessions = formData.totalSessions" />
      </uni-forms-item>
      <uni-forms-item required label="剩余课时" name="remainingSessions">
        <uni-easyinput type="text" v-model="formData.remainingSessions" placeholder="请输入剩余课时" />
      </uni-forms-item>
    </uni-forms>
    <button @click="submitForm">创建学生</button>
  </view>
</template>

<script>
import { getEduCourseList } from '@/api/organization/course'
import { createUser } from '@/api/user/user'
export default {
  data () {
    return {
      formData: {
        name: '',
        course: '',
        phone: '',
        totalSessions: '',
        remainingSessions: ''
      },
      courseList: [],
      rules: {
        name: {
          rules: [{
            required: true,
            errorMessage: '请输入姓名',
          }
          ]
        },
        course: {
          rules: [{
            required: true,
            errorMessage: '请选择科目',
          }
          ]
        },
        totalSessions: {
          rules: [{
            required: true,
            errorMessage: '请选输入总课时',
          }
          ]
        },
        remainingSessions: {
          rules: [{
            required: true,
            errorMessage: '请选输入剩余课时',
          }
          ]
        }
      }
    }
  },
  created () {
    this.initEduCourseList()
  },
  methods: {
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

    //提交表单
    submitForm () {
      uni.showLoading()
      this.$refs.form.validate().then(res => {
        uni.hideLoading()
        console.log('表单数据信息：', res);
        createUser(res.name, res.course, res.totalSessions, res.remainingSessions, res.phone, this.studentRoleId).then(res => {
          if (res.code === 0) {
            uni.showToast({
              title: '创建成功',
              icon: 'success',
              duration: 2000
            })
            setTimeout(() => {
              uni.reLaunch({
                url: "/pages/students/index",
                success: function (e) {
                  var page = getCurrentPages()[0];
                  if (page == undefined || page == null) return;
                  page.onLoad(); //或者其它操作
                }
              });
            }, 500)
          }
        })
      }).catch(err => {
        uni.hideLoading()
        console.log('表单错误信息：', err);
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
.addStudent {
  margin: 40rpx;
}
</style>