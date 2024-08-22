<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="报名ID:" prop="enrollmentId">
          <el-input v-model.number="formData.enrollmentId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="操作类型（增加或扣除）:" prop="action">
        <el-select v-model="formData.action" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in ['add', 'subtract']" :key="item" :label="item" :value="item" />
        </el-select>
        </el-form-item>
        <el-form-item label="操作原因:" prop="reason">
          <el-input v-model="formData.reason" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="课时数量:" prop="numSessions">
          <el-input v-model.number="formData.numSessions" :clearable="true" placeholder="请输入" />
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
  name: 'EduClassSession'
}
</script>

<script setup>
import {
  createEduClassSession,
  updateEduClassSession,
  findEduClassSession
} from '@/api/eduClassSession'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            enrollmentId: 0,
            reason: '',
            numSessions: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findEduClassSession({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reeduClassSession
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createEduClassSession(formData.value)
               break
             case 'update':
               res = await updateEduClassSession(formData.value)
               break
             default:
               res = await createEduClassSession(formData.value)
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

<style>
</style>
