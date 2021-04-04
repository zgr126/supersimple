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
  beforeCreate(){
    let setTimeout1s = new Promise((resolve, reject)=>{
      setTimeout(e=>{
        resolve()
      },1000)
    })
    let getAdminStatus = this.axios.get('status')
    let getAdminApp = this.axios.get('app')
    this.Utils.promiseX([setTimeout1s,getAdminStatus, getAdminApp]).then(_e=>{
      this.showWelcome = false
      let auth = _e[2]
      let adminS = _e[1]
      let hasCreate = false
      if (!this.IsErrType(adminS)){
        hasCreate = !!adminS.data.CreateTime
      }
      if (this.IsErrType(auth)){
        if (auth.response && auth.response.status == 403){
          this.showWelcome = false
          if (hasCreate){
            this.$router.push('/login')
          }else{
            this.$router.push('/init')
          }
        }
      }else{
        this.$router.push('/index')
      }

      
      // if () 
      // let code = e.response.status
      // let data = e.response.data
      // if (code == 403){
        
      // }
      // this.$router.push('/index')
    })
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
  overflow hidden
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
