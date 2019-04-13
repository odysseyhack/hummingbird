// Hummingbird Dashboard
import Vue from 'vue'
import App from './App'
import Vuex from 'vuex'
import 'es6-promise/auto'
import VuexStore from './store/index'
import Router from 'vue-router'
import Buefy from 'buefy'
import 'onsenui/css/onsenui.css'
import 'onsenui/css/onsen-css-components.css'
import VueOnsen from 'vue-onsenui'
import * as VueGoogleMaps from 'vue2-google-maps'
import Home from './views/Home.vue'
import SettingsPage from './views/SettingsPage.vue'
import VueMqtt from 'vue-mqtt'
import Viewer from 'v-viewer'
import vueCustomElement from 'vue-custom-element'

Vue.use(VueGoogleMaps, {
  load: {
    key: 'AIzaSyCFYHQNTzES5wdpD6AA-agVMafL1kg7brg',
    libraries: 'places, drawing',
  }
})

Vue.use(Vuex)
const store = new Vuex.Store(VuexStore)
Vue.use(Router)

Vue.use(Buefy)
Vue.use(VueOnsen)
VueOnsen.platform.select('ios')

const router = new Router({
  routes: [{
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/settings',
      name: 'settings',
      component: SettingsPage
    },
  ]
})

let options = {
  clientId: 'hummingbird-client-' + Math.random().toString(16).substr(2, 4)
}

Vue.use(VueMqtt, 'ws://broker.hivemq.com:8000/mqtt', options)

Vue.config.productionTip = false
Vue.use(Viewer)
Vue.use(vueCustomElement)

new Vue({
  router: router,
  store: store,
  render: h => h(App)
}).$mount('#hummingbird-ui')