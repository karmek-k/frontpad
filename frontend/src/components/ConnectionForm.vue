<script setup>
import { computed, reactive } from 'vue';

defineEmits(['sessionCreate', 'sessionConnect']);

const state = reactive({
  inputText: '',
  isLoading: false,
  isError: false
});

const inputHasText = computed(() => {
  return state.inputText.length > 0;
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
    .then(sessId => (state.inputText = sessId))
    .catch(() => (state.isError = true))
    .finally(() => (state.isLoading = false));
}
</script>

<template>
  <p v-show="state.isError">An error has occurred</p>
  <input
    type="text"
    v-model="state.inputText"
    placeholder="Leave empty to create"
    @input="$emit('sessionEdit', state.inputText)"
  />
  <button @click="handleNewSession" :disabled="inputHasText">
    Create a session
  </button>
  <button @click="" :disabled="!inputHasText">Connect</button>
</template>

<style scoped>
input {
  width: 200px;
}
</style>
