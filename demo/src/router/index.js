import Vue from 'vue'
import Router from 'vue-router'
import Home from '../views/home' // 导入我们编写的vue组件
import TcsPermission from '../views/tcs_permission' // 导入我们编写的vue组件
import ReviewRecord from '../views/review_record' // 导入我们编写的vue组件
import Test2 from '../views/test2' // 导入我们编写的vue组件
import Wellcome from '../views/wellcome' // 导入我们编写的vue组件
import Test from '../views/test' // 导入我们编写的vue组件
import ElementUI from 'element-ui'

Vue.use(Router)
Vue.use(ElementUI);

export default new Router({
  routes: [
    { path: '/', redirect: '/home' },
    { path: '/test',  name: 'test', component: Test },

    { path: '/home', name: 'home', component: Home, meta: {title: 'wellcome'}, redirect:'wellcome',
        children: [
        { path: '/tcs_permission', name: 'tcs_permission', component: TcsPermission, meta: {title: 'tcs队列权请'}},
         { path: '/review_record', name: 'review_record', component: ReviewRecord, meta: {title: '打压审计'}},
        { path: '/wellcome', name: 'wellcome', component: Wellcome, meta: {title: 'wellcome'}},
        { path: '/test2', name: 'test2', component: Test2, meta: {title: '打压审计'}},


        ]
    },
  ],
})
