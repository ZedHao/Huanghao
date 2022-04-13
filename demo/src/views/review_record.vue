<template>
    <el-main class="el-main">

        <el-form class="el-form" ref="form" :model="loginForm">
            <el-form-xxx prop="gid">
                <el-input class="el-input" id="gid" type="gid" v-model="loginForm.gid" placeholder="gid"></el-input>
            </el-form-xxx>
            <el-button

                        class="el-button" type="primary" @click="handleLogin">搜索</el-button>




        </el-form>
        <a target="_blank"
           href="https://misb"
           class="text-decoration-none"
        >
            打压通知审计手册
        </a>

             <el-table
                :data="suppressInfo"
                   element-loading-text = "加载中，请稍后..."
                :row-class-name="tableRowClassName"
                :xxx="ite12"
                :cell-style="rowClass"
                style="width: 170vh"
        >
                                                               <el-table-column
                    prop="gid"
                    label="gid"
            > </el-table-column>
            <el-table-column
                    prop="suppress_cp_status"
                    label="申诉状态"
            > </el-table-column>
                             <el-table-column
                    prop="audit_suppress_reason"
                    label="mp外漏打压理由"
            > </el-table-column>
                             <el-table-column
                    prop="show_enter"
                    label="是否展示申诉入口"
            > </el-table-column>
                             <el-table-column
                    prop="valid_appeal"
                    label="是否允许申诉"
            > </el-table-column>
                             <el-table-column
                    prop="forbid_reason"
                    label="mp外漏文案"
            > </el-table-column>

        </el-table>
        <el-table
                :data="tableData"
                   element-loading-text = "加载中，请稍后..."
                :row-class-name="tableRowClassName"
                :xxx="itemKey"
                :cell-style="rowClass"
                style="width: 170vh"
        >
            <el-table-column
                    prop="create_time"
                    label="请求通知时间"
            >
            </el-table-column>
            <el-table-column
                    prop="task_id"
                    label="tcs任务id"
            >
            </el-table-column>
            <el-table-column
                    prop="project_name"
                    label="队列审核链接"
            >
                <template slot-scope="scope">
                    <a :href="scope.row.tcs_url" target="_blank">
                        {{scope.row.project_name}}
                    </a>
                </template>
            </el-table-column>
            <el-table-column
                    prop="notification_status_str"
                    label="通知状态"
            >
            </el-table-column>
            <el-table-column
                    prop="errmsg"
                    label="过滤原因"
            >
            </el-table-column>
            <el-table-column
                    prop="fail_reason"
                    label="打压原因"
            >
            </el-table-column>
            <el-table-column prop="verify_result" label="审出结果">
                <template slot-scope="scope">
                    <div style="text-align:center">
                        <el-button class="el-button" type="primary" plain
                                   @click="showIndex(tableData[scope.$index].verify_result)">
                            详情
                        </el-button>
                        <Modal v-show="showModal" width="200"
                               @on-cancel="cancel">
                            <div slot="content">
                                <div class="myBody" style="width:80vh">
                                    <JsonView :json="content"></JsonView>
                                </div>
                            </div>
                        </Modal>

                    </div>
                </template>
            </el-table-column>


            <el-table-column prop="times" label="各环节触达时间">
                <template slot-scope="scope">
                    <div style="text-align:center">
                        <el-button class="el-button" type="primary" plain
                                   @click="showIndex(tableData[scope.$index].times)">
                            详情
                        </el-button>
                        <Modal v-show="showModal" width="200"
                               @on-cancel="cancel">
                            <div slot="content">
                                <div class="myBody" style="width:80vh">
                                    <JsonView :json="content"></JsonView>
                                </div>
                            </div>
                        </Modal>

                    </div>
                </template>
            </el-table-column>
            <el-table-column prop="co123_info" label="配123息">
                <template slot-scope="scope">
                    <div style="text-align:center">
                        <el-button class="el-button" type="primary" plain
                                   @click="showIndex(tableData[scope.$index].123g_info)">
                            详情
                        </el-button>
                        <Modal v-show="showModal" width="200"
                               @on-cancel="cancel">
                            <div slot="content">
                                <div class="myBody" style="width:80vh">
                                    <JsonView :json="content"></JsonView>
                                </div>
                            </div>
                        </Modal>

                    </div>
                </template>
            </el-table-column>
            <el-table-column prop="log_info" label="rd:日志ID&debug">
                <template slot-scope="scope">
                    <div style="text-align:center">
                        <el-button class="el-button" type="primary" plain
                                   @click="showIndex(tableData[scope.$index].log_info)">
                            日志ID
                        </el-button>
                        <Modal v-show="showModal" width="200"
                               @on-cancel="cancel">
                            <div slot="content">
                                <div class="myBody" style="width:80vh">
                                    <JsonView :json="content"></JsonView>
                                </div>
                            </div>
                        </Modal>

                    </div>
                </template>
            </el-table-column>

            <el-table-column prop="req" label="rd:请求详情">
                <template slot-scope="scope">
                    <div style="text-align:center">
                        <el-button class="el-button" type="primary" plain
                                   @click="showIndex(tableData[scope.$index].req)">
                            详情
                        </el-button>
                        <Modal v-show="showModal" width="200"
                               @on-cancel="cancel">
                            <div slot="content">
                                <div class="myBody" style="width:80vh">
                                    <JsonView :json="content"></JsonView>
                                </div>
                            </div>
                        </Modal>

                    </div>
                </template>
            </el-table-column>
        </el-table>


    </el-main>
</template>
<style lang="less" scoped>
    .el-form {
        flex-direction: row;
        display: flex;
    }

    .el-input__inner {
        height: 50px;
        width: 400px;
    }

    .el-button {
        border-radius: 10px;
        padding: 10px;
        width: 80px;
        height: 25px;
        display: flex;
        font-size: 12px;
        background-color: lightgoldenrodyellow;
        justify-content: center
    }

    div {
        display: flex;
        justify-content: center;
    }

</style>
<script>
    import axios from 'axios/index'
    import Modal from './modal.vue'
    import JsonView from './json_view'

    export default {
        methods: {
            getUrlKey: function (name) {
                return decodeURIComponent((new RegExp('[?|&]' + name + '=' + '([^&;]+?)(&|#|;|$)').exec(location.href) || [, ""])[1].replace(/\+/g, '%20')) || null
            },
            tableRowClassName({row}) {
                let fails = ["401", "501", "600"]
                let success = ["410", "700", "710"]
                if (fails.indexOf(row.notification_status) > 0) {
                    return 'warning-row';
                } else if (success.indexOf(row.notification_status) > 0) {
                    return 'success-row';
                }
                return '';
            },
            parseInfo(col) {
                return {
                    create_time: col.create_time,
                    notification_status_str: col.notification_status_str,
                    notification_status: col.notification_status,
                    group_id: col.group_id,
                    times: JSON.stringify(col.times),
                    task_id: col.task_id,
                    tcs_url: col.tcs_url,
                    project_name: col.project_name,
                    fail_reason: col.fail_reason,
                    verify_result: col.verify_result,
                    req: col.req,
                    errmsg: col.errmsg,
                    log_info: JSON.stringify(col.log_info),
                    article_type: col.article_type,
                    config_info: JSON.stringify(col.123),
                    event_type: col.event_type,
                    top_msg: col.top_msg,
                }
            },

            handleLogin() {
                // this.listLoading = true;
                this.$refs['form'].validate(() => {
                    axios.get('https://fn.miss.net/get_record/?item_id=' + this.loginForm.gid)
                        .then(res => {
                            let arr = [];

                            // this.listLoading = false;

                            Object.xxx(res.data.data).forEach(v => {
                                let o = this.parseInfo(res.data.data[v])
                                o[v] = res.data.data[v];
                                arr.push(o)
                            })
                            this.tableData = arr
                            this.itemKey = Math.random()
                        })
                        .catch(() => {
                            // this.listLoading = false;//如果捕获到错误去掉加载效果
                            window.alert(
                                "查询错误请重试"
                            );
                        });
                })
                // https://575l81nm.fn.miss.net/get_suppress/?item_id=7010593644436619783
                this.$refs['form'].validate(() => {
                    axios.get('https://fn.miss.net/get_suppress/?item_id=' + this.loginForm.gid)
                        .then(res1 => {
                            this.suppressInfo = []
                                    this.suppressInfo.push(res1.data.data)
                       /*     this.suppressInfo[0].suppress_cp_status = res1.data.data.suppress_cp_status
                            this.suppressInfo[0].audit_suppress_reason = res1.data.data.audit_suppress_reason
                            this.suppressInfo[0].show_enter = res1.data.data.show_enter
                            this.suppressInfo[0].valid_appeal = res1.data.data.valid_appeal
                            this.suppressInfo[0].forbid_reason = res1.data.data.forbid_reason
                   */         // this.listLoading = false;
                            console.log(res1.data.data)
                        })
                        .catch(() => {
                            // this.listLoading = false;//如果捕获到错误去掉加载效果
                            window.alert(
                                "查询错误请重试"
                            );
                        });
                })

            },
            showIndex(s) {
                this.showModal = !this.showModal;
                this.content = JSON.parse(s)
            },
            cancel() {
                this.showModal = false;
            },
            rowClass() { //表格数据居中显示
                return "text-align:center"
            },
        },
        data() {
            //这样就拿到了参数中的数
            return {
                loginForm: {
                    gid: this.getUrlKey('item_id')
                },
                tableData: [],
                suppressInfo: [],
                itemKey: "",
                showModal: false,
                //listLoading: false,
                content: '',
            }

        },
        components: {
            'Modal': Modal,
            'JsonView': JsonView
        },

    }
</script>
