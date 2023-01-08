<script setup>
import ComponentsOrders from './components/ComponentsOrders.vue'
import ComponentsSpOrders from './components/ComponentsSpOrders.vue'
import ComponentsUser from './components/ComponentsUser.vue'
import ComponentsSetting from './components/ComponentsSetting.vue'
import ComponentsRules from './components/ComponentsRules.vue'
</script>

<script>
export default {
  data() {
    return {
      tabN: 'spOrder',
      user: {
        // Dummy Keys
        api: {
          key: 'a20241aab5a808ffa832a99f0be21aeb67824953bbe0a9eee1eaf10bce46eb8b',
          secret: 'a20241aab5a808ffa832a99f0be21aeb67824953bbe0a9eee1eaf10bce46eb8b.username.e829e44ffb65154eb7206db046700173122970e188d3018e1faa8047deec2e9d',
        }
      },
      setting: {
        token: 'e829e44ffb65154eb7206db046700173122970e188d3018e1faa8047deec2e9d',
        types: [],
        tif: 'GTC'
      },
      order: {
        step: 5,
        isReady: true,
        exchange: 1,
        code: '0000',
        price: 25_700,
        size: 1,
        type: 'limit'
      },
      sp: {
        step: 5,
        isReady: true,
        exchange: 1,
        code: '0000',
        price: 25_700,
        size: 1,
        lossTriggerPrice: 25_000,
        lossFixedPrice: 25_000,
        orderType: 'limit',
        triggerOrderType: 'market',
      }
    }
  },
  methods: {
    sleep(ms) {
      return new Promise(resolve => setTimeout(resolve, ms));
    }
  }
}
</script>

<template>
  <header>
    <!-- <img alt="Vue logo" class="logo" src="./assets/logo.svg" width="125" height="125" /> -->
  </header>

  <main>
    <el-tabs v-model="tabN" class="tab-style">
      <el-tab-pane label="Order" name="order">
        <ComponentsOrders :order="order" />
      </el-tab-pane>
      <el-tab-pane label="SP-Orders" name="spOrder">
        <ComponentsSpOrders :sp="sp" />
      </el-tab-pane>
      <el-tab-pane label="User" name="user">
        <ComponentsUser :user="user" />
      </el-tab-pane>
      <el-tab-pane label="Setting" name="setting">
        <ComponentsSetting :setting="setting" />
      </el-tab-pane>
      <el-tab-pane label="Rules" name="rules">
        <ComponentsRules />
      </el-tab-pane>
    </el-tabs>

  </main>
</template>

<style scoped>
header {
  line-height: 1.5;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

.el-row {
  margin-bottom: 1rem;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }
}
</style>
