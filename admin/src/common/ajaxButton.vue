<template>
  <div>
    <el-button class="ok" size="mini" type="primary" @click="push">{{label}}</el-button>
  </div>
</template>

<script>
function isPromise(obj) {
  return !!obj  
    && (typeof obj === 'object' || typeof obj === 'function')
    && typeof obj.then === 'function';
}
export default {
    data(){
        return{
            status: 0, //0: input 1: loading 2: success 3: failed
        }
    },
    props: ['label', 'promise','success', 'error'],
    methods:{
        push(){
            let _p = this.promise()
            // if _p not typeof promise, return error hook
            if (!isPromise(_p)){
                this.error()
                return 
            }
            this.status = 1
            _p.then(e=>{
                this.success(e)
                this.status = 2
            }).catch(e=>{
                this.error(e)
                this.status = 3
            })
        },
        reset(){
            this.status = 0
        }
    },
    mounted(){
    }
}
</script>

<style lang="stylus" scoped>
</style>