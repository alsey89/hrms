<template>
</template>

<script setup>
import { useUserStore } from '@/stores/User';
import { onBeforeMount } from 'vue';
import { useRouter } from 'vue-router';
const router = useRouter();

onBeforeMount(() => {
    const userStore = useUserStore();
    userStore.signout(router);
});
</script>