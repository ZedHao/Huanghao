<template>
    <div class="modal-backdrop">
        <div class="modal" :style="mainStyles">
            <div class="modal-header">
          <slot name="header">表头</slot>
            </div>
            <div class="modal-body" slot="content" style="overflow:auto;"  >
            <slot name="content">body</slot>
            </div>
            <div class="modal-footer">
                <button type="button" class="btn-close" @click="closeSelf">关闭</button>
                <button type="button" class="btn-confirm" @click="confirm">确认</button>
            </div>
        </div>

    </div>
</template>

<style>
    .modal-backdrop {
        position: fixed;
        top: 0;
        right: 0;
        bottom: 0;
        left: 0;
        display: flex;
        justify-content: center;
        align-items: center;
    }

    .modal {
        background-color: #fff;
        box-shadow: 2px 2px 20px 1px;
        overflow-x: auto;
        display: flex;
        flex-direction: column;
        border-radius: 16px;
        width: 700px;
    }

    .modal-header {
        border-bottom: 1px solid #eee;
        color: #313131;
        justify-content: space-between;
        padding: 15px;
        display: flex;
    }

    .modal-footer {
        border-top: 1px solid #eee;
        justify-content: flex-end;
        padding: 15px;
        display: flex;
    }

    .modal-body {
        position: relative;
        padding: 20px 10px;
        height: 80vh;
	text-overflow:ellipsis;overflow:hidden;
	word-wrap:break-word;


    }


    .btn-close, .btn-confirm {
        border-radius: 8px;
        margin-left: 16px;
        width: 56px;
        height: 36px;
        border: none;
        cursor: pointer;
    }

    .btn-close {
        color: #313131;
        background-color: transparent;
    }

    .btn-confirm {
        color: #fff;
        background-color: #2d8cf0;
    }

</style>

<script>
   // https://www.jianshu.com/p/ee2169c1ca07
    export default {
        name: 'Modal',
        props: {
            width: {
                type: [Number, String],//类型检测
                default: 600 //父组件没传width时的默认值
            },
                height: {
                type: [Number, String],//类型检测
                default: 600 //父组件没传width时的默认值
            }
        },
        data() {
            return {}
        },
        computed: {
            //计算属性来响应width属性，实时绑定到相应DOM元素的style上
            mainStyles() {
                let style = {};
                style.width = `${parseInt(this.width*5)}px`;
                style.height = `${parseInt(this.height)}px`;

                return style;
            }
        },
        methods: {
            //响应关闭按钮点击事件，通过$emit api通知父组件执行父组件的on-cancel方法
            closeSelf() {
                this.$emit('on-cancel');
            },
            confirm() {
                this.$emit('on-cancel');
            }
        }
    }
</script>
