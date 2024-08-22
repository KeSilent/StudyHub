<!--
 * @Author: Yang
 * @Date: 2023-04-10 18:22:12
 * @Description: 请填写简介
-->
<template>
  <div style="height:80vh">
    <iframe width="100%" height="100%" :src="`${basePath}/form-generator/#/`" frameborder="0" />
  </div>
</template>

<script>
export default {
  name: 'FormGenerator'
}
</script>

<script setup>
import { ref } from 'vue'
const basePath = ref(import.meta.env.VITE_BASE_PATH)
const basePort = ref(import.meta.env.VITE_SERVER_PORT)
</script>
