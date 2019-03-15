// vue.config.js
module.exports = {
  devServer: {
    proxy: 'http://localhost:4000'
  },
  chainWebpack: config => {
    config.module
      .rule('wasm')
      .test(/\.wasm$/)
      .use('wasm-loader')
     .loader('wasm-loader')
     .end()
  }
}
