<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="创建时间">
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间" />
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间" />
        </el-form-item>
        <el-form-item label="userId">

          <el-input v-model.number="searchInfo.userId" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="财务类型" prop="type">
          <el-select v-model="searchInfo.type" clearable placeholder="请选择" @clear="()=>{searchInfo.type=undefined}">
            <el-option v-for="(item,key) in caiwuOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="钱包类型" prop="ptype">
          <el-select v-model="searchInfo.ptype" clearable placeholder="请选择" @clear="()=>{searchInfo.ptype=undefined}">
            <el-option v-for="(item,key) in walletOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="订单号">
          <el-input v-model="searchInfo.orderId" placeholder="搜索条件" />

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
        <el-button size="small">金额总计：{{ sum }}</el-button>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="userId" prop="userId" width="120" />
        <el-table-column align="left" label="用户名" prop="account" width="120" />
        <el-table-column align="left" label="原金额" prop="yprice" width="120" />
        <el-table-column align="left" label="操作金额" prop="price" width="120" />
        <el-table-column align="left" label="当前金额" prop="nprice" width="120" />
        <el-table-column align="left" label="操作说明" prop="memo" width="120" />
        <el-table-column sortable align="left" label="财务类型" prop="type" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.type,caiwuOptions) }}
          </template>
        </el-table-column>
        <el-table-column sortable align="left" label="钱包类型" prop="ptype" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.ptype,walletOptions) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="订单号" prop="orderId" width="120" />
        <el-table-column align="left" label="操作" min-width="150" fixed="right">
          <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateFaCaiwuFunc(scope.row)">变更</el-button>
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
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="userId:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户名:" prop="account">
          <el-input v-model="formData.account" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="原金额:" prop="yprice">
          <el-input-number v-model="formData.yprice" style="width:100%" :precision="4" :clearable="true" />
        </el-form-item>
        <el-form-item label="操作金额:" prop="price">
          <el-input-number v-model="formData.price" style="width:100%" :precision="4" :clearable="true" />
        </el-form-item>
        <el-form-item label="当前金额:" prop="nprice">
          <el-input-number v-model="formData.nprice" style="width:100%" :precision="4" :clearable="true" />
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
  name: 'FaCaiwu'
}
</script>

<script setup>
import {
  createFaCaiwu,
  deleteFaCaiwu,
  deleteFaCaiwuByIds,
  updateFaCaiwu,
  findFaCaiwu,
  getFaCaiwuList
} from '@/api/faCaiwu'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
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

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const sum = ref(0)
// 排序
const sortChange = ({ prop, order }) => {
  searchInfo.value.sort = prop
  searchInfo.value.order = order
  getTableData()
}

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
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
  const table = await getFaCaiwuList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    sum.value = table.data.price // 统计
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {
  caiwuOptions.value = await getDictFunc('caiwu')
  walletOptions.value = await getDictFunc('wallet')
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
    deleteFaCaiwuFunc(row)
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
  const res = await deleteFaCaiwuByIds({ ids })
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
const updateFaCaiwuFunc = async(row) => {
  const res = await findFaCaiwu({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.refaCaiwu
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteFaCaiwuFunc = async(row) => {
  const res = await deleteFaCaiwu({ ID: row.ID })
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
    account: '',
    yprice: 0,
    price: 0,
    nprice: 0,
    memo: '',
    type: '',
    ptype: '',
    orderId: '',
  }
}
// 弹窗确定
const enterDialog = async() => {
     elFormRef.value?.validate(async(valid) => {
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
         closeDialog()
         getTableData()
       }
     })
}
</script>

<style>
</style>
