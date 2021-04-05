<template>
  <div id="app">
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <!-- <div class="logo">
      <img src="./assets/logo.png">
    </div> -->
    <welcome :show-welcome="showWelcome"></welcome>
    <div v-show="!showWelcome" class="logo">
      <img src="./assets/logo.png">
    </div>
    <div class="mainpage">
      <router-view></router-view>
    </div>
    
  </div>
</template>

<script>

export default {
  name: 'App',
  data(){
    return {
      showWelcome: false,
    }
  },
  mounted(){
    this.showWelcome = true
  },
  watch:{
    $route(v){
      if (!(v.path == '/' || v.path == '')){
        this.showWelcome = false
      }
    }
  },
  async beforeCreate(){
    this.axios.get('app').then(
      this.$router.push('index')
    )
    // register global getapplicationStatus function
    // this.Global.getApplicationStatus = async ()=>{
    //   // return   0：need init 1: need login  2: all fine
    //   let status = -1
    //   let getAdminStatus = this.axios.get('status')
    //   let getAdminApp = this.axios.get('app')
    //   await this.Utils.promiseX([getAdminStatus, getAdminApp]).then(async _e=>{
    //     let auth = _e[1]
    //     let adminS = _e[0]
    //     let hasCreate = false
    //     if (!this.IsErrType(adminS)){
    //       hasCreate = !!adminS.createTime
    //     }
    //     if (this.IsErrType(auth)){
    //       if (auth.response && auth.response.status == 403){
    //         if (hasCreate){
    //           status= 1
    //         }else{
    //           status= 0
    //         }
    //       }
    //     }else{
    //       status= 2
    //     }
    //   })
    //   return status
    // }
    // let v = await this.Global.getApplicationStatus()
    // setTimeout(e=>{
    //   this.showWelcome = false
    //   switch(v){
    //     case 0:
    //       this.$router.push('/init')
    //       break
    //     case 1:
    //       this.$router.push('/login')
    //       break
    //     case 2:
    //       this.$router.push('/index')
    //   }
    // },500)
  }
}
</script>

<style lang="stylus">
html, body, #app
  height 100%
  width 100%
html, body, #app, .mainpage
  transition none
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  padding: 0
  margin 0
}
body 
  padding 0
  margin 0
.apptop
  position absolute
  z-index: 2;
  left: 50%;
  margin-left: -400px;
.mainpage
  // height 100%
  // width 100%
  // padding 20px
  position relative
  // overflow hidden
* 
  transition: All .2s ease-in-out
  outline:none;
*:hover 
  transition: All .2s ease-in-out
.logo
  display block
  text-align center
  margin-bottom 20px
  img 
    height: 100px;
    margin: 15px;
    // background: white;
  span
    display: block;
    padding: 5px;
</style>
