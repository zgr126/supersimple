<template>
  <div class="passbox">
    
    <el-card shadow="hover">
      <el-form ref="form" :model="form" :rules="rules">
        <el-form-item label="Set Admin Password" prop="adminpass">
          <el-input class="input_c" v-model="form.adminpass" placeholder="password" show-password></el-input>
        </el-form-item>
        <el-form-item label="Password Again" prop="adminpassR">
          <el-input class="input_c" v-model="form.adminpassR" placeholder="password" show-password></el-input>
        </el-form-item>
        <ajax-button label="Commit" @click="pushAjax" ref="ajaxbtn"></ajax-button>
      </el-form>
    </el-card>
  </div>
</template>

<script>
import ajaxButton from '../common/ajaxButton.vue';
export default {
  components: { ajaxButton },
  data(){
    var validatePass = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('Please input Password'));
      } else {
        callback();
      }
    };
    var validatePass2 = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('Please input Password'));
      } else if (value !== this.form.adminpass) {
        callback(new Error('Password does not match'));
      } else {
        callback();
      }
    };
    return {
      adminpass: '',
      form: {
        adminpass:'',
        adminpassR: '',
        
      },
      rules: {
        adminpass: [
          { validator: validatePass, trigger: 'blur' }
        ],
        adminpassR: [
          { validator: validatePass2, trigger: 'blur' }
        ],
      }
    }
  },
  methods:{
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        console.log(valid)
        if (valid) {
          console.log("ok")
        } else {
          console.log('error submit!!');
          return false;
        }
      });
    },
    pushAjax(){
      // this.axios.get('password').then(e=>{
      //   console.log(e)
      // })
      let p = this.axios.post('password', {
        passWord: this.form.adminpass
      })
      let v = this.$refs['ajaxbtn'].push(p)
      v.then(e=>{
        console.log('1')
      })
      return p
    }
  }
}
</script>

<style lang="stylus" scoped>
.passbox
  padding 10px
/deep/ .el-card
  max-width 300px
  margin auto
/deep/ .el-form-item
  margin 0 0 8px 0
  /deep/ .el-form-item__label
    line-height 30px
.ok
  margin-top 10px
  width: 100%;
</style>