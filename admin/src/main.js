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
  //   let userName = global.globalVars.username
  //   let url = request.url
  //   let match = url.match(/\?/)
  //   if (match) {
  //     let paramsStr = match.input.slice(match.index)
  //     let matchUserName = paramsStr.match('username')
  //     if (!matchUserName) {
  //       request.url = request.url + '&username=' + userName
  //     }
  //   } else {
  //     request.url = request.url + '?username=' + userName
  //   }
  // if (request.method === 'post'){
  //   request.data.username = global.globalVars.username
  // }
  // console.log(request)
  return request
})
axios.interceptors.response.use(function (response) {
  // 判断code
  let data = response.data
  // 登陆超时，重新登陆
  if (data.code === 403) {
    // console.log(data.status)
    localStorage.removeItem('user')
    document.cookie = ''
    document.location.hash = '#/'
    window.location.reload()
  }
  if (response.config.method == 'options') {
    return '1'
  }else{
    if (response.headers['content-type'] !== '' && response.headers['content-type'] !== "text/plain" && response.headers['content-type'].indexOf('octet-stream') === -1) {
      if (data.code == 500 || !data.code){
        console.log(data)
        return 
      }else if (!(data.code == 100 || data.code == 200)) {
        ElementUI.Message({
          type: 'error',
          message: data.msg|| data.message || '出错了'
        })
        // throw new Error(data.msg || '出错了')
      }
    }
    return response.data
  }
  
}, function (error) {
  ElementUI.Message({
    type: 'error',
    message: error.message
  })
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

const router = new VueRouter({
  routes // (缩写) 相当于 routes: routes
})
router.beforeEach((to, from, next) => {
  /* 路由发生变化修改页面title */
  document.title = '虚拟实验室管理平台'
  next()
})
new Vue({
  render: h => h(App),
  router
}).$mount('#app')
