<template>
  <el-container id="container" @keyup.esc.exact="message.clear()">
    <el-aside class="hidden-sm-and-down" width="35%">
      <history ref="history" :window-width="windowWidth"></history>
    </el-aside>
    <el-main>
      <el-tabs v-model="mode">
        <el-tab-pane label="标准模式" name="normal"></el-tab-pane>
        <el-tab-pane label="键盘模式" name="keyboard"></el-tab-pane>
        <el-tab-pane v-if="this.windowWidth <= 992" label="历史记录" name="history"></el-tab-pane>
      </el-tabs>
      <normal v-if="mode==='normal'" @calc="calc"></normal>
      <keyboard v-if="mode==='keyboard'" @calc="calc"></keyboard>
      <history v-if="mode==='history'" :window-width="windowWidth"></history>
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
  height: 95%;
}
</style>

<script>
import axios from "./utils/axios";
import message from "./utils/message";
import history from "./components/history.vue"
import keyboard from "./components/keyboard.vue"
import normal from "./components/normal.vue"
import store from "./store";
export default {
  name: "app",
  data() {
    return {
      mode: "normal",
      windowWidth: NaN,
    };
  },
  methods: {
    calc(newExp) {
      this.isLoading=true;
      let errorMsg = "";
      let correctedExp = "";
      let answer = NaN;
      let data = { "orig-exp": newExp.origExp };
      let resHandler = (result) => {
        this.isLoading=false;
        message.clear();
        errorMsg = result["error-msg"];
        correctedExp = result["corrected-exp"];
        answer = result["answer"];
        if (this.windowWidth <= 992) {
          message.success(`计算结果: ${answer}`)
        }
        if (errorMsg !== "") {
          message.warning(`${errorMsg}`);
          this.origExp="";
        } else {
          this.origExp=answer.toString();
          this.$store.commit('updateAnsHistory',{
            correctedExp: correctedExp,
            answer: answer,
            index: 0,
          });
        }
      };
      let errHandler = (error) => {
        this.isLoading=false;
        this.origExp="";
        message.alert({
          title: '后端出错', msg: `${error}`
        });
      }
      axios.post({
        url: "http://localhost:3001/process",
        data: data,
        loading: false,
        timeout: 100,
        confirm: false,
        success: resHandler,
        error: errHandler,
      });
    },
  },
  computed: {
    origExp: {
      set(newVal) {
        store.commit('setOrigExp', newVal)
      }
    },
    isLoading: {
      get() {
        return this.$store.state.isLoading;
      },
      set(newVal) {
        store.commit('setLoading', newVal);
      }
    }
  },
  watch:{
    windowWidth(newValue){
      if(newValue>=993 && this.mode=="history"){
        this.mode="normal";
      }
    }
  },
  mounted() {
    this.windowWidth=window.innerWidth;
    window.onresize = () => {
      return (() => {
        this.windowWidth=window.innerWidth;
      })()
    }
  },
  components: {
    history,
    normal,
    keyboard,
  }
};
</script>