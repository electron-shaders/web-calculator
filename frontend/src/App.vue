<template>
  <div class="demo">
    <el-tabs v-model="mode">
      <el-tab-pane label="Normal" name="normal"></el-tab-pane>
      <el-tab-pane label="Keyboard" name="keyboard"></el-tab-pane>
    </el-tabs>
    <router-view @calc="calc" />
  </div>
</template>

<style>
.demo {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 28%;
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
      origExp: "",
      errorMsg: "",
      correctedExp: "",
      answer: NaN,
    };
  },
  methods: {
    calc(newExp) {
      this.origExp = newExp.origExp;
      var data = { "orig-exp": this.origExp };
      console.log(data);
      var resHandler = function(result){
        this.errorMsg = result["error-msg"];
        this.correctedExp = result["corrected-exp"];
        this.answer = result["answer"];
        console.log(result);
        if (this.errorMsg !== "") {
          message.warning(`${this.errorMsg}`);
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
    router.push({ name: "keyboard" });
  },
};
</script>