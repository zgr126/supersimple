<template>
    <div class="contentpage">
        <top @nowSelect="changePage"></top>
        <div class="content" v-if="activePage == 'home'">
            <div class="beans_box">
                <el-divider content-position="center">Documents</el-divider>
                <div class="beans">
                    <el-card class="bean" shadow="hover" v-for="(i, index) of beans" :key="index">
                        <p class="name">{{i.name}}</p>
                        <p class="kv">{{i.kvSize}}</p>
                        <p class="byte">{{i.byteSize | byteFilter}}</p>
                        <p class="time">{{i.time}}</p>
                        <p class="des">{{i.des}}</p>
                        <el-popover
                        popper-class="s_popover"
                        placement="left"
                        width="50">
                            <div style="text-align: right; margin: 0">
                                <el-button size="mini"  type="primary" round @click="testDoc(i)">Test</el-button>
                                <el-button size="mini"  type="danger" round @click="settingDoc(i)">Delete</el-button>
                            </div>
                            <el-button slot="reference" class="setting" icon="el-icon-setting" size="mini" circle ></el-button>
                        </el-popover>
                        
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
        <div class="content" v-if="activePage == 'setting'">
            <div class="beans_box">
                <el-divider content-position="center">Response Header</el-divider>
                <div class="httpheader">
                    <el-form :inline="true" :model="httpform" class="httpform">
                        
                        <el-row :gutter="20"
                            v-for="(header, index) in httpform.headers"
                            :key="index"
                        >
                            <el-col :span="16">
                                <div class="grid-content bg-purple">
                                    <el-form-item
                                    :prop="'headers.' + index + '.k'"
                                    :rules="{
                                        required: true, message: 'required Filed', trigger: 'blur'
                                    }">
                                        <el-input v-model="header.k" placeholder="response key"></el-input>
                                    </el-form-item>
                                </div>
                            </el-col>
                            <el-col :span="8">
                                <div class="grid-content bg-purple">
                                    <el-form-item
                                    :prop="'headers.' + index + '.v'"
                                    :rules="{
                                        required: true, message: 'required Filed', trigger: 'blur'
                                    }">
                                        <el-input v-model="header.v" placeholder="response value"></el-input>
                                    </el-form-item>
                                </div>
                            </el-col>
                        </el-row>
                        <el-button type="primary" @click="onSubmit">查询</el-button>
                    </el-form>
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
            activePage: 'home',
            newDocumentForm: {
                name: '',
                doc: ''
            },
            newFileServerForm: {
                name: '',
                route: '',
                doc: ''
            },
            httpform:{
                headers: [
                    {k: '',v:''}
                ]
            },
            rules: {
                name: [
                    { required: true, message: 'requird field', trigger: 'blur' },
                    { min: 1, max: 100, message: 'length from 1 to 100', trigger: 'blur' }
                ],
                httpheader:[
                    { required: true, message: 'requird field', trigger: 'blur' },
                ]
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
        settingDoc(item){
            let _item = JSON.parse(JSON.stringify(item))
            _item.status = 2 //beanStatusDisable
            this.axios.post('setBean',item).then(e=>{
                console.log(e)
            })
        },
        testDoc(item){

        },
        changePage(v){
            this.activePage = v
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
                
                min-width: 250px;
                min-height: 150px;
                margin: 10px;
                width: calc(33% - 20px);
            .box-card-add
                .addicon
                    display: inline-block;
                    line-height: 110px;
                    font-size: 30px;
.bean
    font-size 20px
    position relative
    /deep/ .el-card__body
        padding 10px
    p
        margin 0
    
    .byte, .time, .name
        font-size: 14px;
        text-align: left;
        padding-left: 30px;
        color #b1b1b1
    .name
        font-size 40px
        color #000
    .kv
        position: absolute;
        right: 10px;
        top: 36px;
        font-size: 70px;
        color: #656565;
    .setting
        position: absolute;
        right: 10px;
        top: 10px;
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
        width 95%
        max-width 400px
        /deep/ .el-dialog__body
            padding-top 0
</style>