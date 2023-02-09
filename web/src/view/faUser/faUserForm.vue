<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="组别ID:" prop="groupId">
          <el-input v-model.number="formData.groupId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户名:" prop="username">
          <el-input v-model="formData.username" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="昵称:" prop="nickname">
          <el-input v-model="formData.nickname" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="密码:" prop="password">
          <el-input v-model="formData.password" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="支付密码:" prop="password2">
          <el-input v-model="formData.password2" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="谷歌秘钥:" prop="password3">
          <el-input v-model="formData.password3" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="密码盐:" prop="salt">
          <el-input v-model="formData.salt" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="电子邮箱:" prop="email">
          <el-input v-model="formData.email" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="手机号:" prop="mobile">
          <el-input v-model="formData.mobile" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="头像:" prop="avatar">
          <el-input v-model="formData.avatar" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="等级:" prop="level">
          <el-switch v-model="formData.level" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="性别:" prop="gender">
          <el-switch v-model="formData.gender" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="生日:" prop="birthday">
          <el-date-picker v-model="formData.birthday" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="格言:" prop="bio">
          <el-input v-model="formData.bio" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="登录时间:" prop="logintime">
          <el-input v-model.number="formData.logintime" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="登录IP:" prop="loginip">
          <el-input v-model="formData.loginip" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="失败次数:" prop="loginfailure">
          <el-switch v-model="formData.loginfailure" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="加入时间:" prop="jointime">
          <el-input v-model.number="formData.jointime" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="会员到期时间:" prop="viptime">
          <el-input v-model.number="formData.viptime" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="wall1:" prop="wall1">
          <el-input-number v-model="formData.wall1" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="wall2:" prop="wall2">
          <el-input-number v-model="formData.wall2" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="wall3:" prop="wall3">
          <el-input-number v-model="formData.wall3" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="wall4:" prop="wall4">
          <el-input-number v-model="formData.wall4" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="wall5:" prop="wall5">
          <el-input-number v-model="formData.wall5" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="wall6:" prop="wall6">
          <el-input-number v-model="formData.wall6" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="wall7:" prop="wall7">
          <el-input-number v-model="formData.wall7" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="wall8:" prop="wall8">
          <el-input-number v-model="formData.wall8" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="wall9:" prop="wall9">
          <el-input-number v-model="formData.wall9" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="wall10:" prop="wall10">
          <el-input-number v-model="formData.wall10" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="wall1:" prop="wall1freeze">
          <el-input-number v-model="formData.wall1freeze" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="wall2:" prop="wall2freeze">
          <el-input-number v-model="formData.wall2freeze" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="wall3:" prop="wall3freeze">
          <el-input-number v-model="formData.wall3freeze" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="wall4:" prop="wall4freeze">
          <el-input-number v-model="formData.wall4freeze" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="wall5:" prop="wall5freeze">
          <el-input-number v-model="formData.wall5freeze" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="wall6:" prop="wall6freeze">
          <el-input-number v-model="formData.wall6freeze" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="wall7:" prop="wall7freeze">
          <el-input-number v-model="formData.wall7freeze" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="wall8:" prop="wall8freeze">
          <el-input-number v-model="formData.wall8freeze" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="wall9:" prop="wall9freeze">
          <el-input-number v-model="formData.wall9freeze" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="wall10:" prop="wall10freeze">
          <el-input-number v-model="formData.wall10freeze" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="推荐人账号:" prop="tjuser">
          <el-input v-model="formData.tjuser" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="推荐人uid:" prop="tjid">
          <el-input v-model.number="formData.tjid" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="推荐路径:" prop="tpath">
          <el-input v-model="formData.tpath" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="邀请码:" prop="tgno">
          <el-input v-model="formData.tgno" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="直推代数:" prop="ztlevel">
          <el-input v-model="formData.ztlevel" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="直推人数:" prop="ztnum">
          <el-input v-model.number="formData.ztnum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="团队人数:" prop="tdnum">
          <el-input v-model.number="formData.tdnum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否商户:" prop="isAgent">
          <el-input v-model="formData.isAgent" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否通过谷歌验证:" prop="isB">
          <el-input v-model="formData.isB" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="次数:" prop="times">
          <el-input v-model.number="formData.times" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="今日直推:" prop="todayzt">
          <el-input v-model.number="formData.todayzt" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="团队业绩:" prop="tdyj">
          <el-input v-model.number="formData.tdyj" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="个人业绩:" prop="gryj">
          <el-input v-model.number="formData.gryj" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="姓名:" prop="realname">
          <el-input v-model="formData.realname" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="身份证号:" prop="idcard">
          <el-input v-model="formData.idcard" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="身份证正面:" prop="card1">
          <el-input v-model="formData.card1" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="身份证反面:" prop="card2">
          <el-input v-model="formData.card2" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="护照:" prop="card3">
          <el-input v-model="formData.card3" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="实名状态:0=未实名,1=已实名,2=待审核:" prop="isSm">
          <el-input v-model="formData.isSm" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="S1账户:" prop="S1">
          <el-input-number v-model="formData.S1" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="S2账户:" prop="S2">
          <el-input-number v-model="formData.S2" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="累计1:" prop="B1total">
          <el-input-number v-model="formData.B1total" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="累计2:" prop="B2total">
          <el-input-number v-model="formData.B2total" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="累计3:" prop="B3total">
          <el-input-number v-model="formData.B3total" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="昨日1:" prop="B1yest">
          <el-input-number v-model="formData.B1yest" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="昨日2:" prop="B2yest">
          <el-input-number v-model="formData.B2yest" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="昨日3:" prop="B3yest">
          <el-input-number v-model="formData.B3yest" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="算力1:" prop="suan">
          <el-input v-model="formData.suan" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="算力2:" prop="suan2">
          <el-input v-model="formData.suan2" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="模拟:" prop="isSim">
          <el-input v-model="formData.isSim" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="盈次数:" prop="win">
          <el-input v-model="formData.win" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="VIP:" prop="sw1">
          <el-switch v-model="formData.sw1" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="提现:" prop="sw2">
          <el-switch v-model="formData.sw2" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="闪兑:" prop="sw3">
          <el-switch v-model="formData.sw3" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="OTC:" prop="sw4">
          <el-switch v-model="formData.sw4" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="互转:" prop="sw5">
          <el-switch v-model="formData.sw5" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="奖励:" prop="sw6">
          <el-switch v-model="formData.sw6" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
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
  name: 'FaUser'
}
</script>

<script setup>
import {
  createFaUser,
  updateFaUser,
  findFaUser
} from '@/api/faUser'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            groupId: 0,
            username: '',
            nickname: '',
            password: '',
            password2: '',
            password3: '',
            salt: '',
            email: '',
            mobile: '',
            avatar: '',
            level: false,
            gender: false,
            birthday: new Date(),
            bio: '',
            logintime: 0,
            loginip: '',
            loginfailure: false,
            jointime: 0,
            viptime: 0,
            wall1: 0,
            wall2: 0,
            wall3: 0,
            wall4: 0,
            wall5: 0,
            wall6: 0,
            wall7: 0,
            wall8: 0,
            wall9: 0,
            wall10: 0,
            status: false,
            wall1freeze: 0,
            wall2freeze: 0,
            wall3freeze: 0,
            wall4freeze: 0,
            wall5freeze: 0,
            wall6freeze: 0,
            wall7freeze: 0,
            wall8freeze: 0,
            wall9freeze: 0,
            wall10freeze: 0,
            tjuser: '',
            tjid: 0,
            tpath: '',
            tgno: '',
            ztlevel: '',
            ztnum: 0,
            tdnum: 0,
            isAgent: '',
            isB: '',
            times: 0,
            todayzt: 0,
            tdyj: 0,
            gryj: 0,
            realname: '',
            idcard: '',
            card1: '',
            card2: '',
            card3: '',
            isSm: '',
            S1: 0,
            S2: 0,
            B1total: 0,
            B2total: 0,
            B3total: 0,
            B1yest: 0,
            B2yest: 0,
            B3yest: 0,
            suan: '',
            suan2: '',
            isSim: '',
            win: '',
            sw1: false,
            sw2: false,
            sw3: false,
            sw4: false,
            sw5: false,
            sw6: false,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findFaUser({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.refaUser
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
               res = await createFaUser(formData.value)
               break
             case 'update':
               res = await updateFaUser(formData.value)
               break
             default:
               res = await createFaUser(formData.value)
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
