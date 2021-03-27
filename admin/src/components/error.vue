<template>
  <div>
    <transition name="bounce">
      <div class="errorbox" v-show="haserror">
        <p>
          Net Error
          <i class="el-icon-loading"></i>
        </p>
        
      </div>
    </transition>
  </div>
</template>

<script>
export default {
  props: ['haserror'],
  methods:{
    
  }
}
</script>

<style lang="stylus" scoped>
@import './main.styl'
.errorbox 
    height 100%
    width 100%
    position absolute
    display table
    background #fff
    z-index 101
p    
    color: #545454;
    font-size: 20px;
    text-align: center;
    display: table-cell
    vertical-align: middle;
    margin 0
    height 100%
    width 100%

</style>