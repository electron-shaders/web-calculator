<template>
  <div class="demo">
    <el-tabs v-model="currentPage">
      <el-tab-pane label="Home" name="home"></el-tab-pane>
      <el-tab-pane label="Demo" name="demo"></el-tab-pane>
    </el-tabs>
    <router-view/>
    <p class="log">{{ storeLog }}</p>
    <p class="log">{{ axiosLog }}</p>
  </div>
</template>

<style>
.demo {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  border: 2px solid grey;
  border-width: auto;
  border-radius: 10px;
  padding: 15px;
  vertical-align: middle;
  text-align: center;
}
.log {
  text-align: left;
  color: white;
  padding: 3px;
  padding-left: 4px;
  border-left: 4px solid darkgrey;
  background-color: grey;
}
</style>

<script>
import { send } from "./api/app"
import router from "./router"
export default {
  name: "app",
  data() {
    return {
      answer: "",
      currentPage: "home",
      storeLog: "store>>> " + this.$store.state.app
    }
  },
  computed: {
    axiosLog() {
      return "res from https://yesno.wtf/api>>> " + this.answer;
    }
  },
  watch:{
    currentPage(newPage,oldPage){
      router.push({name:`${newPage}`})
    }
  },
  async mounted() {
    const res = await send();
    this.answer = res.answer;
  }
}
</script>
