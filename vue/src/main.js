import Vue from 'vue'
import App from './App.vue'

import wasm from './wasm_exec'

const go = new wasm()
// const go = new Go()
let inst
WebAssembly.instantiateStreaming(fetch("calc.wasm"), go.importObject)
  .then(async (result) => {
    inst = result.instance
    await go.run(inst)
  })

// Vue.prototype.$go = { go }; // How to make it??? T^T

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')
