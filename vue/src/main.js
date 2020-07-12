import Go from './wasm_exec'

const go = new Go()
WebAssembly.instantiateStreaming(fetch("calc.wasm"), go.importObject)
  .then(async (result) => {
    await go.run(result.instance)
  })

import Vue from 'vue'
import App from './App.vue'

new Vue({
  render: h => h(App)
}).$mount('#app')
