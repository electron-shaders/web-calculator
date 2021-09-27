<template>
    <div id="orig-exp-textarea">
        <el-input ref="input-box" v-model="origExp" @keyup.enter.native="this.$emit('calc',{origExp}); origExp=''" autosize type="text" placeholder="请输入一个表达式......">
            <template #append>
                <el-button type="primary" @click="this.$emit('calc',{origExp}); origExp=''">
                    <span class="iconfont el-icon-third-calculator" style="vertical-align: middle;"> 计算 </span>                
                </el-button>
            </template>
        </el-input>
    </div>
</template>

<script>
export default {
    name: "keyboard",
    data(){
        return{
            origExp:""
        }
    },
    mounted() {
        this.$refs['input-box'].focus()
    },
}
</script>