<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="userId:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户名:" prop="account">
          <el-input v-model="formData.account" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="原金额:" prop="yprice">
          <el-input-number v-model="formData.yprice" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="操作金额:" prop="price">
          <el-input-number v-model="formData.price" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="当前金额:" prop="nprice">
          <el-input-number v-model="formData.nprice" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="操作说明:" prop="memo">
          <el-input v-model="formData.memo" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="财务类型:" prop="type">
          <el-input v-model="formData.type" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="钱包类型:" prop="ptype">
          <el-input v-model="formData.ptype" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="订单号:" prop="orderId">
          <el-input v-model="formData.orderId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" @click="save">保存</el-button>
          <el-button size="small" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'FaCaiwu'
}
</script>

<script setup>
import {
  createFaCaiwu,
  updateFaCaiwu,
  findFaCaiwu
} from '@/api/faCaiwu'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const caiwuOptions = ref([])
const walletOptions = ref([])
const formData = ref({
            userId: 0,
            account: '',
            yprice: 0,
            price: 0,
            nprice: 0,
            memo: '',
            type: '',
            ptype: '',
            orderId: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findFaCaiwu({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.refaCaiwu
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    caiwuOptions.value = await getDictFunc('caiwu')
    walletOptions.value = await getDictFunc('wallet')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createFaCaiwu(formData.value)
               break
             case 'update':
               res = await updateFaCaiwu(formData.value)
               break
             default:
               res = await createFaCaiwu(formData.value)
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
