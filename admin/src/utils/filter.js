function formatDate (timeStamp, fmt) {
  fmt = fmt || 'yyyy-MM-dd hh:mm:ss'
  var date = new Date(timeStamp)
  var o = {
    'M+': date.getMonth() + 1, // 月份
    'd+': date.getDate(), // 日
    'h+': date.getHours(), // 小时
    'm+': date.getMinutes(), // 分
    's+': date.getSeconds(), // 秒
    'q+': Math.floor((date.getMonth() + 3) / 3), // 季度
    'S': date.getMilliseconds() // 毫秒
  }
  if (/(y+)/.test(fmt)) fmt = fmt.replace(RegExp.$1, (date.getFullYear() + '').substr(4 - RegExp.$1.length))
  for (var k in o) { if (new RegExp('(' + k + ')').test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length === 1) ? (o[k]) : (('00' + o[k]).substr(('' + o[k]).length))) }
  return fmt
}

function formatMin (value, fmt) {
  let minAndSecond = value.toString().split('.')
  let minStr = minAndSecond[0] ? minAndSecond[0] + '分' : ''
  let secondStr = minAndSecond[1] ? minAndSecond[1] + '秒' : ''
  return minStr + secondStr
}

function formatMinToHoursMin (value) {
  let h = Math.floor(value / 60)
  let m = value % 60
  return h === 0 ? m + '分钟' : h + '小时' + m + '分钟'
}

function formatSecondeToHoursMin (value) {
  let m = Math.floor(value / 60)
  let h = Math.floor(m / 60)
  return h === 0 ? (m === 0 ? '1分钟' : m + '分钟') : h + '小时' + m + '分钟'
}
const Dayago = (v, v2)=>{
  if (!(v && v2))return ''
  const sub = v2-v
  const d = sub/1000/60/60/24
  return Math.floor(d)
}

export default {
  install (Vue, options) {
    Vue.filter('Time', function (value, format) {
      if (!value) return ''
      return formatDate(value, format)
    })
    Vue.filter('Date', function (value) {
      if (!value) return ''
      return formatDate(value, 'yyyy-MM-dd')
    })
    Vue.filter('Min', function (value, format) {
      if (!value) return '0分'
      return formatMin(value, format)
    })
    Vue.filter('MinToHoursMin', function (value) {
      if (!value) return '0分钟'
      return formatMinToHoursMin(value)
    })
    Vue.filter('formatSecondeToHoursMin', function (value) {
      if (!value) return '0分钟'
      return formatSecondeToHoursMin(value)
    })
    Vue.filter('Dayago', function (value,v, v2) {
      if (!value) return '0天'
      return Dayago(v, v2)
    })
  }
}
