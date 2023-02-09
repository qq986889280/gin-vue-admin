<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
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
          <el-switch v-model="formData.vodStatus" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
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
        <el-form-item label="vodContent字段:" prop="vodContent">
          <el-input v-model="formData.vodContent" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="vodPlayUrl字段:" prop="vodPlayUrl">
          <el-input v-model="formData.vodPlayUrl" :clearable="true" placeholder="请输入" />
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
  name: 'VideoDetails'
}
</script>

<script setup>
import {
  createVideoDetails,
  updateVideoDetails,
  findVideoDetails
} from '@/api/videoDetails'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
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
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findVideoDetails({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.revideoDetails
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
