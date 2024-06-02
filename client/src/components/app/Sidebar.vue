<template>
    <aside ref="sidebar" :class="{ 'w-3/4 z-50': isMobile, 'w-[250px]': !isMobile }"
        class="fixed top-0 left-0 h-full select-none bg-background border-r-2 border-gray-300 overflow-y-auto">
        <menu class="flex flex-col gap-2 p-4">
            <div class="text-xl font-bold"> People Matter </div>
            <hr class="border-b border-gray-300" />

            <div v-for="(section, sectionIndex) in menuSections" :key="sectionIndex">
                <div class="text-sm text-gray-500">{{ section.label }}</div>
                <div v-for="(item, itemIndex) in section.items" :key="itemIndex">
                    <div v-if="item.children" @click="toggleSubmenu(sectionIndex)"
                        class="flex justify-between items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                        <div class="flex gap-2">
                            <Icon :icon="item.icon" class="w-5 h-5" />
                            <div>{{ item.name }}</div>
                        </div>
                        <Icon
                            :icon="activeSubmenu[sectionIndex] ? 'material-symbols-light:keyboard-arrow-down' : 'material-symbols-light:keyboard-arrow-right'"
                            class="w-5 h-5" />
                    </div>
                    <div v-if="item.children" ref="submenus" class="flex flex-col gap-2 pl-2 overflow-hidden">
                        <div v-for="(child, childIndex) in item.children" :key="childIndex"
                            class="flex items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                            <Icon :icon="child.icon" class="w-5 h-5" />
                            <div>{{ child.name }}</div>
                        </div>
                    </div>
                    <div v-else
                        class="flex items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                        <Icon :icon="item.icon" class="w-5 h-5" />
                        <div>{{ item.name }}</div>
                    </div>
                </div>
                <hr class="border-b border-gray-300 my-1" v-if="sectionIndex < menuSections.length - 1" />
            </div>

            <div
                class="flex items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                <Icon icon="material-symbols:settings" class="w-5 h-5" />
                <div>Settings</div>
            </div>
            <div
                class="flex items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                <Icon icon="material-symbols:logout" class="w-5 h-5" />
                <div>Sign Out</div>
            </div>
        </menu>
    </aside>
</template>

<script setup>
import { ref, reactive, watch, onMounted, nextTick } from 'vue';
import { gsap } from 'gsap';
import { Icon } from '@iconify/vue';
import { defineProps } from 'vue';

const props = defineProps({
    showSidebar: Boolean,
    isMobile: Boolean
});

// Menu sections
const menuSections = ref([
    {
        label: 'Administrator',
        items: [
            { name: 'Dashboard', icon: 'material-symbols:dashboard', path: '/admin' },
            {
                name: 'Company',
                icon: 'mdi:office-building',
                children: [
                    { name: 'Profile', icon: 'ooui:view-details-ltr', path: '/admin/company' },
                    { name: 'Locations', icon: 'material-symbols:map', path: '/admin/locations' },
                    { name: 'Departments', icon: 'system-uicons:hierarchy', path: '/admin/departments' },
                    { name: 'Positions', icon: 'hugeicons:job-link', path: '/admin/positions' },
                ]
            }
        ]
    },
    {
        label: 'Location: {location}',
        items: [
            { name: 'Dashboard', icon: 'material-symbols:dashboard', path: '/manager' },
            {
                name: 'Location',
                icon: 'material-symbols:map',
                children: [
                    { name: 'Users', icon: 'ph:user-list', path: '/manager/users' },
                    { name: 'Leave', icon: 'material-symbols:sick-outline', path: '/manager/leave' },
                    { name: 'Salary', icon: 'material-symbols:payments', path: '/manager/salary' },
                    { name: 'Claims', icon: 'material-symbols:money-outline', path: '/manager/claims' },
                ]
            }
        ]
    },
    {
        label: 'User: {userId}',
        items: [
            { name: 'Dashboard', icon: 'material-symbols:dashboard', path: '/user' },
            {
                name: 'User',
                icon: 'material-symbols:frame-person-outline',
                children: [
                    { name: 'Profile', icon: 'fluent:slide-text-person-16-filled', path: '/user/profile' },
                    { name: 'Leave', icon: 'material-symbols:sick-outline', path: '/user/leave' },
                    { name: 'Salary', icon: 'material-symbols:payments', path: '/user/salary' },
                    { name: 'Claims', icon: 'material-symbols:money-outline', path: '/user/claims' },
                ]
            }
        ]
    }
]);

const activeSubmenu = reactive({});
const submenus = ref([]);

const toggleSubmenu = (sectionIndex) => {
    activeSubmenu[sectionIndex] = !activeSubmenu[sectionIndex];
};

watch(() => activeSubmenu, () => {
    submenus.value.forEach((submenu, index) => {
        gsap.to(submenu, { height: activeSubmenu[index] ? 'auto' : 0, duration: 0.5 });
    });
}, { deep: true });

// Sidebar animation
const sidebar = ref(null);

const animateSidebar = async () => {
    await nextTick();  // Ensure DOM updates are completed
    const sidebarWidth = sidebar.value.offsetWidth;
    gsap.to(sidebar.value, { x: props.showSidebar ? 0 : -sidebarWidth, duration: 0.5 });
};

watch(() => props.showSidebar, () => {
    animateSidebar();
});

onMounted(() => {
    animateSidebar();
    submenus.value.forEach(submenu => {
        gsap.set(submenu, { height: 0 });
    });
});
</script>
