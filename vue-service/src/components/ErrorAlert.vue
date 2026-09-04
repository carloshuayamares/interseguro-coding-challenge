<script setup lang="ts">
defineProps<{ message: string }>();
defineEmits<{ dismiss: [] }>();
</script>

<template>
  <v-alert type="error" variant="tonal" closable class="error-alert" @click:close="$emit('dismiss')">
    {{ message }}
  </v-alert>
</template>
