<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="创建时间">
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间" />
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间" />
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
        <el-table-column align="left" label="按钮组" width="150">
          <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateFaUserFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
        <el-table-column align="left" label="状态">
          <template #default="scope">
            <el-switch
              v-model="scope.row.status"
              inline-prompt
              :active-value="1"
              :inactive-value="2"
              @change="()=>{switchEnable(scope.row)}"
            />
          </template>
        </el-table-column>
        <el-table-column align="left" label="VIP">
          <template #default="scope">
            <el-switch
              v-model="scope.row.sw1"
              inline-prompt
              :active-value="1"
              :inactive-value="2"
              @change="()=>{switchEnable(scope.row)}"
            />
          </template>
        </el-table-column>
        <el-table-column align="left" label="提现">
          <template #default="scope">
            <el-switch
              v-model="scope.row.sw2"
              inline-prompt
              :active-value="1"
              :inactive-value="2"
              @change="()=>{switchEnable(scope.row)}"
            />
          </template>
        </el-table-column>
        <el-table-column align="left" label="闪兑">
          <template #default="scope">
            <el-switch
              v-model="scope.row.sw3"
              inline-prompt
              :active-value="1"
              :inactive-value="2"
              @change="()=>{switchEnable(scope.row)}"
            />
          </template>
        </el-table-column>
        <el-table-column align="left" label="OTC">
          <template #default="scope">
            <el-switch
              v-model="scope.row.sw4"
              inline-prompt
              :active-value="1"
              :inactive-value="2"
              @change="()=>{switchEnable(scope.row)}"
            />
          </template>
        </el-table-column>
        <el-table-column align="left" label="互转">
          <template #default="scope">
            <el-switch
              v-model="scope.row.sw5"
              inline-prompt
              :active-value="1"
              :inactive-value="2"
              @change="()=>{switchEnable(scope.row)}"
            />
          </template>
        </el-table-column>
        <el-table-column align="left" label="奖励">
          <template #default="scope">
            <el-switch
              v-model="scope.row.sw6"
              inline-prompt
              :active-value="1"
              :inactive-value="2"
              @change="()=>{switchEnable(scope.row)}"
            />
          </template>
        </el-table-column>
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>

        <el-table-column align="left" label="组别ID" prop="groupId" width="120" />
        <el-table-column align="left" label="用户名" prop="username" width="120" />
        <el-table-column align="left" label="昵称" prop="nickname" width="120" />
        <!-- <el-table-column align="left" label="密码" prop="password" width="120" /> -->
        <!-- <el-table-column align="left" label="支付密码" prop="password2" width="120" /> -->
        <!-- <el-table-column align="left" label="谷歌秘钥" prop="password3" width="120" /> -->
        <!-- <el-table-column align="left" label="密码盐" prop="salt" width="120" /> -->
        <el-table-column align="left" label="电子邮箱" prop="email" width="120" />
        <el-table-column align="left" label="手机号" prop="mobile" width="120" />
        <el-table-column align="left" label="头像" prop="avatar" width="120" />
        <el-table-column align="left" label="等级" prop="level" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.level) }}</template>
        </el-table-column>
        <el-table-column align="left" label="性别" prop="gender" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.gender) }}</template>
        </el-table-column>
        <el-table-column align="left" label="生日" width="180">
          <template #default="scope">{{ formatDate(scope.row.birthday) }}</template>
        </el-table-column>
        <el-table-column align="left" label="格言" prop="bio" width="120" />
        <el-table-column align="left" label="登录时间" prop="logintime" width="120" />
        <el-table-column align="left" label="登录IP" prop="loginip" width="120" />
        <el-table-column align="left" label="失败次数" prop="loginfailure" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.loginfailure) }}</template>
        </el-table-column>
        <el-table-column align="left" label="加入时间" prop="jointime" width="120" />
        <el-table-column align="left" label="会员到期时间" prop="viptime" width="120" />
        <el-table-column align="left" label="wall1" prop="wall1" width="120" />
        <el-table-column align="left" label="wall2" prop="wall2" width="120" />
        <el-table-column align="left" label="wall3" prop="wall3" width="120" />
        <el-table-column align="left" label="wall4" prop="wall4" width="120" />
        <el-table-column align="left" label="wall5" prop="wall5" width="120" />
        <el-table-column align="left" label="wall6" prop="wall6" width="120" />
        <el-table-column align="left" label="wall7" prop="wall7" width="120" />
        <el-table-column align="left" label="wall8" prop="wall8" width="120" />
        <el-table-column align="left" label="wall9" prop="wall9" width="120" />
        <el-table-column align="left" label="wall10" prop="wall10" width="120" />

        <!-- <el-table-column align="left" label="状态" prop="status" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.status) }}</template>
        </el-table-column> -->
        <el-table-column align="left" label="wall1" prop="wall1freeze" width="120" />
        <el-table-column align="left" label="wall2" prop="wall2freeze" width="120" />
        <el-table-column align="left" label="wall3" prop="wall3freeze" width="120" />
        <el-table-column align="left" label="wall4" prop="wall4freeze" width="120" />
        <el-table-column align="left" label="wall5" prop="wall5freeze" width="120" />
        <el-table-column align="left" label="wall6" prop="wall6freeze" width="120" />
        <el-table-column align="left" label="wall7" prop="wall7freeze" width="120" />
        <el-table-column align="left" label="wall8" prop="wall8freeze" width="120" />
        <el-table-column align="left" label="wall9" prop="wall9freeze" width="120" />
        <el-table-column align="left" label="wall10" prop="wall10freeze" width="120" />
        <el-table-column align="left" label="推荐人账号" prop="tjuser" width="120" />
        <el-table-column align="left" label="推荐人uid" prop="tjid" width="120" />
        <el-table-column align="left" label="推荐路径" prop="tpath" width="120" />
        <el-table-column align="left" label="邀请码" prop="tgno" width="120" />
        <el-table-column align="left" label="直推代数" prop="ztlevel" width="120" />
        <el-table-column align="left" label="直推人数" prop="ztnum" width="120" />
        <el-table-column align="left" label="团队人数" prop="tdnum" width="120" />
        <el-table-column align="left" label="是否商户" prop="isAgent" width="120" />
        <el-table-column align="left" label="是否通过谷歌验证" prop="isB" width="120" />
        <el-table-column align="left" label="次数" prop="times" width="120" />
        <el-table-column align="left" label="今日直推" prop="todayzt" width="120" />
        <el-table-column align="left" label="团队业绩" prop="tdyj" width="120" />
        <el-table-column align="left" label="个人业绩" prop="gryj" width="120" />
        <el-table-column align="left" label="姓名" prop="realname" width="120" />
        <el-table-column align="left" label="身份证号" prop="idcard" width="120" />
        <el-table-column align="left" label="身份证正面" prop="card1" width="120" />
        <el-table-column align="left" label="身份证反面" prop="card2" width="120" />
        <el-table-column align="left" label="护照" prop="card3" width="120" />
        <el-table-column align="left" label="实名状态:0=未实名,1=已实名,2=待审核" prop="isSm" width="120" />
        <el-table-column align="left" label="S1账户" prop="S1" width="120" />
        <el-table-column align="left" label="S2账户" prop="S2" width="120" />
        <el-table-column align="left" label="累计1" prop="B1total" width="120" />
        <el-table-column align="left" label="累计2" prop="B2total" width="120" />
        <el-table-column align="left" label="累计3" prop="B3total" width="120" />
        <el-table-column align="left" label="昨日1" prop="B1yest" width="120" />
        <el-table-column align="left" label="昨日2" prop="B2yest" width="120" />
        <el-table-column align="left" label="昨日3" prop="B3yest" width="120" />
        <el-table-column align="left" label="算力1" prop="suan" width="120" />
        <el-table-column align="left" label="算力2" prop="suan2" width="120" />
        <el-table-column align="left" label="模拟" prop="isSim" width="120" />
        <el-table-column align="left" label="盈次数" prop="win" width="120" />
        <!-- <el-table-column align="left" label="VIP" prop="sw1" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.sw1) }}</template>
        </el-table-column>
        <el-table-column align="left" label="提现" prop="sw2" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.sw2) }}</template>
        </el-table-column>
        <el-table-column align="left" label="闪兑" prop="sw3" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.sw3) }}</template>
        </el-table-column>
        <el-table-column align="left" label="OTC" prop="sw4" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.sw4) }}</template>
        </el-table-column>
        <el-table-column align="left" label="互转" prop="sw5" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.sw5) }}</template>
        </el-table-column>
        <el-table-column align="left" label="奖励" prop="sw6" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.sw6) }}</template>
        </el-table-column> -->

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
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="150px">
        <!-- <el-form-item label="组别ID:" prop="groupId">
          <el-input v-model.number="formData.groupId" :clearable="true" placeholder="请输入" />
        </el-form-item> -->
        <el-form-item label="用户名:" prop="username">
          <el-input v-model="formData.username" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="昵称:" prop="nickname">
          <el-input v-model="formData.nickname" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <!-- <el-form-item label="密码:" prop="password">
          <el-input v-model="formData.password" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="支付密码:" prop="password2">
          <el-input v-model="formData.password2" :clearable="true" placeholder="请输入" />
        </el-form-item> -->
        <!-- <el-form-item label="谷歌秘钥:" prop="password3">
          <el-input v-model="formData.password3" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="密码盐:" prop="salt">
          <el-input v-model="formData.salt" :clearable="true" placeholder="请输入" />
        </el-form-item> -->
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
          <el-input v-model="formData.level" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <!-- <el-form-item label="性别:" prop="gender">
          <el-switch v-model="formData.gender" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable />
        </el-form-item>
        <el-form-item label="生日:" prop="birthday">
          <el-date-picker v-model="formData.birthday" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
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
          <el-switch v-model="formData.loginfailure" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable />
        </el-form-item>
        <el-form-item label="加入时间:" prop="jointime">
          <el-input v-model.number="formData.jointime" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="会员到期时间:" prop="viptime">
          <el-input v-model.number="formData.viptime" :clearable="true" placeholder="请输入" />
        </el-form-item> -->
        <!-- <el-form-item label="wall1:" prop="wall1">
          <el-input-number v-model="formData.wall1" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="wall2:" prop="wall2">
          <el-input-number v-model="formData.wall2" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="wall3:" prop="wall3">
          <el-input-number v-model="formData.wall3" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="wall4:" prop="wall4">
          <el-input-number v-model="formData.wall4" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="wall5:" prop="wall5">
          <el-input-number v-model="formData.wall5" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="wall6:" prop="wall6">
          <el-input-number v-model="formData.wall6" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="wall7:" prop="wall7">
          <el-input-number v-model="formData.wall7" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="wall8:" prop="wall8">
          <el-input-number v-model="formData.wall8" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="wall9:" prop="wall9">
          <el-input-number v-model="formData.wall9" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="wall10:" prop="wall10">
          <el-input-number v-model="formData.wall10" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item> -->
        <el-form-item label="状态" prop="disabled">
          <el-switch
            v-model="formData.status"
            inline-prompt
            :active-value="1"
            :inactive-value="2"
          />
        </el-form-item>
        <el-form-item label="VIP" prop="disabled">
          <el-switch
            v-model="formData.sw1"
            inline-prompt
            :active-value="1"
            :inactive-value="2"
          />
        </el-form-item>
        <el-form-item label="提现" prop="disabled">
          <el-switch
            v-model="formData.sw2"
            inline-prompt
            :active-value="1"
            :inactive-value="2"
          />
        </el-form-item>
        <el-form-item label="闪兑" prop="disabled">
          <el-switch
            v-model="formData.sw3"
            inline-prompt
            :active-value="1"
            :inactive-value="2"
          />
        </el-form-item>
        <el-form-item label="OTC" prop="disabled">
          <el-switch
            v-model="formData.sw4"
            inline-prompt
            :active-value="1"
            :inactive-value="2"
          />
        </el-form-item>
        <el-form-item label="互转" prop="disabled">
          <el-switch
            v-model="formData.sw5"
            inline-prompt
            :active-value="1"
            :inactive-value="2"
          />
        </el-form-item>
        <el-form-item label="奖励" prop="disabled">
          <el-switch
            v-model="formData.sw6"
            inline-prompt
            :active-value="1"
            :inactive-value="2"
          />
        </el-form-item>

        <!-- <el-form-item label="wall1:" prop="wall1freeze">
          <el-input-number v-model="formData.wall1freeze" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="wall2:" prop="wall2freeze">
          <el-input-number v-model="formData.wall2freeze" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="wall3:" prop="wall3freeze">
          <el-input-number v-model="formData.wall3freeze" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="wall4:" prop="wall4freeze">
          <el-input-number v-model="formData.wall4freeze" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="wall5:" prop="wall5freeze">
          <el-input-number v-model="formData.wall5freeze" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="wall6:" prop="wall6freeze">
          <el-input-number v-model="formData.wall6freeze" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="wall7:" prop="wall7freeze">
          <el-input-number v-model="formData.wall7freeze" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="wall8:" prop="wall8freeze">
          <el-input-number v-model="formData.wall8freeze" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="wall9:" prop="wall9freeze">
          <el-input-number v-model="formData.wall9freeze" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="wall10:" prop="wall10freeze">
          <el-input-number v-model="formData.wall10freeze" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item> -->
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
        <!-- <el-form-item label="直推代数:" prop="ztlevel">
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
        </el-form-item> -->
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
        <!-- <el-form-item label="S1账户:" prop="S1">
          <el-input-number v-model="formData.S1" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="S2账户:" prop="S2">
          <el-input-number v-model="formData.S2" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="累计1:" prop="B1total">
          <el-input-number v-model="formData.B1total" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="累计2:" prop="B2total">
          <el-input-number v-model="formData.B2total" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="累计3:" prop="B3total">
          <el-input-number v-model="formData.B3total" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="昨日1:" prop="B1yest">
          <el-input-number v-model="formData.B1yest" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="昨日2:" prop="B2yest">
          <el-input-number v-model="formData.B2yest" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="昨日3:" prop="B3yest">
          <el-input-number v-model="formData.B3yest" style="width:100%" :precision="2" :clearable="true" />
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
        </el-form-item> -->

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
  name: 'FaUser'
}
</script>

<script setup>
import {
  createFaUser,
  deleteFaUser,
  deleteFaUserByIds,
  updateFaUser,
  findFaUser,
  getFaUserList
} from '@/api/faUser'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { nextTick, ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
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
  level: 1,
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
  status: 1,
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
  sw1: 1,
  sw2: 1,
  sw3: 1,
  sw4: 1,
  sw5: 1,
  sw6: 1,
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
  if (searchInfo.value.level === '') {
    searchInfo.value.level = null
  }
  if (searchInfo.value.gender === '') {
    searchInfo.value.gender = null
  }
  if (searchInfo.value.loginfailure === '') {
    searchInfo.value.loginfailure = null
  }
  if (searchInfo.value.status === '') {
    searchInfo.value.status = null
  }
  if (searchInfo.value.sw1 === '') {
    searchInfo.value.sw1 = null
  }
  if (searchInfo.value.sw2 === '') {
    searchInfo.value.sw2 = null
  }
  if (searchInfo.value.sw3 === '') {
    searchInfo.value.sw3 = null
  }
  if (searchInfo.value.sw4 === '') {
    searchInfo.value.sw4 = null
  }
  if (searchInfo.value.sw5 === '') {
    searchInfo.value.sw5 = null
  }
  if (searchInfo.value.sw6 === '') {
    searchInfo.value.sw6 = null
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
  const table = await getFaUserList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    deleteFaUserFunc(row)
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
  const res = await deleteFaUserByIds({ ids })
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
const updateFaUserFunc = async(row) => {
  const res = await findFaUser({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.refaUser
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteFaUserFunc = async(row) => {
  const res = await deleteFaUser({ ID: row.ID })
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
  }
}
// 弹窗确定
const enterDialog = async() => {
     elFormRef.value?.validate(async(valid) => {
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
         closeDialog()
         getTableData()
       }
     })
}

const switchEnable = async(row) => {
  formData.value = JSON.parse(JSON.stringify(row))
  await nextTick()
  const req = {
    ...formData.value
  }
  const res = await updateFaUser(req)
  if (res.code === 0) {
    ElMessage({ type: 'success', message: `${req.status === 2 ? '禁用' : '启用'}成功` })
    await getTableData()
  }
}
</script>

<style>
</style>
