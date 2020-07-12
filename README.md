# Go 1.14 WASM Vue.js practice
A `WASM Calculator` practice within `Vue.js`

## Need
* `Vue-Cli3`
* `Go 1.14`

## Run
```sh
$ cd wasm
set GOARCH=wasm
set GOOS=js
$ go build

$ cp (or copy) calc.wasm ../vue/public/ (or ../vue/assets)

$ cd ..

$ cd vue
$ npm install
$ npm run serve
```

## Server
* A http server for testing `index.html` in `wasm` folder

## Not tested with
* `Electron`

## wasm_exec.js
Few modified from `$GOROOT/misc/wasm/wasm_exec.js`

## Problem
* Cannot make to load golang codes before vue codes loading. ES-Lint ignore set are always needed like `// eslint-disable-next-line` or `/* global function_name */`
