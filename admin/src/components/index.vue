<template>
  <div class="contentpage">
      <top></top>
      <div class="content">
        <div class="beans_box">
            <el-divider content-position="center">Documents</el-divider>
            <div class="beans">
                <el-card class="bean" shadow="hover" v-for="(i, index) of beans" :key="index">
                    <span>{{i.name}}</span>
                </el-card>
                <el-card class="bean box-card-add" shadow="hover" @click.native="showDialogDocument = true">
                    <i class="addicon el-icon-plus"></i>
                </el-card>
            </div>
        </div>
        <div class="beans_box">
            <el-divider content-position="center">FileServer</el-divider>
            <div class="beans">
                <el-card class="bean box-card-add" shadow="hover" @click.native="showDialogFileServer = true">
                    <i class="addicon el-icon-plus"></i>
                </el-card>
            </div>
        </div>

      </div>
        <el-dialog
        title="Add Document"
        :visible.sync="showDialogDocument"
        class="common-dialog">
            <el-form ref="newDocumentForm" :model="newDocumentForm" :rules="rules">
                <el-form-item label="name" prop="name">
                    <el-input class="input_c" v-model="newDocumentForm.name" show-password></el-input>
                </el-form-item>
                <el-form-item label="doc" prop="doc">
                    <el-input class="input_c" v-model="newDocumentForm.doc" show-password></el-input>
                </el-form-item>
                <ajax-button label="Commit" @click.native="submitForm('newDocumentForm')" ref="ajaxbtn"></ajax-button>
            </el-form>
        </el-dialog>
        <el-dialog
        title="Add FileServer"
        :visible.sync="showDialogFileServer"
        class="dialog">
            <span>这是一段信息</span>
            <span slot="footer" class="dialog-footer">
                <el-button @click="dialogVisible = false">取 消</el-button>
            </span>
        </el-dialog>
  </div>
</template>

<script>
import top from './top'
export default {
    data(){
        return{
            beans:[],
            httpHeaders: [],
            showDialogDocument: false,
            showDialogFileServer: false,
            newDocumentForm: {
                name: '',
                doc: ''
            },
            newFileServerForm: {
                name: '',
                route: '',
                doc: ''
            },
            rules: {
                name: [
                    { required: true, message: 'requird field', trigger: 'blur' },
                    { min: 1, max: 100, message: 'length from 1 to 100', trigger: 'blur' }
                ],
            }
        }
    },
    methods:{
        submitForm(formName){
            this.$refs[formName].validate((valid) => {
                if (valid) {
                    this.pushAjax(formName)
                } else {
                    return false;
                }
            });
        },
        pushAjax(formName){
            switch(formName){
                case 'newDocumentForm':
                    this.addDocument(formName)
                    break
                case 'newFile':
                    break
            }
        },
        addDocument(formName){
            this.axios.post('addBean', {
                name: this.newDocumentForm.name,
                doc: this.newDocumentForm.doc
            }).then(e=>{
                this.$refs[formName].resetFields()
                this.showDialogDocument = false
                this.$message.success('Add Document success')
                this.refrushPage()
            })
        },
        addFileServer(){
            
        },
        refrushPage(){
            this.axios.get('/app').then(e=>{
                this.beans = e.beans
                this.httpHeaders = e.httpHeaders
            })
        }
    },
    mounted(){
        this.refrushPage()
    },
    beforeCreate(){
        
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
.content
    padding 20px 0
    .beans_box
        max-width 1000px
        text-align center
        margin auto
        .beans
            
            display: flex;
            flex-flow: wrap;
            /deep/ .el-card
                cursor pointer
                min-width: 250px;
                min-height: 150px;
                margin: 10px;
                width: calc(33% - 20px);
            .box-card-add
                .addicon
                    display: inline-block;
                    line-height: 110px;
                    font-size: 30px;
@media (max-width: 860px) and (min-width: 570px) {
    /deep/ .el-card{
        width calc(50% - 30px)!important
    }
}
@media (max-width: 570px)  {
    /deep/ .el-card{
        width calc(100% - 20px)!important
    }
}
/deep/ .el-divider__text
    font-size 20px
.common-dialog
    /deep/ .el-dialog
        width 500px
        /deep/ .el-dialog__body
            padding-top 0
</style>