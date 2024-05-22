<template>
    <div class="w-full h-screen flex bg-background">
        <!-- overlay to hide sidebar -->
        <div v-if="showSidebar && screenWidth < 768" @click="toggleSidebar"
            class="fixed top-0 left-0 w-full h-full bg-black bg-opacity-50 z-40"></div>
        <!-- navigation sidebar -->
        <aside ref="sidebar"
            :class="{ 'w-3/4 sm:w-[200px]': screenWidth < 768, 'w-[200px]': screenWidth >= 768, 'hidden': !showSidebar && screenWidth < 768 }"
            class="fixed top-0 left-0 h-full select-none bg-card z-50 md:z-0" overflow-y-scroll>
            <menu class="flex flex-col gap-2 p-4">
                <div>
                    <div class="text-xl font-bold"> People Matter </div>
                    <div class="text-sm text-gray-500">Navigation</div>
                </div>
                <hr class="border-b border-gray-300" />
                <div class="flex items-center gap-2 p-2 hover:bg-accent hover:cursor-pointer">
                    <Icon icon="mdi-light:home" class="w-5 h-5" />
                    <div>Dashboard</div>
                </div>
                <!-- company -->
                <div class="flex justify-between items-center gap-2 p-2 hover:bg-accent hover:cursor-pointer">
                    <div class="flex gap-2">
                        <Icon icon="mdi-light:home" class="w-5 h-5" />
                        <div>Company</div>
                    </div>
                    <Icon @click="toggleCompanySubmenu" icon="material-symbols-light:keyboard-arrow-down"
                        class="w-5 h-5" />
                </div>
                <!-- company submenu -->
                <div v-show="showCompanySubmenu" ref="companySubmenu" class="flex flex-col gap-2 pl-4 overflow-hidden">
                    <div class="flex items-center gap-2 p-2 hover:bg-accent hover:cursor-pointer">
                        <Icon icon="mdi-light:home" class="w-5 h-5" />
                        <div>Company</div>
                    </div>
                    <div class="flex items-center gap-2 p-2 hover:bg-accent hover:cursor-pointer">
                        <Icon icon="mdi-light:home" class="w-5 h-5" />
                        <div>Departments</div>
                    </div>
                    <div class="flex items-center gap-2 p-2 hover:bg-accent hover:cursor-pointer">
                        <Icon icon="mdi-light:home" class="w-5 h-5" />
                        <div>Roles</div>
                    </div>
                    <div class="flex items-center gap-2 p-2 hover:bg-accent hover:cursor-pointer">
                        <Icon icon="mdi-light:home" class="w-5 h-5" />
                        <div>Employees</div>
                    </div>
                </div>
                <!-- user -->
                <div class="flex justify-between items-center gap-2 p-2 hover:bg-accent hover:cursor-pointer">
                    <div class="flex gap-2">
                        <Icon icon="mdi-light:home" class="w-5 h-5" />
                        <div>User</div>
                    </div>
                    <Icon @click="toggleUserSubmenu" icon="material-symbols-light:keyboard-arrow-down"
                        class="w-5 h-5" />
                </div>
                <!-- user submenu -->
                <div v-show="showUserSubmenu" ref="userSubmenu" class="flex flex-col gap-2 pl-4 overflow-hidden">
                    <div class="flex items-center gap-2 p-2 hover:bg-accent hover:cursor-pointer">
                        <Icon icon="mdi-light:home" class="w-5 h-5" />
                        <div>Profile</div>
                    </div>
                    <div class="flex items-center gap-2 p-2 hover:bg-accent hover:cursor-pointer">
                        <Icon icon="mdi-light:home" class="w-5 h-5" />
                        <div>Settings</div>
                    </div>
                    <div class="flex items-center gap-2 p-2 hover:bg-accent hover:cursor-pointer">
                        <Icon icon="mdi-light:home" class="w-5 h-5" />
                        <div>Profile</div>
                    </div>
                    <div class="flex items-center gap-2 p-2 hover:bg-accent hover:cursor-pointer">
                        <Icon icon="mdi-light:home" class="w-5 h-5" />
                        <div>Settings</div>
                    </div>
                </div>
                <div class="flex items-center gap-2 p-2 hover:bg-accent hover:cursor-pointer">
                    <Icon icon="mdi-light:home" class="w-5 h-5" />
                    <div>Settings</div>
                </div>
            </menu>
        </aside>
        <div :class="{ 'ml-[200px]': showSidebar && screenWidth >= 768, 'ml-0': !showSidebar || screenWidth < 768 }"
            class="w-full h-full overflow-y-scroll">
            <Icon v-if="screenWidth < 768 && !showSidebar" @click="toggleSidebar" icon="material-symbols:menu-rounded"
                class="fixed top-3 right-3 w-6 h-6 text-primary z-10" />
            <router-view></router-view>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount, watch } from 'vue';
import { gsap } from 'gsap';
import { Icon } from '@iconify/vue';

const showSidebar = ref(window.innerWidth >= 768);
const sidebar = ref(null);
const screenWidth = ref(window.innerWidth); // Reactive variable for screen width
const showCompanySubmenu = ref(true);
const companySubmenu = ref(null);
const showUserSubmenu = ref(false);
const userSubmenu = ref(null);

const toggleSidebar = () => {
    showSidebar.value = !showSidebar.value;
};

const toggleCompanySubmenu = () => {
    showCompanySubmenu.value = !showCompanySubmenu.value;
};

const toggleUserSubmenu = () => {
    showUserSubmenu.value = !showUserSubmenu.value;
};

// Function to handle window resize and update screen width
const handleResize = () => {
    screenWidth.value = window.innerWidth;
    if (screenWidth.value >= 768) {
        showSidebar.value = true; // Show sidebar by default on desktop
    } else {
        showSidebar.value = false; // Hide sidebar by default on mobile
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

// Watch the showCompanySubmenu ref to animate the submenu
watch(showCompanySubmenu, (newVal) => {
    if (newVal) {
        gsap.to(companySubmenu.value, { height: 'auto', duration: 0.5 });
    } else {
        gsap.to(companySubmenu.value, { height: 0, duration: 0.5 });
    }
});

// Watch the showUserSubmenu ref to animate the submenu
watch(showUserSubmenu, (newVal) => {
    if (newVal) {
        gsap.to(userSubmenu.value, { height: 'auto', duration: 0.5 });
    } else {
        gsap.to(userSubmenu.value, { height: 0, duration: 0.5 });
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
    // Set initial height of the submenus
    gsap.set(companySubmenu.value, { height: 'auto' });
    gsap.set(userSubmenu.value, { height: 0 });
});

onBeforeUnmount(() => {
    window.removeEventListener('resize', handleResize); // Clean up event listener
});
</script>

<style scoped>
/* Add any necessary styles here */
</style>
