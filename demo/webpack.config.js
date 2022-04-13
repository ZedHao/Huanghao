const VueLoaderPlugin = require('vue-loader/lib/plugin');
module.exports = {
    devtool: "sourcemap",
    entry: './src/main.js', // 入口文件
    output: {
        filename: 'bundle.js' // 打包出来的wenjian
    },
    rules:[{test:"",loader: 'vue-loader'}],
    plugins: [
        // make sure to include the plugin for the magic
        new VueLoaderPlugin()
    ],
    module : {
    ...
``}
}
