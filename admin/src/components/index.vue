<template>
  <div class="contentpage">
      <top></top>
      <div>
        <div class="content_box">
            <el-divider content-position="center">Documents</el-divider>
            <div class="content">
                <el-card class="box-card-add" shadow="hover" @click="addBean">
                    <i class="addicon el-icon-plus"></i>
                </el-card>
            </div>
        </div>
        <div class="content_box">
            <el-divider content-position="center">FileServer</el-divider>
            <div class="content">
                <el-card class="box-card-add" shadow="hover" @click="addFileServer">
                    <i class="addicon el-icon-plus"></i>
                </el-card>
            </div>
        </div>

      </div>
  </div>
</template>

<script>
import top from './top'
export default {
    data(){
        return{
            beans:[],
            httpHeaders: []
        }
    },
    methods:{
        addBean(){},
        addFileServer(){}
    },
    mounted(){
    },
    beforeCreate(){
        this.axios.get('/app').then(e=>{
            this.beans = e.beans
            this.httpHeaders = e.httpHeaders
        })
    },
    components:{
        top
        // bk
    }
}
</script>

<style lang="stylus" scoped>
.contentpage
    height 100%
    width 100%
    overflow hidden
.content_box
    max-width 1000px
    text-align center
    margin auto
    .content
        display: flex;
        flex-flow: wrap;
        >div
            min-width: 250px;
            min-height: 150px;
        .box-card-add
            .addicon
                display: inline-block;
                line-height: 110px;
                font-size: 30px;
/deep/ .el-divider__text
    font-size 20px
</style>