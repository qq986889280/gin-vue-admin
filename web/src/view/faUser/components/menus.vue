<template>
  <div>
    <!-- <div class="clearfix sticky-button">
      <el-input v-model="filterText" class="fitler" placeholder="筛选" />
    </div> -->
    <div class="tree-content">
      <!-- :default-checked-keys="menuTreeIds" -->
      <!-- show-checkbox -->
      <!-- node-key="id" -->
      <el-tree ref="menuTree" :data="menuTreeData" :props="menuDefaultProps" default-expand-all highlight-current
        :filter-node-method="filterNode" @check="nodeChange">
        <template #default="{ node, data }">
          <span class="custom-tree-node">
            <span>{{ node.ID }}</span>
            <span>{{ data.username }} 个人业绩：{{ data.tdyj }} 团队业绩：{{ data.tdyj }}</span>
          </span>
        </template>
      </el-tree>
    </div>

  </div>
</template>

<script setup>
import { userTeam } from '@/api/faUser'
import {
  updateAuthority
} from '@/api/authority'
import { getAuthorityBtnApi, setAuthorityBtnApi } from '@/api/authorityBtn'
import { nextTick, ref, watch } from 'vue'
import { ElMessage } from 'element-plus'

const props = defineProps({
  row: {
    default: function () {
      return {}
    },
    type: Object
  }
})

const emit = defineEmits(['changeRow'])
const filterText = ref('')
const menuTreeData = ref([])
const menuTreeIds = ref([])
const needConfirm = ref(false)
const menuDefaultProps = ref({
  children: 'children',
  label: function (data) {
    return data.meta.title
  }
})

const init = async () => {
  // 获取所有菜单树
  const res = await userTeam({ id: props.row.ID })
  menuTreeData.value = res.data
  // const res1 = await getMenuAuthority({ authorityId: props.row.authorityId })
  // const menus = res1.data.menus
  // const arr = []
  // menus.forEach(item => {
  //   // 防止直接选中父级造成全选
  //   if (!menus.some(same => same.parentId === item.menuId)) {
  //     arr.push(Number(item.menuId))
  //   }
  // })
  // menuTreeIds.value = arr
}

init()

const setDefault = async (data) => {
  const res = await updateAuthority({ authorityId: props.row.authorityId, AuthorityName: props.row.authorityName, parentId: props.row.parentId, defaultRouter: data.name })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: '设置成功' })
    emit('changeRow', 'defaultRouter', res.data.authority.defaultRouter)
  }
}
const nodeChange = () => {
  needConfirm.value = true
}
// 暴露给外层使用的切换拦截统一方法
const enterAndNext = () => {
  relation()
}
// 关联树 确认方法
const menuTree = ref(null)
const relation = async () => {
  const checkArr = menuTree.value.getCheckedNodes(false, true)
  const res = await addMenuAuthority({
    menus: checkArr,
    authorityId: props.row.authorityId
  })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '菜单设置成功!'
    })
  }
}

defineExpose({ enterAndNext, needConfirm })

const btnVisible = ref(false)

const btnData = ref([])
const multipleSelection = ref([])
const btnTableRef = ref()
let menuID = ''
const OpenBtn = async (data) => {
  menuID = data.ID
  const res = await getAuthorityBtnApi({ menuID: menuID, authorityId: props.row.authorityId })
  if (res.code === 0) {
    openDialog(data)
    await nextTick()
    if (res.data.selected) {
      res.data.selected.forEach(id => {
        btnData.value.some(item => {
          if (item.ID === id) {
            btnTableRef.value.toggleRowSelection(item, true)
          }
        })
      })
    }
  }
}

const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

const openDialog = (data) => {
  btnVisible.value = true
  btnData.value = data.menuBtn
}

const closeDialog = () => {
  btnVisible.value = false
}
const enterDialog = async () => {
  const selected = multipleSelection.value.map(item => item.ID)
  const res = await setAuthorityBtnApi({
    menuID,
    selected,
    authorityId: props.row.authorityId
  })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: '设置成功' })
    btnVisible.value = false
  }
}

const filterNode = (value, data) => {
  if (!value) return true
  console.log(data.username)
  return data.username.indexOf(value) !== -1
}

watch(filterText, (val) => {
  menuTree.value.filter(val)
})

</script>

<script>

export default {
  name: 'Menus'
}
</script>

<style lang="scss" scope>
@import "@/style/main.scss";

.custom-tree-node {
  span+span {
    margin-left: 12px;
  }
}
</style>
