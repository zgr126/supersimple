<template>
  <div class="passbox">
    
    <el-card shadow="hover">
      <label class="cardlabel">Set Admin Password</label>
      <el-form ref="form" :model="form" :rules="rules">
        <el-form-item label="password" prop="adminpass">
          <el-input class="input_c" v-model="form.adminpass" placeholder="password" show-password></el-input>
        </el-form-item>
        <el-form-item label="Password Again" prop="adminpassR">
          <el-input class="input_c" v-model="form.adminpassR" placeholder="password" show-password></el-input>
        </el-form-item>
        <ajax-button label="Commit" @click="submitForm('form')" ref="ajaxbtn"></ajax-button>
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
          { validator: validatePass, trigger: 'blur' },
          { min: 8, max: 100, message: 'length from 8 to 100', trigger: 'blur' }
        ],
        adminpassR: [
          { validator: validatePass2, trigger: 'blur' },
          { min: 8, max: 100, message: 'length from 8 to 100', trigger: 'blur' }
        ],
      }
    }
  },
  methods:{
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        console.log(valid)
        if (valid) {
          this.pushAjax()
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
      let p = this.axios.post('setPassword', {
        password: this.form.adminpass
      })
      let v = this.$refs['ajaxbtn'].push(p)
      v.then(e=>{
        this.$router.replace('/index')
      })
      return p
    }
  }
}
</script>

<style lang="stylus" scoped>
@import '../main.styl'
.passbox
  padding-top: 50px;

/deep/ .el-card
  max-width 300px
  margin auto
/deep/ .el-form-item
  // margin 0 0 8px 0
  /deep/ .el-form-item__label
    line-height 30px
/deep/ .el-card
  overflow inherit
  /deep/ .el-card__body
    position relative
.ok
  margin-top 10px
  width: 100%;
</style>