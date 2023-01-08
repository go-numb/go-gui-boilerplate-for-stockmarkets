<script setup>
import WelcomeItem from './WelcomeItem.vue'
import { Document, User, SetUp, CircleCheck, ChatDotRound } from '@element-plus/icons-vue'
</script>

<template>
  <el-row>
    <el-col :span="22" :offset="1">
      <WelcomeItem>
        <template #heading>
          <el-icon>
            <Document />
          </el-icon>
          ドキュメント
        </template>

        使用例やルールなどを記載するつもり
      </WelcomeItem>

      <WelcomeItem>
        <template #heading>
          <el-icon>
            <SetUp />
          </el-icon>
          開発ツール
        </template>

        開発ツールなどなど

        <ul>
          <li>
            Background: Golang
          </li>
          <li>
            GUI: Electon
          </li>
          <li>
            Frontend: Vue3 SPA
          </li>
          <li>
            Design: Element-Plus
          </li>
          <li>
            Requests: Axios
          </li>
        </ul>
      </WelcomeItem>

      <WelcomeItem>
        <template #heading>
          <el-icon>
            <CircleCheck />
          </el-icon>
          対象
        </template>

        <ul>
          <li>auカブドットコム証券</li>
          <li>J-Quants</li>
          <li>AlpacaAPI</li>
        </ul>
      </WelcomeItem>

      <WelcomeItem>
        <template #heading>
          <el-icon>
            <ChatDotRound />
          </el-icon>
          コミュニティ
        </template>

        Discord Community
      </WelcomeItem>

      <WelcomeItem>
        <template #heading>
          <el-icon>
            <User />
          </el-icon>
          Support Me
        </template>
        Twitter:
        <a href="https://twitter.com/_numbP" target="_blank">
          <el-avatar :size="40" src="https://pbs.twimg.com/profile_images/1121699376835878913/CPVPitJ-_400x400.png"
            fit="cover" />
        </a>.
      </WelcomeItem>
    </el-col>
  </el-row>
</template>
