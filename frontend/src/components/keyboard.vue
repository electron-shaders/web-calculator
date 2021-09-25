<template>
    <div id="orig-exp-textarea">
        <link rel="stylesheet" href="../iconfont.css">
        <el-input v-model="origExp" autosize type="text" placeholder="Please input an expression">
            <template #append>
                <el-button type="primary" @click="this.$emit('calc',{origExp})">
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
}
</script>