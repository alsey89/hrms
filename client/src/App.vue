<template>
  <router-view />
</template>

<script setup>
import { onMounted } from 'vue';
import { useUserStore } from '@/stores/User';
import posthog from 'posthog-js';

const userStore = useUserStore();

onMounted(() => {
  if (userStore.isAuthenticated) {
    posthog.identify(authStore.user.id);
    authStore.setIdentified(true); // Set a flag in your auth store to mark as identified
  }
});


</script>