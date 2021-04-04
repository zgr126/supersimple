
import ajaxButton from './ajaxButton'
import welcome from './welcome'
import absolutebox from './absolutebox'
const registerComponents={
    install:function(Vue){
        Vue.component('AjaxButton',ajaxButton)
        Vue.component('welcome', welcome)
        Vue.component('absolutebox', absolutebox)
    }
}

export default registerComponents