<template>
  <div>
      <div class="top">
        <div class="t_l">
          <img src="../assets/logo.png">
        </div>
        <div class="t_m">
          <div v-for="(i, index) of top" :key="index" @click="changeRoute(i)" :class="{active: nowSelect == i.name}">
            <p class="p2">{{i.elabel}}</p>
          </div>
        </div>
        <div class="t_r">
          <span class="right">
          </span>
        </div>
      </div>
  </div>
</template>

<script>
// import mimavue from '../mimavue'
import qs from 'qs'
export default {
  data(){
    var validatePass2 = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请再次输入密码'))
      } else if (value !== this.signForm.mima) {
        callback(new Error('两次输入密码不一致!'))
      } else {
        callback()
      }
    }
    return {
      isZhuceLoading: false,
      iswangji: false,
      wangjiForm:{
        zhanghao: ''
      },
      fileUseLst:[{name: '城市规划'}, {name: '防灾规划'}, {name: '韧性评价'}, {name: '其他'}],
      top:[
        {name: '', label: '首页', elabel: 'HOME'},
        {name: 'doc', label: '文件', elabel: 'DOCUMENT'},
        {name: 'share', label: '共享', elabel: 'SHARE'},
      ],
      username:'',
      nowSelect: '',
      dialogVisible: false,
      loginDialogVisible: false,
      signForm: {
        zhanghao: '',
        zhicheng: '',
        xingming: '',
        danwei: '',
        mima: '',
        shiyanleixing: '',
        chongfumima: '',
        filePurpose: '城市规划',
        needFile: ''
      },
      loginForm:{
        zhanghao: '',
        mima: ''
      },
      fileTypeLst: [],
      rules: {
        zhanghao: [
          { required: true, message: '请输入账号', trigger: 'blur' },
          { type: 'email', message: '请输入正确的邮箱地址', trigger: 'blur'}
        ],
        zhicheng: [
          { required: true, message: '请输入职称', trigger: 'blur' }
        ],
        xingming: [
          { required: true, message: '请输入姓名', trigger: 'blur' }
        ],
        danwei: [
          {required: true, message: '请输入单位', trigger: 'blur' }
        ],
        mima: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 8, max: 30, message: '密码长度为8-30个字符', trigger: 'blur' },
          { required: true, pattern: /^[0-9a-zA-Z]*$/g, message: '只能输入数字和英文', trigger: 'blur' },
        ],
        needFile: [
          { required: true, message: '请填写所需文件', trigger: 'blur' }
        ],
        chongfumima: [
          { required: true, validator: validatePass2, trigger: 'blur' },
          { required: true, pattern: /^[0-9a-zA-Z]*$/g,message: '只能输入数字和英文', trigger: 'blur' },
        ],
        filePurpose: [
          { required: true, message: '请输入文件用途', trigger: 'blur' }
        ],
      }

    }
  },
  methods:{
    wangjipost(){
      this.$refs['wang'].validate((valid) => {
        if (valid) {
          this.axios.post('api-service-base-earthquake/baseUser/getPwdByE', {
            email: this.wangjiForm.zhanghao,
          }).then(e=>{
            if (e.code == 101){return}
            this.$refs['wang'].resetFields()
            this.iswangji = false
            this.$message({
              showClose: true,
              message: '提交成功，请去邮箱进行密码验证',
              type: 'success'
            });
          })
        }
      })
    },
    wangjimima(){
      this.loginDialogVisible = false
      this.iswangji = true
    },
    hasZhanghao(){
      this.login()
      setTimeout(e=>{
        this.$refs['signindialog'].resetFields()
        this.dialogVisible = false
      })
    },
    clickname(){
      this.$router.push('/self')
    },
    signin(){
      this.dialogVisible = true
    },
    login(){
      this.loginDialogVisible = true
      this.loginForm.zhanghao = ''
      this.loginForm.mima = ''
    },
    logout(){
      this.axios({
          url: '/logout',
          method: 'post',
          headers: {
              'Content-Type': 'application/x-www-form-urlencoded'
          }
      }).then(e=>{
        localStorage.removeItem('user')
        document.location.hash = '#/'
        window.location.reload()
      })
    },
    custom(){
      this.loginForm.zhanghao = 'zz'
      this.loginForm.mima = '123456'
      this.Global.username = 'zz'
      this.loginPost()
    },
    loginPost(){
      
      var forms = new FormData()
      var configs = {
          headers:{'Content-Type':'application/x-www-from-urlencoded'}
      }
      forms.append('username', this.loginForm.zhanghao)
      forms.append('password', this.loginForm.mima)
      // this.Global.username = this.loginForm.zhanghao
      let _d = {
        username: this.loginForm.zhanghao,
        password: this.loginForm.mima
      }
      let s = qs.stringify(_d)
      this.axios({
          url: 'authentication/form',
          method: 'post',
          headers: {
              'Content-Type': 'application/x-www-form-urlencoded'
          },
          data: s
      })
      // this.axios.post('authentication/form', s, configs)
      .then(e=>{
        console.log(e)
        this.Global.username = e.data.name
        this.Global.user = e.data
        localStorage.setItem('user', this.loginForm.zhanghao)
        // if (this.Global.username != 'zz'){
        //   this.Global.Cookie.set('user', e.data.name)
        // }
        // this.$emit('loginMethod', e.data.name)
        this.loginDialogVisible = false
        window.location.reload()
      })
    },
    isSignin(){
      this.$refs['signindialog'].validate((valid) => {
        if (valid) {
          this.isZhuceLoading = true
          this.Global.username = this.signForm.zhanghao
          this.axios.post('api-service-base-earthquake/baseUser/register', {
            username: this.signForm.zhanghao,
            nick: this.signForm.xingming,
            companyName: this.signForm.danwei,
            password: this.signForm.mima,
            needFile: this.signForm.needFile,
            filePurpose: this.signForm.filePurpose
          }).then(e=>{
            this.isZhuceLoading = false
            if (e){
              this.$refs['signindialog'].resetFields()
              this.dialogVisible = false
              this.$message({
                showClose: true,
                message: '注册成功',
                type: 'success'
              });
            }
            
            
          }).catch(e=>{
            this.isZhuceLoading = false
          })
        } else {
          // console.log('error submit!!');
          return false;
        }
      })
    },
    changeRoute(i){
      // if (!this.username && (i.name != 'about' && i.name != '')){
      //   this.$message({
      //         showClose: true,
      //         message: '请登录后查看',
      //         type: 'warning'
      //       });
      //   return
      // }
      let path
      if (this.$route.path.slice(1) == i.name)return
      if (!i.name){path = '/'}else{path = i.name}
      this.$router.push(path)
    },
  },
  mounted(){
    this.nowSelect = this.$route.path.slice(1)
    
  },
  watch:{
    '$route.path':function(n,o){
      this.nowSelect = n.slice(1)
    }
  },
  components:{
    // mimavue
  }
}
</script>

<style lang="stylus" scoped>
.top
  width 100%
  height 100px
  max-width 800px
  min-width 800px
  margin: auto;
  // color #fff
  .t_l
    height 100%
    display inline-block
    img 
      height: 60px;
      margin: 15px;
      float left
      // background: white;
    span
      display: block;
      padding: 5px;
  .t_m
    display: inline-block;
    height: 100px;
    float right
    line-height: 100px;
    vertical-align: top;
    >div
      line-height 100px
      
      height 100%
      display inline-block
      margin 0 5px
      &.active
        // background #eaeaea
        color #3d5e80
      p
        margin 5px
        cursor: pointer;
        display inline
        cursor pointer
        &:hover
          color #3d5e80
      .p2
        font-size: 18px;
        font-weight: bold;
  .t_r
    float: right;
    position: absolute;
    top: 0;
    right: 0;
    height:50px;
    margin: 0 20px;
    line-height:50px;
.signin_dialog 
  overflow hidden
  color #1c6cb1
  /deep/ >div
    // position absolute
    // margin-top: 300px!important
    width: 700px!important
    height: 450px;
    text-align left
    // left: 700px;
    // margin: 0;
    border-radius: 5px;
  /deep/ .el-dialog__body
    padding 0px 20px!important
  /deep/ .sign_label
    font-size 1.2rem
  /deep/ .el-select, .el-input {
    width: 200px!important;
  }
.login_dialog 
  
  overflow hidden
  color #1c6cb1
  /deep/ >div
    // position absolute
    // margin-top: 300px!important
    width: 400px!important
    height: 260px;
    // left: 700px;
    // margin: 0;
    border-radius: 5px;
  /deep/ .el-dialog__body
    padding 0px 20px!important
  /deep/ .el-select, .el-input {
    width: 200px!important;
  }
.wangji /deep/ >div
      height 200px
</style>