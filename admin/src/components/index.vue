<template>
  <div class="contentpage">
      <!-- <loadzip></loadzip> -->
  </div>
</template>

<script>

export default {
    data(){
        return{
            
        }
    },
    methods:{
    },
    mounted(){
        this.axios.get('state').then(e=>{
            console.log(e)
        })
    },
    components:{
        
        // bk
    }
}
</script>

<style lang="stylus" scoped>
.contentpage
    height 100%
    width 100%
    overflow hidden
</style>