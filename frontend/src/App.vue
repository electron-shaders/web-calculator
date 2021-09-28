<template>
  <el-container id="container" @keyup.esc.exact="message.clear()">
    <el-aside class="hidden-sm-and-down" width="35%">
      <history ref="history"></history>
    </el-aside>
    <el-main>
      <el-tabs v-model="mode">
        <el-tab-pane label="标准模式" name="normal"></el-tab-pane>
        <el-tab-pane label="键盘模式" name="keyboard"></el-tab-pane>
      </el-tabs>
      <router-view @calc="calc" />
    </el-main>
  </el-container>
</template>

<style>
#tab-operations {
  position: fixed;
  bottom: 0%;
  padding: 1% 0;
  text-align: left;
}
#container {
  height: 92%;
}
</style>

<script>
import router from "./router";
import axios from "./utils/axios";
import message from "./utils/message";
import history from "./components/history.vue"
export default {
  name: "app",
  data() {
    return {
      mode: "normal",
    };
  },
  methods: {
    calc(newExp) {
      let errorMsg = "";
      let correctedExp = "";
      let answer = NaN;
      let data = { "orig-exp": newExp.origExp };
      let ansHistory = this.$refs['history'].getAnsHistory();
      let resHandler = function (result) {
        message.clear();
        errorMsg = result["error-msg"];
        correctedExp = result["corrected-exp"];
        answer = result["answer"];
        if (errorMsg !== "") {
          message.warning(`${errorMsg}`);
        } else {
          if (window.innerWidth <= 768) {
            //TODO: 小屏自动复制结果
          }
          //TODO: 结果自动上输入框
          ansHistory.unshift({
            correctedExp: correctedExp,
            answer: answer,
            index: 0,
          });
          for (let i = 1; i < ansHistory.length; i++) {
            ansHistory[i].index = i;
          }
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
    router.push({ name: "normal" });
  },
  components: {
    history
  }
};
</script>