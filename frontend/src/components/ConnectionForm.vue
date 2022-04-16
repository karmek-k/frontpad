<script setup>
import { computed, reactive } from 'vue';

defineEmits(['sessionCreate', 'wsConnect', 'wsDisconnect']);

const props = defineProps({
  isConnected: {
    type: Boolean
  },
  isError: {
    type: Boolean
  },
  disableCreateSession: {
    type: Boolean,
    default: false
  }
});

const state = reactive({
  inputText: ''
});
</script>

<template>
  <p v-show="props.isError">An error has occurred</p>
  <input
    type="text"
    v-model="state.inputText"
    placeholder="Leave empty to create"
    :disabled="props.isConnected"
  />
  <button
    @click="$emit('sessionCreate')"
    :disabled="props.disableCreateSession"
  >
    Create a session
  </button>
  <button
    @click="$emit('wsConnect')"
    :disabled="props.isConnected || !props.disableCreateSession"
  >
    Connect
  </button>
  <button @click="$emit('wsDisconnect')" :disabled="!props.isConnected">
    Disconnect
  </button>
</template>

<style scoped>
input {
  width: 200px;
}
</style>
