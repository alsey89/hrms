<!-- MainComponent.vue -->
<template>
    <div class="w-full h-screen flex bg-background">
        <!-- button to toggle sidebar -->
        <Icon v-if="screenWidth < 768 && !showSidebar" @click="toggleSidebar" icon="material-symbols:menu-rounded"
            class="fixed top-3 left-3 w-6 h-6 text-primary z-10" />
        <!-- overlay to hide sidebar -->
        <div v-if="showSidebar && screenWidth < 768" @click="toggleSidebar"
            class="fixed inset-0 bg-black bg-opacity-40 z-40"></div>
        <!-- sidebar -->
        <Sidebar :showSidebar="showSidebar" :screenWidth="screenWidth" />
        <!-- main content -->
        <div :class="{ 'ml-[200px]': showSidebar && screenWidth >= 768, 'ml-0': !showSidebar || screenWidth < 768 }"
            class="w-full h-full overflow-y-scroll">
            <router-view></router-view>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount } from 'vue';
import { Sidebar } from "@/components/app";
import { Icon } from '@iconify/vue';

const showSidebar = ref(window.innerWidth >= 768);
const screenWidth = ref(window.innerWidth);
const toggleSidebar = () => {
    showSidebar.value = !showSidebar.value;
};

const handleResize = () => {
    screenWidth.value = window.innerWidth;
    if (screenWidth.value >= 768) {
        showSidebar.value = true;
    } else {
        showSidebar.value = false;
    }
};

onMounted(() => {
    window.addEventListener('resize', handleResize);
    handleResize();
});

onBeforeUnmount(() => {
    window.removeEventListener('resize', handleResize);
});
</script>