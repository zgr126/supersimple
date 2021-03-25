<template>
  <div id="app">
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <top class="apptop"></top>
    <div class="mainpage">
      <router-view></router-view>
    </div>
    
  </div>
</template>

<script>
// import HelloWorld from './components/HelloWorld.vue'
////////////////////////////////////////////////////////////////
// 改动源码 THREE orbitcontrol handleMouseWheel 
// 搜索 evt.initMouseEvent("wheel"
////////////////////////////////////////////////////////////////
import top from './components/top'
export default {
  name: 'App',
  mounted(){
  },
  components: {
    // HelloWorld,
    top
  }
}
</script>

<style lang="stylus">
html, body, #app, .mainpage
  height 100%
  width 100%
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
  height 100%
  width 100%
  position relative
  overflow hidden
* 
  transition: All .2s ease-in-out
  outline:none;
*:hover 
  transition: All .2s ease-in-out
</style>
