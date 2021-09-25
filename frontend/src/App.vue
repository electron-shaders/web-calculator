<template>
  <div class="demo">
    <el-tabs v-model="mode">
      <el-tab-pane label="Normal" name="normal"></el-tab-pane>
      <el-tab-pane label="Keyboard" name="keyboard"></el-tab-pane>
    </el-tabs>
    <router-view @calc="calc" />
    {{ origExp }}
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
import axios from "axios"
export default {
  name: "app",
  data() {
    return {
      mode: "keyboard",
      origExp: "",
      error: "",
      correctedExp: "",
      answer: NaN,
    };
  },
  methods: {
    calc(newExp) {
      this.origExp = newExp.origExp;
      var data = { "orig-exp": this.origExp };
      console.log(data);
      axios({
        method: "POST",
        url: "http://127.0.0.1:3001/process",
        data: data,
        headers: { "content-type": "text/plain" },
      })
        .then((result) => {
          this.error = result.data["error"];
          this.correctedExp = result.data["corrected-exp"];
          this.answer = result.data["answer"];
          console.log(result.data);
        })
        .catch((error) => {
          console.error(error);
        });
    },
  },
  watch: {
    mode(newMode) {
      router.push({ name: `${newMode}` });
      this.origExp = "";
      this.error = "";
      this.correctedExp = "";
      this.answer = NaN;
    },
  },
  mounted() {
    router.push({ name: "keyboard" });
  },
};
</script>