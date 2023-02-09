<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <!-- <el-form-item label="创建时间">
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间" />
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间" />
        </el-form-item> -->
        <el-form-item label="电影" prop="flim">
          <el-select v-model="searchInfo.typeId" clearable placeholder="请选择" @clear="()=>{searchInfo.typeId=undefined}">
            <el-option v-for="(item,key) in flimOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="综艺" prop="zongyi">
          <el-select v-model="searchInfo.typeId" clearable placeholder="请选择" @clear="()=>{searchInfo.typeId=undefined}">
            <el-option v-for="(item,key) in zongyiOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="连续剧" prop="lianxuju">
          <el-select v-model="searchInfo.typeId" clearable placeholder="请选择" @clear="()=>{searchInfo.typeId=undefined}">
            <el-option v-for="(item,key) in lianxujuOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="动漫" prop="dongman">
          <el-select v-model="searchInfo.typeId" clearable placeholder="请选择" @clear="()=>{searchInfo.typeId=undefined}">
            <el-option v-for="(item,key) in dongmanOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
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
        <el-table-column align="left" label="资源id" prop="vodId" width="120" />
        <el-table-column align="left" label="类型" prop="typeName" width="120" />
        <el-table-column align="left" label="类型id" prop="typeId" width="120" />
        <el-table-column align="left" label="类型id1" prop="typeId1" width="120" />
        <!-- <el-table-column align="left" label="group_id" prop="groupId" width="120" /> -->
        <el-table-column align="left" label="图片" prop="vodPic" width="150">
          <template #default="scope"> <img :src="scope.row.vodPic" alt="" width="140" height="220"></template>
        </el-table-column>
        <el-table-column align="left" label="标题" prop="vodName" width="220" />
        <el-table-column align="left" label="简介" prop="vodSub" width="520" />
        <el-table-column align="left" label="拼音" prop="vodEn" width="120" />
        <el-table-column align="left" label="状态" prop="vodStatus" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.vodStatus) }}</template>
        </el-table-column>
        <el-table-column align="left" label="字母" prop="vodLetter" width="120" />
        <!-- <el-table-column align="left" label="标签" prop="vodTag" width="120" /> -->
        <el-table-column align="left" label="类型" prop="vodClass" width="120" />

        <!-- <el-table-column align="left" label="图片" prop="vodPicScreenshot" width="120" /> -->
        <el-table-column align="left" label="演员" prop="vodActor" width="620" />
        <el-table-column align="left" label="导演" prop="vodDirector" width="120" />
        <el-table-column align="left" label="编剧" prop="vodWriter" width="120" />
        <!-- <el-table-column align="left" label="后台" prop="vodBehind" width="120" /> -->
        <el-table-column align="left" label="简介" prop="vodBlurb" width="920" />
        <el-table-column align="left" label="书签" prop="vodRemarks" width="120" />
        <el-table-column align="left" label="发行日期" prop="vodPubdate" width="120" />
        <el-table-column align="left" label="总计" prop="vodTotal" width="120" />
        <el-table-column align="left" label="序列" prop="vodSerial" width="120" />
        <el-table-column align="left" label="区域" prop="vodArea" width="120" />
        <el-table-column align="left" label="语言" prop="vodLang" width="120" />
        <el-table-column align="left" label="年代" prop="vodYear" width="120" />
        <el-table-column align="left" label="vodHits字段" prop="vodHits" width="120" />
        <el-table-column align="left" label="vodHitsDay字段" prop="vodHitsDay" width="120" />
        <el-table-column align="left" label="vodHitsWeek字段" prop="vodHitsWeek" width="120" />
        <el-table-column align="left" label="vodHitsMonth字段" prop="vodHitsMonth" width="120" />
        <el-table-column align="left" label="时间" prop="vodDuration" width="120" />
        <el-table-column align="left" label="vodUp字段" prop="vodUp" width="120" />
        <el-table-column align="left" label="vodDown字段" prop="vodDown" width="120" />
        <el-table-column align="left" label="vodScore字段" prop="vodScore" width="120" />
        <el-table-column align="left" label="vodScoreAll字段" prop="vodScoreAll" width="120" />
        <el-table-column align="left" label="vodScoreNum字段" prop="vodScoreNum" width="120" />
        <el-table-column align="left" label="vodTime字段" prop="vodTime" width="220" />
        <!-- <el-table-column align="left" label="vodTimeAdd字段" prop="vodTimeAdd" width="120" /> -->
        <!-- <el-table-column align="left" label="vodDoubanId字段" prop="vodDoubanId" width="120" /> -->
        <el-table-column align="left" label="vodDoubanScore字段" prop="vodDoubanScore" width="120" />
        <!-- <el-table-column align="left" label="详情" prop="vodContent" width="120" /> -->
        <!-- <el-table-column align="left" label="vodPlayUrl字段" prop="vodPlayUrl" width="120" /> -->
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateVideoDetailsFunc(scope.row)">变更</el-button>
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
        <el-form-item label="资源id:" prop="vodId">
          <el-input v-model.number="formData.vodId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="类型:" prop="typeName">
          <el-input v-model="formData.typeName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="类型id:" prop="typeId">
          <el-input v-model.number="formData.typeId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="类型id1:" prop="typeId1">
          <el-input v-model.number="formData.typeId1" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="group_id:" prop="groupId">
          <el-input v-model.number="formData.groupId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="标题:" prop="vodName">
          <el-input v-model="formData.vodName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="简介:" prop="vodSub">
          <el-input v-model="formData.vodSub" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="拼音:" prop="vodEn">
          <el-input v-model="formData.vodEn" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:" prop="vodStatus">
          <el-switch v-model="formData.vodStatus" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable />
        </el-form-item>
        <el-form-item label="字母:" prop="vodLetter">
          <el-input v-model="formData.vodLetter" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="标签:" prop="vodTag">
          <el-input v-model="formData.vodTag" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="类型:" prop="vodClass">
          <el-input v-model="formData.vodClass" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="图片:" prop="vodPic">
          <el-input v-model="formData.vodPic" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="图片:" prop="vodPicScreenshot">
          <el-input v-model="formData.vodPicScreenshot" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="演员:" prop="vodActor">
          <el-input v-model="formData.vodActor" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="导演:" prop="vodDirector">
          <el-input v-model="formData.vodDirector" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="编剧:" prop="vodWriter">
          <el-input v-model="formData.vodWriter" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="后台:" prop="vodBehind">
          <el-input v-model="formData.vodBehind" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="简介:" prop="vodBlurb">
          <el-input v-model="formData.vodBlurb" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="书签:" prop="vodRemarks">
          <el-input v-model="formData.vodRemarks" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="发行日期:" prop="vodPubdate">
          <el-input v-model="formData.vodPubdate" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="总计:" prop="vodTotal">
          <el-input v-model.number="formData.vodTotal" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="序列:" prop="vodSerial">
          <el-input v-model="formData.vodSerial" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="区域:" prop="vodArea">
          <el-input v-model="formData.vodArea" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="语言:" prop="vodLang">
          <el-input v-model="formData.vodLang" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="年代:" prop="vodYear">
          <el-input v-model="formData.vodYear" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="vodHits字段:" prop="vodHits">
          <el-input v-model.number="formData.vodHits" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="vodHitsDay字段:" prop="vodHitsDay">
          <el-input v-model.number="formData.vodHitsDay" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="vodHitsWeek字段:" prop="vodHitsWeek">
          <el-input v-model.number="formData.vodHitsWeek" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="vodHitsMonth字段:" prop="vodHitsMonth">
          <el-input v-model.number="formData.vodHitsMonth" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="时间:" prop="vodDuration">
          <el-input v-model="formData.vodDuration" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="vodUp字段:" prop="vodUp">
          <el-input v-model.number="formData.vodUp" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="vodDown字段:" prop="vodDown">
          <el-input v-model.number="formData.vodDown" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="vodScore字段:" prop="vodScore">
          <el-input v-model="formData.vodScore" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="vodScoreAll字段:" prop="vodScoreAll">
          <el-input v-model.number="formData.vodScoreAll" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="vodScoreNum字段:" prop="vodScoreNum">
          <el-input v-model.number="formData.vodScoreNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="vodTime字段:" prop="vodTime">
          <el-input v-model="formData.vodTime" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="vodTimeAdd字段:" prop="vodTimeAdd">
          <el-input v-model.number="formData.vodTimeAdd" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="vodDoubanId字段:" prop="vodDoubanId">
          <el-input v-model.number="formData.vodDoubanId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="vodDoubanScore字段:" prop="vodDoubanScore">
          <el-input v-model="formData.vodDoubanScore" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="详情:" prop="vodContent">
          <el-input v-model="formData.vodContent" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="vodPlayUrl字段:" prop="vodPlayUrl">
          <el-input v-model="formData.vodPlayUrl" :clearable="true" placeholder="请输入" />
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
  name: 'VideoDetails'
}
</script>

<script setup>
import {
  createVideoDetails,
  deleteVideoDetails,
  deleteVideoDetailsByIds,
  updateVideoDetails,
  findVideoDetails,
  getVideoDetailsList
} from '@/api/videoDetails'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const flimOptions = ref([])
const zongyiOptions = ref([])
const lianxujuOptions = ref([])
const dongmanOptions = ref([])
const formData = ref({
  vodId: 0,
  typeName: '',
  typeId: 0,
  typeId1: 0,
  groupId: 0,
  vodName: '',
  vodSub: '',
  vodEn: '',
  vodStatus: false,
  vodLetter: '',
  vodTag: '',
  vodClass: '',
  vodPic: '',
  vodPicScreenshot: '',
  vodActor: '',
  vodDirector: '',
  vodWriter: '',
  vodBehind: '',
  vodBlurb: '',
  vodRemarks: '',
  vodPubdate: '',
  vodTotal: 0,
  vodSerial: '',
  vodArea: '',
  vodLang: '',
  vodYear: '',
  vodHits: 0,
  vodHitsDay: 0,
  vodHitsWeek: 0,
  vodHitsMonth: 0,
  vodDuration: '',
  vodUp: 0,
  vodDown: 0,
  vodScore: '',
  vodScoreAll: 0,
  vodScoreNum: 0,
  vodTime: '',
  vodTimeAdd: 0,
  vodDoubanId: 0,
  vodDoubanScore: '',
  vodContent: '',
  vodPlayUrl: '',
  // net: undefined,
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
  getTableData()
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  if (searchInfo.value.vodStatus === '') {
    searchInfo.value.vodStatus = null
  }
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
  const table = await getVideoDetailsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
const setOptions = async() => {
  flimOptions.value = await getDictFunc('flim')
  zongyiOptions.value = await getDictFunc('zongyi')
  lianxujuOptions.value = await getDictFunc('lianxuju')
  dongmanOptions.value = await getDictFunc('dongman')
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
    deleteVideoDetailsFunc(row)
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
  const res = await deleteVideoDetailsByIds({ ids })
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
const updateVideoDetailsFunc = async(row) => {
  const res = await findVideoDetails({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.revideoDetails
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteVideoDetailsFunc = async(row) => {
  const res = await deleteVideoDetails({ ID: row.ID })
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
    vodId: 0,
    typeName: '',
    typeId: 0,
    typeId1: 0,
    groupId: 0,
    vodName: '',
    vodSub: '',
    vodEn: '',
    vodStatus: false,
    vodLetter: '',
    vodTag: '',
    vodClass: '',
    vodPic: '',
    vodPicScreenshot: '',
    vodActor: '',
    vodDirector: '',
    vodWriter: '',
    vodBehind: '',
    vodBlurb: '',
    vodRemarks: '',
    vodPubdate: '',
    vodTotal: 0,
    vodSerial: '',
    vodArea: '',
    vodLang: '',
    vodYear: '',
    vodHits: 0,
    vodHitsDay: 0,
    vodHitsWeek: 0,
    vodHitsMonth: 0,
    vodDuration: '',
    vodUp: 0,
    vodDown: 0,
    vodScore: '',
    vodScoreAll: 0,
    vodScoreNum: 0,
    vodTime: '',
    vodTimeAdd: 0,
    vodDoubanId: 0,
    vodDoubanScore: '',
    vodContent: '',
    vodPlayUrl: '',
  }
}
// 弹窗确定
const enterDialog = async() => {
     elFormRef.value?.validate(async(valid) => {
       if (!valid) return
       let res
       switch (type.value) {
         case 'create':
           res = await createVideoDetails(formData.value)
           break
         case 'update':
           res = await updateVideoDetails(formData.value)
           break
         default:
           res = await createVideoDetails(formData.value)
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
