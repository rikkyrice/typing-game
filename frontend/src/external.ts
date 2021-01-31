// vue-browser-detect-plugin
// ブラウザデフォルトのcssを無効化
import 'normalize.css'

// import vue-toasted
import Vue from 'vue'
import Toasted from 'vue-toasted'

const browserDetect = require('vue-browser-detect-plugin')

const vueSmoothScroll = require('vue-smoothscroll')
Vue.use(vueSmoothScroll)

Vue.use(browserDetect)
Vue.use(Toasted)
