<template>
    <el-container class="home-container">
        <!-- 侧边栏 -->
        <el-aside class="el-aside">
            <!-- 头部logo -->
            <template>
                <div class="header-title" v-if="!isCollapse">
                    <img width="35px" height="35px" src="../assets/logo.png" style="border-radius: 5px">
                    <span style="font-weight:bold">团队小工具</span>
                </div>
                <div class="header-title-hiddle" v-else>
                    <img width="40px" height="40px" src="../assets/logo.png" style="border-radius: 5px">
                </div>
            </template>
            <!-- 菜单区域 -->
            <el-menu class="el-menu"
                     background-color="#6A75CA"
                     text-color="#fff"
                     active-text-color="#fff"
                     :collapse="isCollapse"
                     :collapse-transition="false"
                     unique-opened
                     router>
                <!-- 首页 -->
                <el-menu-item
                        class="el-menu-item"
                        :index="item.path" v-for="item in menuList"
                        :miss="item.id"
                        @click="selectMenu(item)">
                    <i class="el-icon-location"></i>
                    <span>{{item.label}}</span>
                </el-menu-item>
            </el-menu>
        </el-aside>
        <!-- 主体 -->
        <el-container class="home-container1">
            <!-- 头部 -->
            <el-header>
                <div class="header-row">
                    <div class="toggle-button">
                        <div>
                            <i :title="isCollapse ? '展开' : '收起'" class="fa fa-bars" @click="toggleCollapse"></i>
                        </div>
                        <div style="height:25px;width:100%;margin-left:10px">
                            <div style="float:left;color:#E74405;font-size:16px;height:25px;line-height:25px;">
                                <i class="fa fa-bullhorn"></i>
                            </div>
                            <el-carousel height="25px" direction="vertical" indicator-position="none" autoplay
                                         :interval="8000">
                                <el-carousel-item v-for="item in mottoList" :miss="item">
                                    <h3 class="medium">{{ item }}</h3>
                                </el-carousel-item>
                            </el-carousel>
                        </div>
                    </div>
                    <div class="header-avatar">
                        <div class="user">
                            管理员，您好！
                        </div>
                        <el-dropdown @command="handleCommand">
              <span class="el-dropdown-link">
                <img width="35" height="35" style="border-radius:50%;background:#dddddd" src="../assets/logo.png"
                     alt="">
                <i class="el-icon-arrow-down el-icon&#45;&#45;right"></i>
              </span>
                            <el-dropdown-menu slot="dropdown">
                                <el-dropdown-item command="update-123">123</el-dropdown-item>
                                <el-dropdown-item command="logout">123</el-dropdown-item>
                            </el-dropdown-menu>
                        </el-dropdown>
                    </div>
                </div>
            </el-header>
            <!-- 内容区 -->
            <el-main class="el-main">
                <!-- 路由占位符，用于展示内容区的内容 -->
                <div style="padding:20px">
                    <keep-alive>
                        <router-view/>
                    </keep-alive>
                </div>
            </el-main>
        </el-container>
    </el-container>
</template>
<script>
    export default {
        components: {},
        //组件被创建
        created() {
            //加载菜单
            this.loadMenu();
        },
        computed: {},
        data() {
            return {
                //菜单列表
                menuList: [],
                // 折叠-展开 默认false不折叠
                isCollapse: false,
                // 系统公告
                mottoList: [
                    '来一个系统公告'
                ],
            }
        },
        methods: {
            // 右上角下拉菜单点击事件
            handleCommand(command) {
                switch (command) {
                    // 退出
                    case 'logout':
                        //消息提示
                        this.msg.success('123')
                        //重置vuex中的数据
                        this.$store.commit('clearVUEX')
                        //跳转到首页
                        this.$router.push("/index");
                        break
                    //修12
                    case 'update-123':
                        //消息提示
                        this.msg.success('123')
                        break
                }
            },
            // 点击折叠 展开菜单
            toggleCollapse() {
                this.isCollapse = !this.isCollapse;
            },
            // 点击菜单 - 传入name，添加到keepalive缓存页面
            selectMenu(item) {
                // 加入keepalive缓存
                this.$store.commit('addKeepAliveCache', item.name)
                //添加tags标签
                //访问wellcome 就代表home
                var name = item.name === 'wellcome' ? 'home' : item.name
                var submenu = {
                    path: name,
                    name: name,
                    label: item.label
                }
                //修改选中菜单
                this.$store.commit('selectMenu', submenu)
            },
            //加载菜单
            loadMenu() {
                this.menuList = [

                               {
                                   id: 'number-02',
                                   class: 'el-icon-document',
                                   path: '/wellcome',
                                   label: '首页',
                                   name: 'wellcome'
                               },
                    {
                        id: 'number-03',
                        class: 'el-icon-document',
                        path: '/tcs_permission',
                        label: 'tcs队列权限',
                        name: 'tcs_permission'
                    },
                   {
                        id: 'number-04',
                        class: 'el-icon-document',
                       path: '/review_record',
                        label: '打压通知审计',
                        name: 'review_record'
                    },
                ]
            }
        },
    };
</script>
<style lang="less" scoped>
    .home-container {
        //height: 100%;
        min-height: 100vh;
        width: 100%;
        flex-direction: row;
        display: flex
    }

    .home-container1 {
        flex-direction: column;
        flex: 90;
    }

    .el-aside {
        background-color: #6A75CA;
        flex-direction: column;
        flex: 10;

        .header-title {
            padding-left: 10px;
            height: 60px;
            font-weight: 300;
            display: flex;
            font-size: 20px;
            align-items: center;
            cursor: pointer;
            color: #ffffff;

            span {
                margin-left: 10px;
            }
        }

        .header-title-hiddle {
            width: 64px;
            height: 60px;
            display: table-cell;
            vertical-align: middle;
            text-align: center;
            color: #ffffff;
            cursor: pointer;
        }

        .el-menu {
            border-right: none;
            display: flex;
            flex-direction: column;
        }
    }

    .el-main {
        background-color: #eaedf1;
        height: 100vh;
        display: flex;
        flex-direction: column;
    }

    .el-header {
        color: rgb(0, 0, 0);
        font-size: 20px;
        border-bottom: 1px solid #dddddd;
        height: 103px !important;
        padding: 0;
        background: #fff;
    }

    .header-row {
        height: 60px;
        width: 100%;
        display: flex;
        flex-direction: row;
        justify-content: center;
        border-bottom: 1px solid #dddddd;
        overflow: hidden;
    }

    ul li {
        list-style-type: none;
        margin: 0;
        padding: 0;
    }

    ul {
        padding-left: 0px;
    }

    li {
        display: inline
    }

    .header-avatar {
        float: right;
        width: 40%;
        display: flex;
        align-items: center;
        justify-content: flex-end;
        padding-right: 20px;

        .user {
            font-size: 14px;
            font-weight: bold;
            padding: 0 10px;
        }
    }

    .el-col-align-middle {
        line-height: 40px;
        text-align: left;
        font-size: 14px;
    }

    // 菜单选中背景色
    .el-menu-item.is-active {
        background-color: #1890FF !important;
    }

    // 菜单悬浮背景色
    .el-menu-item:hover {
        background-color: #1890FF !important;
    }

    .el-menu-item {
        border-radius: 10px;
        height: 50px;
        width: 150px;
    }

    // 走马灯
    .el-carousel__item h3 {
        color: #ee7c12;
        font-size: 14px;
        opacity: 0.75;
        line-height: 25px;
        margin: 0;
    }

    .fa {
        margin-right: 10px;
    }

    // 点击展开/折叠按钮
    .toggle-button {
        width: 80%;
        font-size: 20px;
        line-height: 40px;
        color: #595959;
        text-align: left;
        display: flex;
        align-items: center;
        float: left;
        padding-left: 20px;

        i {
            cursor: pointer;
        }
    }

    // 面包屑导航
    .el-breadcrumb {
        margin-bottom: 0;
    }

    // tab页
    .tabs-switch-page {
        display: flex;
        align-items: center;
        height: 40px;
        background-color: #fff;
        overflow: hidden;
    }

    .el-tag {
        cursor: pointer;
        margin-left: 10px;
        border-radius: 2px;
        font-size: 12px;
        color: #1890FF;
        border-color: #1890FF;
    }

    .el-tag--dark {
        color: #fff;
        background-color: #1890FF;
    }

    .el-dropdown-link {
        cursor: pointer;
    }

    .el-icon-arrow-down {
        font-size: 12px;
    }

    .submit-row {
        display: flex;
        flex-direction: row;
        justify-content: flex-end;
        align-items: center;
    }
</style>
