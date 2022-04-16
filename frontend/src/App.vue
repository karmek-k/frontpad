<script setup>
import AppHeader from './components/AppHeader.vue';
import ConnectionForm from './components/ConnectionForm.vue';
import CodeEditors from './components/CodeEditors.vue';
import { reactive } from 'vue';

const state = reactive({
  ws: null,
  isError: false,
  sessionId: '',
  isLoading: false
});

async function createNewSession(apiUrl) {
  const res = await fetch(apiUrl, {
    method: 'POST'
  });
  const json = await res.json();

  return json.id;
}

function handleNewSession() {
  state.isError = false;
  state.isLoading = true;

  createNewSession('http://localhost:8000/api/session')
    .then(id => (state.sessionId = id))
    .catch(() => (state.isError = true))
    .finally(() => (state.isLoading = false));
}

function handleConnectWs() {
  state.ws = connectWs('ws://localhost:8000/ws');
}

function handleDisconnectWs() {
  disconnectWs(false);
}

function connectWs(url) {
  if (state.ws !== null) {
    return null;
  }

  const ws = new WebSocket(url);
  ws.addEventListener('error', () => {
    state.isError = true;
  });

  ws.addEventListener('open', () => {
    state.isError = false;
  });

  ws.addEventListener('close', () => {
    disconnectWs(true);
  });

  return ws;
}

function disconnectWs(fromWsCloseEvent) {
  if (state.ws === null) {
    return;
  }

  if (!fromWsCloseEvent) {
    state.ws.close();
  }

  state.ws = null;
}
</script>

<template>
  <AppHeader />
  <ConnectionForm
    :is-error="state.isError"
    :is-connected="isConnected"
    :disable-create-session="state.sessionId.length > 0"
    @session-create="handleNewSession"
    @ws-connect="handleConnectWs"
    @ws-disconnect="handleDisconnectWs"
  />
  <CodeEditors v-show="isConnected" />
</template>
