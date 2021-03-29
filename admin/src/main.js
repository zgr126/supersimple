import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import routes from './router.js'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import global from './global'
import filter from './filter'
import "babel-polyfill"
import Es6Promise from 'es6-promise'
Es6Promise.polyfill()
import 'fetch-detector';
import 'fetch-ie8';
import registerComponents from './common/registerComponents'
import axios from 'axios'
import VueAxios from 'vue-axios'
import VueWorker from 'vue-worker'

Vue.use(VueAxios, axios)
Vue.use(global.install)
Vue.use(filter)
Vue.use(VueWorker)

// var env = process.env.NODE_ENV
Vue.config.productionTip = false
axios.defaults.withCredentials = true;
axios.defaults.timeout =  5000;
axios.interceptors.request.use( (request)=> {
  return request
})
axios.interceptors.response.use(function (response) {
  let data = response.data
  if (data.code === 403) {
    // console.log(data.status)
    localStorage.removeItem('user')
    document.cookie = ''
    document.location.hash = '#/'
    window.location.reload()
  }
  if (response.config.method == 'options') {
    return '1'
  }
  if (data && data.code === 200) {            
    return data  
  } else {            
    ElementUI.Message({
      type: 'error',
      message: data.msg || 'some error'
    })
    return Promise.reject(response);        
  }
  
}, function (error) {
  // ElementUI.Message({
  //   type: 'error',
  //   message: error.message
  // })
})

ElementUI.Dialog.props.closeOnClickModal.default = false
ElementUI.Dialog.props.closeOnPressEscape.default = false
ElementUI.Dialog.props.destroyOnClose.default = true
const originalPush = VueRouter.prototype.push;
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
};

axios.defaults.baseURL = process.env.VUE_APP_BASIC_URL

Vue.use(VueRouter)
Vue.use(ElementUI)
Vue.use(registerComponents)

const router = new VueRouter({
  routes
})
router.beforeEach((to, from, next) => {
  document.title = 'Super-Simple'
  next()
})

new Vue({
  render: h => h(App),
  router
}).$mount('#app')
