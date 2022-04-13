module.exports = {
  root: true,
  parser: 'babel-eslint',
  parserOptions: {
    sourceType: 'module'
  },
  env:{
    "browser": true,
    "node": true,
    "es6": true
  },
  // required to lint *.vue files
  // extends: 'vue',
  plugins: [
    'vuefix',
    'html',
  ],
  // add your custom rules here
  rules: {}
}
/*module.export  = {
    rout: true,
    env: {
        browser: true,
        node: true
    },
    parserOptions: {
        parser: 'babel-eslint'
    },
    extends: [
        'plugin:vue/recommended'
    ],
    plugins: [
        'vue'
    ],
    // 添加自定义规则
    rules: {}
}*/


