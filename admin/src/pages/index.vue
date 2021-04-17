<template>
    <div class="contentpage">
        <top @nowSelect="changePage"></top>
        <!-- home -->
        <div class="content" v-if="activePage == 'home'">
            <div class="beans_box">
                <el-divider content-position="center">
                    Documents
                    <el-button class="document_filter" :class="{active: activeDocument =='alive'}" icon="el-icon-video-play" size="mini" circle @click="activeDocument = 'alive'"></el-button>
                    <el-button class="document_filter" :class="{active: activeDocument =='disable'}" icon="el-icon-delete" size="mini" circle @click="activeDocument = 'disable'"></el-button>
                </el-divider>
                <div class="beans">
                    <el-card class="bean" shadow="hover" v-for="(i, index) of beans" :key="index" @click.native="beanDetailsDialog = true;activeBean = i">
                        <p class="name">{{i.name}}</p>
                        <p class="kv">{{i.kvSize}}</p>
                        <p class="byte">{{i.byteSize | byteFilter}}</p>
                        <!-- <p class="time">{{i.time}}</p> -->
                        <p class="des">{{i.des}}</p>
                        
                    </el-card>
                    <el-card class="bean box-card-add" shadow="hover" @click.native="showDialogDocument = true; newDocumentForm.isFileServer = false">
                        <i class="addicon el-icon-plus"></i>
                    </el-card>
                </div>
            </div>
            <div class="beans_box">
                <el-divider content-position="center">
                    FileServer
                    <el-button class="document_filter" :class="{active: activeFileserver =='alive'}" icon="el-icon-video-play" size="mini" circle @click="activeFileserver = 'alive'"></el-button>
                    <el-button class="document_filter" :class="{active: activeFileserver =='disable'}" icon="el-icon-delete" size="mini" circle @click="activeFileserver = 'disable'"></el-button>
                </el-divider>
                <div class="beans">
                    <el-card class="bean" shadow="hover" v-for="(i, index) of fileservers" :key="index">
                        <p class="name">{{i.name}}</p>
                        <p class="kv">{{i.kvSize}}</p>
                        <p class="byte">{{i.byteSize | byteFilter}}</p>
                        <p class="time">{{i.time}}</p>
                        <p class="des">{{i.des}}</p>
                        <el-popover
                        popper-class="s_popover"
                        placement="left"
                        width="160">
                            <div style="text-align: right; margin: 0">
                                <el-button size="mini" v-if="activeDocument == 'alive'" type="primary" round @click="testDoc(i)">Test</el-button>
                                <el-button size="mini" v-if="activeDocument == 'disable'" type="primary" round @click="settingDoc(i, {status:1})">Recover</el-button>
                                <el-button size="mini"  type="danger" round @click="settingDoc(i, {status:2})">Delete</el-button>
                            </div>
                            <el-button slot="reference" class="setting" icon="el-icon-setting" size="mini" circle ></el-button>
                        </el-popover>
                        
                    </el-card>
                    <el-card class="bean box-card-add" shadow="hover" @click.native="showDialogDocument = true; newDocumentForm.isFileServer = true">
                        <i class="addicon el-icon-plus"></i>
                    </el-card>
                </div>
            </div>
            <div class="bottombox"></div>
        </div>
        <!-- setting -->
        <div class="content" v-if="activePage == 'setting'">
            <div class="beans_box">
                <el-divider content-position="center">Common Response Header</el-divider>
                <rheaders></rheaders>
            </div>
        </div>
        <el-dialog
        :title="newDocumentForm.isFileServer? 'Add FileServer':'Add Document'"
        :visible.sync="showDialogDocument"
        class="common-dialog">
            <el-form ref="newDocumentForm" :model="newDocumentForm" :rules="rules">
                <el-form-item label="name" prop="name">
                    <el-input class="input_c" v-model="newDocumentForm.name" ></el-input>
                </el-form-item>
                <el-form-item label="doc" prop="doc">
                    <el-input class="input_c" v-model="newDocumentForm.doc" ></el-input>
                </el-form-item>
                <ajax-button label="Commit" @click.native="submitForm('newDocumentForm')" ref="ajaxbtn"></ajax-button>
            </el-form>
        </el-dialog>
        <!-- bean details -->
        <el-dialog
        title=""
        :visible.sync="beanDetailsDialog"
        class="common-dialog">
            <p class="name">{{activeBean.name}}</p>
            <p class="kv">{{activeBean.kvSize}}</p>
            <p class="byte">{{activeBean.byteSize | byteFilter}}</p>
            <p class="time">{{activeBean.time}}</p>
            <p class="des">{{activeBean.des}}</p>
            <el-divider content-position="center">Set Response Header</el-divider>
            <span slot="footer" class="dialog-footer">
                <el-button @click="beanDetailsDialog = false">取 消</el-button>
            </span>
        </el-dialog>
    </div>
</template>

<script>
import top from './top'
import rheaders from '../components/setresponseheader'
export default {
    data(){
        return{
            sourceBeans:[],
            beans:[],
            fileservers: [],
            httpHeaders: [],
            showDialogDocument: false,
            showDialogFileServer: false,
            beanDetailsDialog: false,
            activeBean: {},
            
            activePage: 'home',
            activeDocument: 'alive',
            activeFileserver: 'alive',
            newDocumentForm: {
                name: '',
                doc: '',
                isFileServer: false
            },
            httpform:{
                headers: [
                    {key: '',value:''}
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
        saveSetting(){
            let headerList = []
            this.httpform.headers.map(e=>{
                headerList.push(e)
            })
            this.axios.post('setting', this.httpform).then(e=>{
                console.log(e)
            })
        },
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
            this.axios.post('addBean', {
                name: this.newDocumentForm.name,
                des: this.newDocumentForm.doc,
                isFileServer: this.newDocumentForm.isFileServer
            }).then(e=>{
                this.$refs[formName].resetFields()
                this.showDialogDocument = false
                this.$message.success('Add Document success')
                this.refrushPage()
            })
        },
        addFileServer(){
            
        },
        settingDoc(item, options){
            let _item = JSON.parse(JSON.stringify(item))
            Object.assign(_item, options)
            // if delete doc
            if (options.status == 2 && item.status == 2){
                _item.status = 3
            }
            
            this.axios.post('setBean',_item).then(e=>{
                this.refrushPage()
            })
        },
        testDoc(item){

        },
        showBeans(){
            this.beans = []
            this.fileservers = []
            this.sourceBeans.map(e=>{
                if ((this.activeDocument == 'alive' && e.status == 1) || (this.activeDocument == 'disable' && e.status == 2))
                {
                    if(!e.isFileServer)this.beans.push(e)
                }
                if ((this.activeFileserver == 'alive' && e.status == 1) || (this.activeFileserver == 'disable' && e.status == 2))
                {
                    if(e.isFileServer)this.fileservers.push(e)
                }
                
            })
        },
        changePage(v){
            this.activePage = v
        },
        refrushPage(){
            this.axios.get('/app').then(e=>{
                this.sourceBeans = e.beans
                this.httpHeaders = e.httpHeaders
                this.showBeans()
            })
        }
    },
    watch:{
        activeDocument(){this.showBeans()},
        activeFileserver(){this.showBeans()}
    },
    mounted(){
        this.refrushPage()
    },
    beforeCreate(){
        
    },
    components:{
        top,rheaders
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
.bean
    font-size 20px
    position relative
    /deep/ .el-card__body
        padding 10px
    p
        margin 0
    
    .byte, .time, .name, .des
        font-size: 14px;
        text-align: left;
        padding-left: 30px;
        color #b1b1b1
    .name
        font-size 40px
        color #000
    .des
        padding-top 30px
        overflow: hidden;
        white-space: nowrap;
        text-overflow: ellipsis;
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
.document_filter
    &.active
        border: solid 1px #808080;
        box-shadow: 0px 0px 3px #ababab;
.bottombox
    height 100px
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
        width 80%
        min-width: 400px;
        max-width: 800px;
        /deep/ .el-dialog__body
            padding-top 0
</style>