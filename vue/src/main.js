import Go from './wasm_exec'

import Vue from 'vue'
import App from './App.vue'

/* eslint no-undef: "off"*/
const go = new Go()
WebAssembly.instantiateStreaming(fetch("calc.wasm"), go.importObject)
  .then(async (result) => {
    // await go.run(result.instance)
    go.run(result.instance)
    // console.log(waAdd(...Array("2", "1")))

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
