<template>
    <div id="orig-exp-textarea">
        <el-input v-model="origExp" autosize type="text" placeholder="Please input an expression">
            <template #append>
                <el-button type="primary" @click="this.$emit('calc',{origExp})">
                    <el-icon style="vertical-align: middle;">
                        <search />
                    </el-icon>
                    <span style="vertical-align: middle;"> 计算 </span>                
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