<template>

  <div class="page-component">
    <section class="content element-doc content">
      <h2 id="an-zhuang">
        快速上手</h2>
      <h3 id="npm-an-zhuang">
        App前端对接， url 携带参数直接跳转到该链接即可</h3>

      <pre>
          <code class="language-shell hljs">http://game.tellwe.xyz/#/?id={{ userStore.userInfo.ID }}&order_sn=订单号
</code></pre>
      <p>订单号推荐使用 base64加密 app内用户id,方便回调时解析充值</p>
      <h3 id="cdn">后端对接 充值回调接口</h3>
      <p>POST 请求</p>
      <pre>
          <code class="language-shell hljs">参数
order_sn 订单号
order_price 订单金额
contract 合约地址
from 转账地址
to 收款地址
sign 签名验证 加密方法 md5(order_sn+ order_price + from + to + 商户id)

收到请返回 success
</code></pre>
      <!-- <div class="tip">
        <p>我们建议使用 CDN 引入 Element 的用户在链接地址上锁定版本，以免将来 Element 升级时受到非兼容性更新的影响。锁定版本的方法请查看 <a
          href="https://unpkg.com"
        >unpkg.com</a>。</p>
      </div> -->

    </section>
    <!-- <div class="footer-nav"><span class="footer-nav-link footer-nav-left"><i class="el-icon-arrow-left" />
      更新日志
    </span><span class="footer-nav-link footer-nav-right">
      快速上手
      <i class="el-icon-arrow-right" /></span></div>
  </div> -->
  </div></template>

<script>
import { useUserStore } from '@/pinia/modules/user'
// const userStore = useUserStore()
// console.log(userStore.userInfo.ID)
export default {
  components: {},
  props: {},
  data() {
    return {
      formData: {},
      rules: {},
      userStore: {},
    }
  },
  computed: {},
  watch: {},
  created() {
    this.userStore = useUserStore()
  },
  mounted() {},
  methods: {
    submitForm() {
      this.$refs['vForm'].validate(valid => {
        if (!valid) return
        // TODO: 提交表单
      })
    },
    resetForm() {
      this.$refs['vForm'].resetFields()
    }
  }
}
</script>

<style>
	.page-component {
		background: #fff;
	}

	.page-component .content>h3 {
		margin: 55px 0 20px;
		font-weight: 400;
		color: #1f2f3d;
	}

	.page-container p {
		font-size: 14px;
		color: #5e6d82;
		line-height: 1.5em;
	}

	.hljs {
		line-height: 1.8;
		font-family: Menlo, Monaco, Consolas, Courier, monospace;
		font-size: 12px;
		padding: 18px 24px;
		background-color: #fafafa;
		border: 1px solid #eaeefb;
		margin-bottom: 25px;
		border-radius: 4px;
		-webkit-font-smoothing: auto;
	}

	pre {
		display: block;
		font-family: monospace;
		white-space: pre;
		margin: 1em 0px;
	}

	.page-container .tip {
		padding: 8px 16px;
		background-color: #ecf8ff;
		border-radius: 4px;
		border-left: 5px solid #50bfff;
		margin: 20px 0;
	}
</style>
