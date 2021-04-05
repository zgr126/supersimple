<template>
  <div>
      <div class="top">
        
        <div class="t_m">
          <div v-for="(i, index) of top" :key="index" @click="topBtnClick(i)" :class="{active: nowSelect == i.name}">
            <p class="p2">{{i.label}}</p>
          </div>
        </div>
        <div class="t_r">
          <span class="right">
          </span>
        </div>
        <div></div>
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
      top:[
        {name: 'home', label: 'Home'},
        {name: 'setting', label: 'Setting'},
        {name: 'logout', label: 'Logout'},
      ],
      nowSelect: ''
    }
  },
  methods:{
    topBtnClick(i){
      if (i.name == 'logout'){
        this.axios.post('/logout').then(e=>{
          this.$router.replace('/')
        }).catch()
        
      }
      this.nowSelect = i.name
    }
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
  margin: auto;
  margin-bottom 20px
  // color #fff
  
  .t_m
    display: block;
    text-align center
    >div
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