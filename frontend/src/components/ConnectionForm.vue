<script setup>
import { computed, reactive } from 'vue';

defineProps(['ws']);

const state = reactive({
  sessionId: ''
});

const inputHasText = computed(() => {
  return state.sessionId.length > 0;
});
</script>

<template>
  <input
    type="text"
    v-model="state.sessionId"
    placeholder="Leave empty to create"
  />
  <button :disabled="inputHasText">Create a session</button>
  <button :disabled="ws !== null || !inputHasText">Connect</button>
  <button :disabled="ws === null">Disconnect</button>
</template>
