<template>
  <div>
    <el-button class="ok" size="mini" type="primary" @click="commit">{{label}}</el-button>
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
            sourceLabel: ''
        }
    },
    props: ['label', 'promise'],
    methods:{
        commit(){
            this.$emit('click')
        },
        push(_p){
            // if _p not typeof promise, return error hook
            if (!isPromise(_p)){
                return _p
            }
            this.status = 1
            return _p.then(async e=>{
                console.log('ajaxok')
                await this.successTimeout()
                // this.success(e)
                this.status = 2
            }).catch(e=>{
                this.status = 3
            })
        },
        async successTimeout(){
            await setTimeout(e=>{
                
            }, 2000)
        },
        reset(){
            this.status = 0
        }
    },
    mounted(){
        this.sourceLabel = this.label
    }
}
</script>

<style lang="stylus" scoped>
.ok
    width 100%
</style>