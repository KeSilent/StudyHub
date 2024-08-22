<!--
 * @Author: Yang
 * @Date: 2019-08-22 19:41:20
 * @Description: 首页
-->
<template>
  <view class="content">
    <view class="card-title">
      <view class="title">🐱学时印记🐱</view>
      <view class="content">欢迎使用🐱学时印记🐱{{ list.length > 0 ? "，以下学生剩余课程少于5节。" : "" }}</view>
    </view>
    <view class="text-area">
      <uni-table border stripe emptyText="暂无更多数据" @sort-change="sortChange()" v-if="list.length > 0">
        <uni-tr>
          <uni-th align="center">学生姓名</uni-th>
          <uni-th align="center">课时</uni-th>
          <!-- <uni-th width="100" align="center">电话</uni-th> -->
        </uni-tr>

        <uni-tr v-for="(item, index) in list" :key="index">
          <uni-td align="center">
            <view @click="showEditStudent(item)">{{ item.nickName }}</view>
          </uni-td>
          <uni-td align="center">
            <view>{{ item.remainingSessions }}</view>
          </uni-td>
          <uni-td align="center">
            <view>{{ item.phone }}</view>
          </uni-td>
        </uni-tr>
      </uni-table>
    </view>
  </view>
</template>

<script>
import { getStudentsWithLessThanFiveSessions } from '@/api/organization/enrollment'
import Vue from 'vue';
export default {
  data () {
    return {
      list: [],
    }
  },
  created () {
    this.init()
  },
  methods: {
    init () {
      let token = uni.getStorageSync('token')
      if (token == "") {
        //如果没用登录则，跳转到初始化页面
        uni.redirectTo({
          url: '/pages/login/index'
        });
      } else {
        let user = uni.getStorageSync('user')
        if (user.authorityId == Vue.prototype.studentRoleId) {
          //学生家长角色
          uni.redirectTo({
            url: '/pages/clientInfo/index'
          });
        } else {
          this.initData()
        }
      }

    },
    //显示学生修改弹窗
    showEditStudent (student) {
      uni.navigateTo({
        url: '/pages/students/info?userId=' + student.ID
      });
    },
    initData () {
      getStudentsWithLessThanFiveSessions().then((resp) => {
        if (resp.data.list.length > 0) {
          this.list = resp.data.list
        }
      }).catch(err => {
        console.log(err)
      })
    }
  }
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
}
</style>
