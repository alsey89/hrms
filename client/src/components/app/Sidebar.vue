<template>
    <aside ref="sidebar"
        :class="['fixed top-0 left-0 h-full flex flex-col px-2 gap-2 border-r-2 border-accent bg-background select-none', isMobile ? 'w-3/4' : 'w-[250px]', 'z-50']">
        <h2 class="text-lg font-semibold px-4 py-2"> People Matter </h2>
        <div v-for="section in menuSections" :key="section.label">
            <h2 class="text-xs text-gray-500 p-2 border-t">{{ section.label }}</h2>
            <ul class="py-2">
                <li v-for="item in section.items" :key="item.name" class="group">
                    <div v-if="item.children" @click="toggle(item, $event)"
                        class="flex items-center p-2 border-2 border-background rounded-md cursor-pointer hover:border-accent">
                        <Icon :icon="item.icon" class="text-xl" />
                        <span class="ml-3">{{ item.name }}</span>
                        <span class="ml-auto transition-transform duration-200"
                            :class="item.isOpen ? 'rotate-180' : ''">
                            <Icon icon="material-symbols:expand-more" />
                        </span>
                    </div>
                    <router-link v-else :to="item.path"
                        class="flex items-center p-2 border-2 border-background rounded-md cursor-pointer hover:border-accent">
                        <Icon :icon="item.icon" class="text-xl" />
                        <span class="ml-3">{{ item.name }}</span>
                    </router-link>
                    <ul ref="submenu" :class="{ 'hidden': !item.isOpen }" class="ml-8 overflow-hidden">
                        <li v-for="child in item.children" :key="child.name"
                            class="p-2 border-2 border-background rounded-md cursor-pointer hover:border-accent">
                            <router-link :to="child.path" class="flex items-center">
                                <Icon :icon="child.icon" class="text-xl" />
                                <span class="ml-2">{{ child.name }}</span>
                            </router-link>
                        </li>
                    </ul>
                </li>
            </ul>
        </div>
    </aside>
</template>

<script setup>
import { ref, watch, onMounted, nextTick } from 'vue';
import { gsap } from 'gsap';
import { Icon } from '@iconify/vue';

const props = defineProps({
    showSidebar: Boolean,
    isMobile: Boolean
});

const menuSections = ref([
    {
        label: 'Administrator',
        items: [
            { name: 'Dashboard', icon: 'material-symbols:dashboard', path: '/admin' },
            {
                name: 'Company',
                icon: 'mdi:office-building',
                isOpen: false,
                children: [
                    { name: 'Profile', icon: 'ooui:view-details-ltr', path: '/admin/company' },
                    { name: 'Locations', icon: 'material-symbols:map', path: '/admin/locations' },
                    { name: 'Departments', icon: 'system-uicons:hierarchy', path: '/admin/departments' },
                    { name: 'Positions', icon: 'hugeicons:job-link', path: '/admin/positions' },
                ]
            },
            {
                name: 'Policies',
                icon: 'material-symbols:policy',
                isOpen: false,
                children: [
                    { name: 'Attendance', icon: 'material-symbols:punch-clock', path: '/admin/company' },
                    { name: 'Leave', icon: 'material-symbols:sick', path: '/admin/locations' },
                    { name: 'Salary', icon: 'material-symbols:payments', path: '/admin/departments' },
                    { name: 'Claims', icon: 'material-symbols:money', path: '/admin/positions' },
                    { name: 'Compliance', icon: 'octicon:law', path: '/admin/positions' },
                ]
            },
        ]
    },
    {
        label: 'Location: {location}',
        items: [
            { name: 'Dashboard', icon: 'material-symbols:dashboard', path: '/manager' },
            {
                name: 'Location',
                icon: 'material-symbols:map',
                isOpen: false,
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
                isOpen: false,
                children: [
                    { name: 'Profile', icon: 'fluent:slide-text-person-16-filled', path: '/user/profile' },
                    { name: 'Leave', icon: 'material-symbols:sick-outline', path: '/user/leave' },
                    { name: 'Salary', icon: 'material-symbols:payments', path: '/user/salary' },
                    { name: 'Claims', icon: 'material-symbols:money-outline', path: '/user/claims' },
                ]
            }
        ]
    },
    {
        label: 'Settings',
        items: [
            { name: 'Settings', icon: 'mdi:cog', path: '/settings' },
            { name: 'Help', icon: 'mdi:help-circle', path: '/help' },
            { name: 'Logout', icon: 'mdi:logout', path: '/auth/signout' },
        ]
    }
]);

const sidebar = ref(null);

const animateSidebar = async () => {
    await nextTick();  // Ensure DOM updates are completed
    const sidebarWidth = sidebar.value.offsetWidth;
    gsap.to(sidebar.value, { x: props.showSidebar ? 0 : -sidebarWidth, duration: 0.5 });
};

const toggle = (item, event) => {
    if (item.children) {
        item.isOpen = !item.isOpen;
        const submenu = event.currentTarget.nextElementSibling;
        if (item.isOpen) {
            gsap.fromTo(submenu, { height: 0 }, { height: 'auto', duration: 0.5 });
        } else {
            gsap.to(submenu, { height: 0, duration: 0.5 });
        }
    }
};

watch(() => props.showSidebar, () => {
    animateSidebar();
});

onMounted(() => {
    animateSidebar();
});
</script>

<style scoped>
.hidden {
    display: none;
}
</style>
