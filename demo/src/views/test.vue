<template>

  <div>
    <!--布局容器-->
      <el-container class ="el-container" style="background: crimson;" direction="vertical" >
        <!--顶部栏-->
         <el-aside class="el-aside" style="background: greenyellow">
            左侧侧边栏
          </el-aside>
        <el-container class ="el-container1" >
          <!--左侧侧边栏-->
          <el-header class ="el-header" style="background: olivedrab">
          顶部栏
        </el-header>

          <!--主要区域容器-->
          <el-main class="el-main"  style="background: aqua">
            主要区域容器
          </el-main>
        </el-container>
      </el-container>

  </div>
</template>
<script type="text/ecmascript-6">

    // import ElAside from "element-ui/packages/aside/src/main";
    // import ElMain from "element-ui/packages/main/src/main";

    export default {
      // components: {
      //   ElMain,
      //   ElAside},
      data() {
            return {
              message:"你好世界",
            }
        },
        methods:{

        },
    }

</script>
<style>

 .el-container{
  //height: 100%;
   min-height:100vh;
    width: 100%;
   flex-direction:row;
   display: flex
  }
  .el-container1{
 flex-direction:column;
      flex:80;
  width: 200px;

  }
  .el-header{
       display: flex;
height: 100%;



  }
    .el-aside{
      flex:20;
 flex-direction:column;

  }
      .el-main{
       display: inline-flex;
height: 100%;
  }
</style>
