import Go from './wasm_exec'

import Vue from 'vue'
import App from './App.vue'

/* eslint no-undef: "off"*/
const go = new Go()
WebAssembly.instantiateStreaming(
  fetch("calc.wasm"), go.importObject
).then((result) => {
  go.run(result.instance)

  Vue.prototype.$go = {
    add: waAdd,
    sub: waSub,
    multi: waMulti,
    divi: waDivi,
  }
})

new Vue({
  render: h => h(App)
}).$mount('#app')
