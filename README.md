# Go 1.12 WASM Vue.js practice
A `WASM Calculator` practice within `Vue.js`

## Need
* `Vue-Cli3`
* `Go 1.12`

## Run
```sh
$ cd wasm
set GOARCH=wasm
set GOOS=js
$ go build -o calc.wasm

$ cp (or copy) calc.wasm ../vue/public/ (or ../vue/assets)

$ cd ..

$ cd vue
$ npm install
$ npm run serve
```

## Not work with
* `Electron`, `nodeJS`

## Problem
* Lint ignore is always needed. - like `// eslint-disable-next-line`
* How can I use `go wasm file` with vue.prototype? - like `Vue.prototype.$go = { go };`
* `Google Chrome` (improved but,) still have problem to <b>die</b> when refrech sometimes.
