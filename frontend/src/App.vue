<script setup>
import AppHeader from './components/AppHeader.vue';
import ConnectionForm from './components/ConnectionForm.vue';
import CodeEditors from './components/CodeEditors.vue';
import { onMounted, onUnmounted, reactive } from 'vue';

const state = reactive({
  ws: null,
  sessionId: ''
});

onMounted(() => {
  const ws = new WebSocket('ws://localhost:8000/ws');
  ws.addEventListener('error', () => {
    alert('WebSocket connection error!');
  });

  ws.addEventListener('message', e => {
    console.log(e.data);
  });

  state.ws = ws;
});

onUnmounted(() => {
  state.ws.close();
  state.ws = null;
});
</script>

<template>
  <AppHeader />
  <ConnectionForm
    @session-edit="id => (state.sessionId = id)"
    @session-connect=""
  />
  <CodeEditors v-show="false" />
</template>
