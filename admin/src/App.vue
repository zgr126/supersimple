<template>
  <div id="app">
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <div class="logo">
      <img src="./assets/logo.png">
    </div>
    <error-page :haserror="netErr"></error-page>
    
    <div class="mainpage">
      <router-view></router-view>
    </div>
    
  </div>
</template>

<script>
import errorPage from './components/error'

export default {
  name: 'App',
  data(){
    return {
      netErr: false,
    }
  },
  mounted(){
    
  },
  beforeCreate(){
    this.axios.get('status').then(e=>{
      if (e.data == null){
        this.$router.push('/init')
      }else{
        this.$router.push('/index')
      }
    }).catch(e=>{
      this.netErr = true
    })
  },
  components:{
    errorPage, 
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
  padding 20px
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
