<template>
    <el-row>
        <el-col :span="23" :offset="1">
            <el-form label-position="right" label-width="200px">
                <el-row>
                    <el-form-item label="API KEY">
                        <el-input v-model="user.api.key" type="password" show-password controls-position="right" />
                    </el-form-item>
                </el-row>
                <el-row>
                    <el-form-item label="API SECRET">
                        <el-input v-model="user.api.secret" type="password" show-password controls-position="right" />
                    </el-form-item>
                </el-row>
            </el-form>
        </el-col>
    </el-row>
</template>

<script>
export default {
    props: {
        user: Object
    },

    data() {
        return {
            _user: {
                api: {
                    key: '',
                    secret: '',
                }
            }
        }
    }

}
</script>

<style scoped>

</style>