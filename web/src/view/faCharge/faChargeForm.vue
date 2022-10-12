<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="会员id:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商户id:" prop="shangId">
          <el-input v-model.number="formData.shangId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="转账地址:" prop="from">
          <el-input v-model="formData.from" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="收账地址:" prop="to">
          <el-input v-model="formData.to" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="金额:" prop="number">
          <el-input-number v-model="formData.number" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="转账费用:" prop="fee">
          <el-input-number v-model="formData.fee" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="区块交易id:" prop="txid">
          <el-input v-model="formData.txid" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="交易时间:" prop="createTime">
          <el-input v-model.number="formData.createTime" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-select v-model="formData.status" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="币种:" prop="coin">
          <el-input v-model="formData.coin" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="日志:" prop="tradelog">
          <el-input v-model="formData.tradelog" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="app订单号:" prop="orderSn">
          <el-input v-model="formData.orderSn" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="contract字段:" prop="contract">
          <el-input v-model="formData.contract" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="质押周期:" prop="cycle">
          <el-input v-model.number="formData.cycle" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="到期时间:" prop="endTime">
          <el-input v-model.number="formData.endTime" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="通知状态:" prop="notice">
          <el-select v-model="formData.notice" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
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
  name: 'FaCharge'
}
</script>

<script setup>
import {
  createFaCharge,
  updateFaCharge,
  findFaCharge
} from '@/api/faCharge'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const statusOptions = ref([])
const formData = ref({
            userId: 0,
            shangId: 0,
            from: '',
            to: '',
            number: 0,
            fee: 0,
            txid: '',
            createTime: 0,
            status: undefined,
            coin: '',
            tradelog: '',
            orderSn: '',
            contract: '',
            cycle: 0,
            endTime: 0,
            notice: undefined,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findFaCharge({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.refaCharge
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    statusOptions.value = await getDictFunc('status')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createFaCharge(formData.value)
               break
             case 'update':
               res = await updateFaCharge(formData.value)
               break
             default:
               res = await createFaCharge(formData.value)
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
