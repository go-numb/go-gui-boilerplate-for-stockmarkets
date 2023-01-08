<template>
    <div>
        <el-row>
            <el-col :span="12" :xs="24">
                <el-form label-position="right" label-width="200px">
                    <el-row>
                        <el-form-item label="銘柄">
                            <el-input v-model="sp.code" controls-position="right" />
                        </el-form-item>
                    </el-row>
                    <el-row>
                        <el-form-item label="注文価格">
                            <el-input-number v-model="sp.price" :step="sp.step" step-strictly
                                controls-position="right" />
                        </el-form-item>
                    </el-row>
                    <el-row>
                        <el-form-item label="注文枚数">
                            <el-input-number v-model="sp.size" :step="1" controls-position="right" />
                        </el-form-item>
                    </el-row>
                    <el-row>
                        <el-form-item label="トリガー価格">
                            <el-input-number v-model="sp.lossTriggerPrice" :step="sp.step" step-strictly
                                controls-position="right" />
                        </el-form-item>
                    </el-row>
                    <el-row>
                        <el-form-item label="トリガー後の注文価格">
                            <el-input-number v-model="sp.lossFixedPrice" :step="sp.step" step-strictly
                                controls-position="right" />
                        </el-form-item>
                    </el-row>
                </el-form>
            </el-col>
            <el-col :span="12" :xs="24">
                <el-form label-position="right" label-width="200px">
                    <el-row>
                        <el-form-item label="市場">
                            <el-radio-group v-model="sp.exchange">
                                <el-radio :label="1">東証</el-radio>
                                <el-radio :label="3">名証</el-radio>
                                <el-radio :label="5">福証</el-radio>
                                <el-radio :label="6">札証</el-radio>
                            </el-radio-group>
                        </el-form-item>
                    </el-row>
                    <el-row>
                        <el-form-item label="注文タイプ">
                            <el-radio-group v-model="sp.orderType">
                                <el-radio label="limit">指値</el-radio>
                                <el-radio label="market">成行</el-radio>
                            </el-radio-group>
                        </el-form-item>
                    </el-row>
                    <el-row>
                        <el-form-item label="トリガー後の注文タイプ">
                            <el-radio-group v-model="sp.triggerOrderType">
                                <el-radio label="limit">指値</el-radio>
                                <el-radio label="market">成行</el-radio>
                            </el-radio-group>
                        </el-form-item>
                    </el-row>
                    <el-row>
                        <el-form-item label="発注価格区切り">
                            <el-radio-group v-model="sp.step">
                                <el-radio :label="5">5</el-radio>
                                <el-radio :label="10">10</el-radio>
                                <el-radio :label="25">25</el-radio>
                            </el-radio-group>
                        </el-form-item>
                    </el-row>
                </el-form>
            </el-col>
        </el-row>
        <el-row :gutter="1">
            <el-col :span="12" :offset="6">
                <el-row>
                    <el-col :span="12">
                        <el-row justify="center">
                            <el-button :type="!sp.isReady ? '' : 'primary'" @click="sendSpOrder"
                                :disabled="!sp.isReady">
                                {{ sp.isReady ? '買い' : '処理中' }}
                            </el-button>
                        </el-row>
                    </el-col>
                    <el-col :span="12">
                        <el-row justify="center">
                            <el-button :type="!sp.isReady ? '' : 'danger'" @click="sendSpOrder" :disabled="!sp.isReady">
                                {{ sp.isReady ? '売り' : '処理中' }}
                            </el-button>
                        </el-row>
                    </el-col>
                </el-row>
            </el-col>
        </el-row>
    </div>
</template>

<script>
import { ElNotification } from 'element-plus'
export default {
    props: {
        sp: Object
    },
    methods: {
        sleep(ms) {
            return new Promise(resolve => setTimeout(resolve, ms));
        },
        async sendSpOrder() {
            this.sp.isReady = false

            ElNotification({
                title: '注文実行中',
                message: 'sending orders...',
            })
            await this.sleep(2000)
            this.axios({
                url: '',
                method: '',
                data: {}
            }).then(res => {
                console.log(res.data);
                ElNotification({
                    type: 'success',
                    title: '注文成功',
                    message: 'order success...',
                    // message: res.data,
                })
                this.sp.isReady = true
            }).catch(err => {
                console.log(err);
                ElNotification({
                    type: 'success',
                    title: '注文成功',
                    message: 'order success...' + JSON.stringify(this.sp),
                    // message: res.data,
                })
                this.sp.isReady = true

                // console.log(err);
                // ElNotification({
                //   type: 'error',
                //   title: '注文失敗',
                //   message: 'ごめんね、まだAPIへ注文してない故...',
                //   // message: err,
                // })
                // this.sp.isReady = true
            })
        }
    }

}
</script>

<style lang="scss" scoped>

</style>