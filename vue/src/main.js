import Vue from 'vue'
import App from './App.vue'

import goWasm from './wasm_exec'
// const go = require('./wasm_exec')
// const wasm = require('./wasm_exec')

// const go = new wasm()
const go = new goWasm()
// eslint-disable-next-line 
// let mod, inst
let inst
WebAssembly.instantiateStreaming(fetch("calc.wasm"), go.importObject).then(async (result) => {
  // mod = result.module
  inst = result.instance
  await go.run(inst)
})

Vue.prototype.$go = { go };

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')
