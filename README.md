# Practice for using WASM within Vue.js

# Not work Go 1.12 wasm cannot run with Vue.js

## Run
```sh
$ cd wasm
set GOARCH=wasm
set GOOS=js
$ go build -o calc.wasm

$ cp (or copy) calc.wasm ../vue/public/

$ cd ..

$ cd vue
$ npm install
$ npm run serve
```

## Sundries

Finding how to use functions in wasm at vue file?

* https://medium.com/@brockreece/vue-webassembly-1a09e38d0389
* https://stackoverflow.com/questions/43608457/how-to-import-functions-from-different-js-file-in-a-vuewebpackvue-loader-proje


What is wasm-loader?

https://github.com/vuejs/vue-cli/issues/763
* https://github.com/greenpdx/vueopencvjs/blob/master/src/components/HelloWorld.vue
* https://github.com/greenpdx/vueopencvjs/blob/master/src/components/add.es6
