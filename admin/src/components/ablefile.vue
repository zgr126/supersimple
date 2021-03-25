<template>
  <div class="dd">
      <div>您还未激活上传/下载文件的权限，如需激活，请下载并填写“<span @click="down">申请模板</span>”后，发送至管理员邮箱（aa@bb.com） </div>
  </div>
</template>

<script>
export default {
    data(){
        return{

        }
    },
    methods:{
        down(){
            // let data = e.data
            const ele = document.createElement('a');
            ele.setAttribute('href', this.axios.defaults.baseURL + '/dizhen/shenqing.docx');
            //this.$options.filters['filterUrl']是调用全局过滤器,filterUrl是你自己项目main.js里面定义的过滤器
            // ele.setAttribute('download','');
            ele.target = "_blank"
            ele.download = '申请模板'
            ele.style.display = 'none';
            document.body.appendChild(ele);
            ele.click();
            document.body.removeChild(ele);
        }
    }
}
</script>

<style lang="stylus" scoped>
.dd
    max-width 1410px
    text-align right
    color: #860101;
.dd span
    cursor: pointer;
    color: #238fff;
    margin: 0 10px;
    text-decoration: underline;
</style>