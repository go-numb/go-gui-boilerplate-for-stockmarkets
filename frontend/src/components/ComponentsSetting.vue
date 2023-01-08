<template>
    <el-row>
        <el-col :span="23" :offset="1">
            <el-row>
                <el-form label-position="right" label-width="200px">
                    <el-row>
                        <el-form-item label="TOKEN">
                            <el-input v-model="setting.token" controls-position="right" />
                        </el-form-item>
                    </el-row>
                    <el-row>
                        <el-form-item label="テストタイプ">
                            <el-checkbox-group v-model="setting.types">
                                <el-checkbox label="Online activities" name="type" />
                                <el-checkbox label="Promotion activities" name="type" />
                                <el-checkbox label="Offline activities" name="type" />
                                <el-checkbox label="Simple brand exposure" name="type" />
                            </el-checkbox-group>
                        </el-form-item>
                    </el-row>
                    <el-row>
                        <el-form-item label="注文タイプ">
                            <el-radio-group v-model="setting.tif">
                                <el-radio label="GTC" />
                                <el-radio label="IOC" />
                                <el-radio label="FOK" />
                            </el-radio-group>
                        </el-form-item>
                    </el-row>
                </el-form>
            </el-row>
        </el-col>
    </el-row>
</template>

<script>
export default {
    props: {
        setting: Object
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