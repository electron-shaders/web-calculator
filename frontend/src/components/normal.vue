<template>
    <div>
        <el-input
            ref="input-box"
            disabled
            v-model="origExp"
            @keyup.enter.native="this.$emit('calc', { origExp });"
            autosize
            type="text"
        >
            <template #append>
                <el-button size="mini" @click="origExp = origExp.substring(0, origExp.length - 1)">
                    <span class="iconfont el-icon-third-tuige" style="vertical-align: middle;" />
                </el-button>
            </template>
        </el-input>
        <el-row :gutter="10">
            <el-col :span="6">
                <el-button type="primary" @click="origExp += '('">(</el-button>
            </el-col>
            <el-col :span="6">
                <el-button type="primary" @click="origExp += ')'">)</el-button>
            </el-col>
            <el-col :span="6">
                <el-button type="primary" @click="origExp += '^2'">
                    x<sup>2</sup>
                </el-button>
            </el-col>
            <el-col :span="6">
                <el-button type="primary" @click="origExp += '√'">√</el-button>
            </el-col>
        </el-row>
        <el-row :gutter="10">
            <el-col :span="6">
                <el-button type="primary" @click="origExp += '7'">7</el-button>
            </el-col>
            <el-col :span="6">
                <el-button type="primary" @click="origExp += '8'">8</el-button>
            </el-col>
            <el-col :span="6">
                <el-button type="primary" @click="origExp += '9'">9</el-button>
            </el-col>
            <el-col :span="6">
                <el-button type="primary" @click="origExp += '/'">÷</el-button>
            </el-col>
        </el-row>
        <el-row :gutter="10">
            <el-col :span="6">
                <el-button type="primary" @click="origExp += '4'">4</el-button>
            </el-col>
            <el-col :span="6">
                <el-button type="primary" @click="origExp += '5'">5</el-button>
            </el-col>
            <el-col :span="6">
                <el-button type="primary" @click="origExp += '6'">6</el-button>
            </el-col>
            <el-col :span="6">
                <el-button type="primary" @click="origExp += '*'">×</el-button>
            </el-col>
        </el-row>
        <el-row :gutter="10">
            <el-col :span="6">
                <el-button type="primary" @click="origExp += '1'">1</el-button>
            </el-col>
            <el-col :span="6">
                <el-button type="primary" @click="origExp += '2'">2</el-button>
            </el-col>
            <el-col :span="6">
                <el-button type="primary" @click="origExp += '3'">3</el-button>
            </el-col>
            <el-col :span="6">
                <el-button type="primary" @click="origExp += '-'">-</el-button>
            </el-col>
        </el-row>
        <el-row :gutter="10">
            <el-col :span="6">
                <el-button type="primary" @click="origExp += '0'">0</el-button>
            </el-col>
            <el-col :span="6">
                <el-button type="primary" @click="origExp += '.'">.</el-button>
            </el-col>

            <el-col :span="6">
                <el-button type="primary" @click="origExp += '%'">%</el-button>
            </el-col>
            <el-col :span="6">
                <el-button type="primary" @click="origExp += '+'">+</el-button>
            </el-col>
        </el-row>
        <el-row :gutter="10">
            <el-col :span="12">
                <el-button type="warning" @click="origExp = ''">C</el-button>
            </el-col>
            <el-col :span="12">
                <el-button type="success" @click="this.$emit('calc', { origExp });" :loading="isLoading">=</el-button>
            </el-col>
        </el-row>
    </div>
</template>

<style scoped>
.el-button {
    width: 100%;
    height: 20%;
}
.el-input {
    margin: 10px 0;
}
.el-row {
    margin-bottom: 2%;
}
</style>

<script>
import store from "../store"
export default {
    name: "normal",
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
    mounted() {
        this.$refs['input-box'].focus()
    },
}
</script>