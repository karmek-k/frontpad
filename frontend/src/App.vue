<script setup>
import AppHeader from './components/AppHeader.vue';
import ConnectionForm from './components/ConnectionForm.vue';
import CodeEditors from './components/CodeEditors.vue';

let ws = null;
</script>

<template>
  <AppHeader />
  <ConnectionForm :ws="ws" />
  <CodeEditors v-show="ws !== null" />
</template>
