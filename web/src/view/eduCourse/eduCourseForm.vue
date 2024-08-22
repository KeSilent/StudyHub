<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="组织:" prop="organizationId">
          <el-select v-model="formData.organizationId" class="m-2" placeholder="请选择组织" size="large">
            <el-option v-for="item in organizationList" :key="item.ID" :label="item.name" :value="item.ID" />
          </el-select>
        </el-form-item>
        <el-form-item label="管理员ID:" prop="adminId">
          <el-input v-model.number="formData.adminId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="课程名称:" prop="courseName">
          <el-input v-model="formData.courseName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="课程描述:" prop="description">
          <el-input v-model="formData.description" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="课程展示图URL:" prop="imageUrl">
          <el-input v-model="formData.imageUrl" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'EduCourse'
}
</script>

<script setup>
import {
  createEduCourse,
  updateEduCourse,
  findEduCourse
} from '@/api/eduCourse'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
import { getEduOrganizationList } from "@/api/eduOrganization";
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
  organizationId: 0,
  adminId: 0,
  courseName: '',
  description: '',
  imageUrl: '',
})
// 验证规则
const rule = reactive({
})
//组织选项
const organizationList = ref([])

//获取组织数据
const getOrganizationList = async () => {
 const result = await getEduOrganizationList()
  organizationList.value = result.data.list
}
getOrganizationList()

const elFormRef = ref()

// 初始化方法
const init = async () => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findEduCourse({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.reeduCourse
      type.value = 'update'
    }
  } else {
    type.value = 'create'
  }


  getEduOrganizationList().then(res => {
    organizationList.value = res.data
    console.log(organizationList.value);
  })
}


init()
// 保存按钮
const save = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createEduCourse(formData.value)
        break
      case 'update':
        res = await updateEduCourse(formData.value)
        break
      default:
        res = await createEduCourse(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
    }
  })
}

// 返回按钮
const back = () => {
  router.go(-1)
}

</script>

<style></style>
