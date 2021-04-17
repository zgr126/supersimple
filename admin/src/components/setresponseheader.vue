<template>
    <div>
        <div class="httpheader">
            <el-form :inline="true" :model="httpform" class="httpform">
                
                <el-row :gutter="20"
                    v-for="(header, index) in httpform.headers"
                    :key="index"
                >
                    <el-col :span="8">
                        <div class="grid-content bg-purple">
                            <el-form-item
                            :prop="'headers.' + index + '.key'"
                            :rules="{
                                required: true, message: 'required Filed', trigger: 'blur'
                            }">
                                <el-input v-model="header.key" placeholder="response key"></el-input>
                            </el-form-item>
                        </div>
                    </el-col>
                    <el-col :span="16">
                        <div class="grid-content bg-purple">
                            <el-form-item
                            :prop="'headers.' + index + '.value'"
                            :rules="{
                                required: true, message: 'required Filed', trigger: 'blur'
                            }">
                                <el-input v-model="header.value" placeholder="response value"></el-input>
                            </el-form-item>
                        </div>
                    </el-col>
                </el-row>
                <el-button class="save_setting_btn" size="mini" @click="saveSetting">Save</el-button>
            </el-form>
        </div>
  </div>
</template>

<script>
export default {
    data(){
        return{
            httpform:{

            }
        }
    },
    props: ['headerlist'],
    beforeCreate(){
        
    },
    methods:{
        saveSetting(){
            
        }
    }
}
</script>

<style lang="stylus" scoped>
// setting
.httpheader
    padding 0 10px
    /deep/ .el-form--inline .el-form-item
        width 100%
        .el-form-item__content
            width 100%
.save_setting_btn
    width: 200px;
    border: solid 1px;
</style>