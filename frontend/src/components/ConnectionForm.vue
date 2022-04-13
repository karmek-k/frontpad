<script setup>
import { computed, reactive } from 'vue';

const state = reactive({
  sessionId: '',
  isError: false,
  ws: null
});

const inputHasText = computed(() => {
  return state.sessionId.length > 0;
});

const isConnected = computed(() => {
  return state.ws !== null;
});

async function newSession() {
  const res = await fetch('http://localhost:8000/api/session', {
    method: 'POST'
  });
  const json = await res.json();

  return json.id;
}

function connectWs(url) {
  if (!inputHasText && isConnected) {
    return null;
  }

  const ws = new WebSocket(url);
  ws.addEventListener('error', () => {
    state.isError = true;
  });

  ws.addEventListener('open', () => {
    state.isError = false;
  });

  return ws;
}

function disconnectWs() {
  if (!isConnected) {
    return;
  }

  state.ws.close();
  state.ws = null;
}
</script>

<template>
  <p v-show="state.isError">An error has occurred</p>
  <input
    type="text"
    v-model="state.sessionId"
    placeholder="Leave empty to create"
    :disabled="isConnected"
  />
  <button
    @click="
      () =>
        newSession()
          .then(id => (state.sessionId = id))
          .catch(() => (state.isError = true))
    "
    :disabled="inputHasText"
  >
    Create a session
  </button>
  <button
    @click="() => (state.ws = connectWs('ws://localhost:8000/ws'))"
    :disabled="isConnected || !inputHasText"
  >
    Connect
  </button>
  <button @click="disconnectWs" :disabled="!isConnected">Disconnect</button>
</template>

<style scoped>
input {
  width: 200px;
}
</style>
