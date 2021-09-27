<template>
  <el-container id="container" @keyup.esc.exact="message.clear()">
    <el-aside width="35%">
      <div v-if="ansHistory.length !== 0">
        <el-table
          :data="ansHistory"
          fit
          highlight-current-row
          style="width: 100%"
          @keyup.delete.native="delSelected"
          @selection-change="handleSelectionChange"
        >
          <el-table-column type="selection" width="auto" />
          <el-table-column
            prop="correctedExp"
            label="修正表达式"
            width="auto"
          />
          <el-table-column
            prop="answer"
            fixed="right"
            label="结果"
            width="auto"
          />
          <el-table-column fixed="right" label="操作" width="auto">
            <template #default="scope" style="text-align: right">
              <el-button
                circle
                type="primary"
                size="mini"
                icon="el-icon-document-copy"
                @click.prevent="copyEle(scope.row.answer)"
              ></el-button>
              <el-button
                circle
                type="danger"
                size="mini"
                icon="el-icon-delete"
                @click.prevent="delEle(scope.$index)"
              ></el-button>
            </template>
          </el-table-column>
        </el-table>
        <div id="tab-operations">
          <el-button
            type="danger"
            :disabled="selected.length === 0"
            icon="el-icon-delete"
            @click="delSelected"
            >删除选中</el-button
          >
        </div>
      </div>
      <div v-else>
        <el-empty
          description="输入一个新的表达式开始计算"
          image="./src/assets/img/null.svg"
        ></el-empty>
      </div>
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
export default {
  name: "app",
  data() {
    return {
      mode: "keyboard",
      ansHistory: [],
      selected: [],
    };
  },
  methods: {
    calc(newExp) {
      let errorMsg = "";
      let correctedExp = "";
      let answer = NaN;
      let data = { "orig-exp": newExp.origExp };
      let ansHistory = this.ansHistory;
      let resHandler = function (result) {
        message.clear();
        errorMsg = result["error-msg"];
        correctedExp = result["corrected-exp"];
        answer = result["answer"];
        if (errorMsg !== "") {
          message.warning(`${errorMsg}`);
        } else {
          ansHistory.unshift({
            correctedExp: correctedExp,
            answer: answer,
            index: 0,
          });
          for (let i = 1; i < ansHistory.length; i++) {
            ansHistory[i].index = i;
          }
          this.ansHistory = ansHistory;
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
    handleSelectionChange(val) {
      this.selected = val;
    },
    delSelected() {
      if (this.selected.length === 0) {
        message.warning("未选中任何结果");
      } else {
        this.selected.sort(function (a, b) {
          let x = a.index;
          let y = b.index;
          return x < y ? 1 : x > y ? -1 : 0;
        });
        for (let i = 0; i < this.selected.length; i++) {
          this.ansHistory.splice(this.selected[i].index, 1);
        }
        for (let i = 0; i < this.ansHistory.length; i++) {
          this.ansHistory[i].index = i;
        }
        message.success("已删除");
      }
    },
    delEle(index) {
      this.ansHistory.splice(index, 1);
      message.success("已删除");
    },
    copyEle: function (val) {
      this.$copyText(val).then(
        function (e) {
          message.success("复制成功");
        },
        function (e) {
          message.success("复制失败");
        }
      );
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