<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
      <el-form-item label="创建时间">
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
      </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button size="small" type="primary" link @click="deleteVisible = false">取消</el-button>
                <el-button size="small" type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="会员id" prop="userId" width="120" />
        <el-table-column align="left" label="商户id" prop="shangId" width="120" />
        <el-table-column align="left" label="转账地址" prop="from" width="120" />
        <el-table-column align="left" label="收账地址" prop="to" width="120" />
        <el-table-column align="left" label="金额" prop="number" width="120" />
        <el-table-column align="left" label="转账费用" prop="fee" width="120" />
        <el-table-column align="left" label="区块交易id" prop="txid" width="120" />
        <el-table-column align="left" label="交易时间" prop="createTime" width="120" />
        <el-table-column align="left" label="状态" prop="status" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.status,statusOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="币种" prop="coin" width="120" />
        <el-table-column align="left" label="日志" prop="tradelog" width="120" />
        <el-table-column align="left" label="app订单号" prop="orderSn" width="120" />
        <el-table-column align="left" label="contract字段" prop="contract" width="120" />
        <el-table-column align="left" label="质押周期" prop="cycle" width="120" />
        <el-table-column align="left" label="到期时间" prop="endTime" width="120" />
        <el-table-column align="left" label="通知状态" prop="notice" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.notice,statusOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateFaChargeFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="会员id:"  prop="userId" >
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商户id:"  prop="shangId" >
          <el-input v-model.number="formData.shangId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="转账地址:"  prop="from" >
          <el-input v-model="formData.from" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="收账地址:"  prop="to" >
          <el-input v-model="formData.to" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="金额:"  prop="number" >
          <el-input-number v-model="formData.number"  style="width:100%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="转账费用:"  prop="fee" >
          <el-input-number v-model="formData.fee"  style="width:100%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="区块交易id:"  prop="txid" >
          <el-input v-model="formData.txid" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="交易时间:"  prop="createTime" >
          <el-input v-model.number="formData.createTime" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:"  prop="status" >
          <el-select v-model="formData.status" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="币种:"  prop="coin" >
          <el-input v-model="formData.coin" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="日志:"  prop="tradelog" >
          <el-input v-model="formData.tradelog" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="app订单号:"  prop="orderSn" >
          <el-input v-model="formData.orderSn" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="contract字段:"  prop="contract" >
          <el-input v-model="formData.contract" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="质押周期:"  prop="cycle" >
          <el-input v-model.number="formData.cycle" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="到期时间:"  prop="endTime" >
          <el-input v-model.number="formData.endTime" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="通知状态:"  prop="notice" >
          <el-select v-model="formData.notice" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
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
  deleteFaCharge,
  deleteFaChargeByIds,
  updateFaCharge,
  findFaCharge,
  getFaChargeList
} from '@/api/faCharge'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
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


// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getFaChargeList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    statusOptions.value = await getDictFunc('status')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteFaChargeFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteFaChargeByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateFaChargeFunc = async(row) => {
    const res = await findFaCharge({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.refaCharge
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteFaChargeFunc = async (row) => {
    const res = await deleteFaCharge({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}
</script>

<style>
</style>
