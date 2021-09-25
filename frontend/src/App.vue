<template>
  <div class="demo">
    <el-tabs v-model="mode">
      <el-tab-pane label="标准模式" name="normal"></el-tab-pane>
      <el-tab-pane label="键盘模式" name="keyboard"></el-tab-pane>
    </el-tabs>
    <router-view @calc="calc" />
    <div v-if="ansHistory.length !== 0" id="history">
      <el-table :data="ansHistory" stripe style="width: 100%">
        <el-table-column type="selection" width="auto" />
        <el-table-column prop="correctedExp" label="修正表达式" width="auto" />
        <el-table-column prop="answer" label="结果" width="auto" />
        <el-table-column fixed="right" label="操作" width="auto">
          <template #default="scope">
            <el-button type="primary" icon="el-icon-document-copy" >复制结果</el-button>
            <el-button type="danger" icon="el-icon-delete" @click="handleDelete(scope.$index)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <div v-else>
      <el-empty description="输入一个新的表达式开始计算" image="./src/assets/img/null.png"></el-empty>
    </div>
  </div>
</template>

<style>
.demo {
  position: absolute;
  font-size: 15px;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 40%;
  border: 2px solid grey;
  border-width: auto;
  border-radius: 10px;
  padding: 15px;
  vertical-align: middle;
  text-align: center;
}
</style>

<script>
import router from "./router";
import axios from "./utils/axios";
import message from "./utils/message";
export default {
  name: "app",
  data() {
    return {
      mode: "keyboard",
      ansHistory: [],
    };
  },
  methods: {
    calc(newExp) {
      message.clear();
      let errorMsg = "";
      let correctedExp = "";
      let answer = NaN;
      let data = { "orig-exp": newExp.origExp };
      let ansHistory = this.ansHistory;
      console.log(data);
      let resHandler = function(result){
        errorMsg = result["error-msg"];
        correctedExp = result["corrected-exp"];
        answer = result["answer"];
        console.log(result);
        if (errorMsg !== "") {
          message.warning(`${errorMsg}`);
        } else {
          console.log(this.ansHistory);
          ansHistory.push({correctedExp:correctedExp,answer:answer});
          this.ansHistory=ansHistory;
        }
      };
      axios.post({
        url: "http://localhost:3001/process",
        data: data,
        loading: false,
        timeout: 100,
        confirm: false,
        success: resHandler,
      });
    },
    handleDelete(index) {
      this.ansHistory.splice(index, 1);
    }
  },
  watch: {
    mode(newMode) {
      router.push({ name: `${newMode}` });
      this.origExp = "";
      this.errorMsg = "";
      this.correctedExp = "";
      this.answer = NaN;
    },
  },
  mounted() {
    router.push({ name: "keyboard" });
  },
};
</script>