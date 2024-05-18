<template>
    <div class="w-full h-screen flex bg-background">
        <!-- overlay to hide sidebar -->
        <div v-if="showSidebar && screenWidth < 768" @click="toggleSidebar"
            class="fixed top-0 left-0 w-full h-full bg-black bg-opacity-50 z-40"></div>
        <!-- navigation sidebar -->
        <aside ref="sidebar"
            :class="{ 'w-3/4 sm:w-[120px]': screenWidth < 768, 'w-[120px]': screenWidth >= 768, 'hidden': !showSidebar && screenWidth < 768 }"
            class="fixed top-0 left-0 h-full bg-primary z-50 md:z-0">
            <menu class="flex flex-col gap-4 p-4">
                <div class="flex items-center gap-2">
                    <Icon icon="mdi-light:home" class="w-5 h-5" />
                    <div>Home</div>
                </div>
                <div class="flex items-center gap-2">
                    <Icon icon="mdi-light:home" class="w-5 h-5" />
                    <div>Home</div>
                </div>
                <div class="flex items-center gap-2">
                    <Icon icon="mdi-light:home" class="w-5 h-5" />
                    <div>Home</div>
                </div>
            </menu>
        </aside>
        <div :class="{ 'ml-[120px]': showSidebar && screenWidth >= 768, 'ml-0': !showSidebar || screenWidth < 768 }"
            class="w-full h-full overflow-y-scroll">
            <Icon v-if="!showSidebar" @click="toggleSidebar" icon="material-symbols:menu-rounded"
                class="fixed top-2 right-2 w-6 h-6 text-primary z-10" />
            <router-view></router-view>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount, watch } from 'vue';
import { gsap } from 'gsap';
import { Icon } from '@iconify/vue';


const showSidebar = ref(window.innerWidth >= 480);
const sidebar = ref(null);
const screenWidth = ref(window.innerWidth); // Reactive variable for screen width

const toggleSidebar = () => {
    showSidebar.value = !showSidebar.value;
};

// Function to handle window resize and update screen width
const handleResize = () => {
    screenWidth.value = window.innerWidth;
    if (screenWidth.value < 480) {
        showSidebar.value = false;
    }
};

// Watch the showSidebar ref to animate the sidebar
watch(showSidebar, (newVal) => {
    if (newVal) {
        gsap.to(sidebar.value, { x: 0, duration: 0.5 });
    } else {
        gsap.to(sidebar.value, { x: -sidebar.value.offsetWidth, duration: 0.5 });
    }
});

onMounted(() => {
    window.addEventListener('resize', handleResize); // Add resize event listener
    handleResize(); // Initial check

    // Set initial position of the sidebar
    if (showSidebar.value) {
        gsap.set(sidebar.value, { x: 0 });
    } else {
        gsap.set(sidebar.value, { x: -sidebar.value.offsetWidth });
    }
});

onBeforeUnmount(() => {
    window.removeEventListener('resize', handleResize); // Clean up event listener
});
</script>
