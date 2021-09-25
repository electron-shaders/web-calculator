<template>
    <div id="orig-exp-textarea">
        <el-input v-model="origExp" autosize type="textarea" placeholder="Please input an expression"/>
        <el-button type="primary" @click="this.$emit('calc',{origExp})">计算</el-button>
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