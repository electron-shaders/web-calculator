<template>
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
      <el-table-column v-if="this.windowWidth >= 1440" prop="correctedExp" label="修正表达式" width="auto" />
      <el-table-column prop="answer" fixed="right" label="结果" width="auto" />
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
        circle
        type="danger"
        :disabled="selected.length === 0"
        icon="el-icon-delete"
        @click="delSelected"
      />
    </div>
  </div>
  <div v-else>
    <el-empty description="输入一个新的表达式开始计算" image="./src/assets/img/null.svg"></el-empty>
  </div>
</template>

<script>
import message from "../utils/message";
export default {
  name: "history",
  data() {
    return {
      ansHistory: [],
      selected: [],
      windowWidth: NaN,
    }
  },
  methods: {
    getAnsHistory() {
      return this.ansHistory;
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
  mounted() {
    this.windowWidth=window.innerWidth;
    window.onresize = () => {
      return (() => {
        this.windowWidth=window.innerWidth;
      })()
    }
  },
}
</script>