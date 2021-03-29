<template>
  <transition name="bounce">
    <!-- <div class="errorbox" v-show="haserror">
      <p>
        Net error
        <i class="el-icon-loading"></i>
      </p>
      
    </div> -->
    <div class="logo" v-show="showWelcome">
      <img src="../assets/logo.png">
    </div>
  </transition>
</template>

<script>
export default {
  props: ['showWelcome'],
  methods:{
    
  }
}
</script>

<style lang="stylus" scoped>
@import '../main.styl'
.errorbox 
    height 100%
    width 100%
    position absolute
    display table
    background #fff
    z-index 101
    top 0
p    
    color: #545454;
    font-size: 20px;
    text-align: center;
    display: table-cell
    vertical-align: middle;
    margin 0
    height 100%
    width 100%
.logo
  z-index 5000
  height 100%
  width 100%
  position absolute
  margin 0
  background #fff
  transition: All .5s ease-in-out
  img 
    position: absolute;
    top: calc(50% - 115px);
    left: calc(50% - 200px);
    box-sizing: border-box;
    margin 15px
    height 200px
</style>