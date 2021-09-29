<template>
  <div>
    <el-input
      ref="input-box"
      v-model="origExp"
      @keyup.enter.native="this.$emit('calc', { origExp });"
      autosize
      type="text"
      placeholder="请输入一个表达式......"
    >
      <template #append>
        <el-button type="primary" @click="this.$emit('calc', { origExp });">
          <span class="iconfont el-icon-third-calculator" style="vertical-align: middle;"> 计算 </span>
        </el-button>
      </template>
    </el-input>
  </div>
</template>

<style scoped>
.el-input {
  margin: 10px 0;
}
</style>

<script>
import store from "../store"
export default {
  name: "keyboard",
  mounted() {
    this.$refs['input-box'].focus()
  },
  computed: {
    origExp: {
      get() {
        return this.$store.state.origExp;
      },
      set(newVal) {
        store.commit('setOrigExp', newVal)
      }
    },
    isLoading() {
      return this.$store.state.isLoading;
    }
  },
}
</script>