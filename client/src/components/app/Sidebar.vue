<template>
    <aside ref="sidebar"
        :class="{ 'w-3/4 z-50': isMobile, 'w-[250px]': !isMobile }"
        class="fixed top-0 left-0 h-full select-none bg-background border-r-2 border-gray-300 overflow-y-auto">
        <menu class="flex flex-col gap-2 p-4">
            <div class="text-xl font-bold"> People Matter </div>
            <hr class="border-b border-gray-300" />
            <!-- !ADMIN -->
            <div class="text-sm text-gray-500">Administrator</div>
            <div class="flex items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                <Icon icon="material-symbols:dashboard" class="w-5 h-5" />
                <div>Dashboard</div>
            </div>
            <!-- company -->
            <div @click="toggleCompanySubmenu" class="flex justify-between items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                <div class="flex gap-2">
                    <Icon icon="mdi:office-building" class="w-5 h-5" />
                    <div>Company</div>
                </div>
                <Icon :icon="showCompanySubmenu ? 'material-symbols-light:keyboard-arrow-down' : 'material-symbols-light:keyboard-arrow-right'"
        class="w-5 h-5" />
            </div>
            <!-- company submenu -->
            <div ref="companySubmenu" class="flex flex-col gap-2 pl-2 overflow-hidden">
                <div class="flex items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                    <Icon icon="ooui:view-details-ltr" class="w-5 h-5" />
                    <div>Details</div>
                </div>
                <div class="flex items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                    <Icon icon="material-symbols:map" class="w-5 h-5" />
                    <div>Locations</div>
                </div>
                <div class="flex items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                    <Icon icon="system-uicons:hierarchy" class="w-5 h-5" />
                    <div>Departments</div>
                </div>
                <div class="flex items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                    <Icon icon="hugeicons:job-link" class="w-5 h-5" />
                    <div>Positions</div>
                </div>
            </div>
            <hr class="border-b border-gray-300" />
            <!-- !MANAGER -->
            <div class="text-sm text-gray-500"> Location: {location} </div>
            <div class="flex items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                <Icon icon="material-symbols:dashboard" class="w-5 h-5" />
                <div>Dashboard</div>
            </div>
            <div @click="toggleManagerSubmenu" class="flex justify-between items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                <div class="flex gap-2">
                    <Icon icon="material-symbols:map" class="w-5 h-5" />
                    <div> Location </div>
                </div>
                <Icon :icon="showManagerSubmenu ? 'material-symbols-light:keyboard-arrow-down' : 'material-symbols-light:keyboard-arrow-right'"
        class="w-5 h-5" />
            </div>
            <!-- manager submenu -->
            <div ref="managerSubmenu" class="flex flex-col gap-2 pl-2 overflow-hidden">
                <div class="flex items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                    <Icon icon="ph:user-list" class="w-5 h-5" />
                    <div>Users</div>
                </div>
                <div class="flex items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                    <Icon icon="material-symbols:sick-outline" class="w-5 h-5" />
                    <div>Leave</div>
                </div>
                <div class="flex items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                    <Icon icon="material-symbols:payments" class="w-5 h-5" />
                    <div>Salary</div>
                </div>
                <div class="flex items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                    <Icon icon="material-symbols:money-outline" class="w-5 h-5" />
                    <div>Claims</div>
                </div>
            </div>
            <hr class="border-b border-gray-300" />
            <!-- !User -->
            <div class="text-sm text-gray-500"> User: {userId} </div>
            <div class="flex items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                <Icon icon="material-symbols:dashboard" class="w-5 h-5" />
                <div>Dashboard</div>
            </div>
            <div @click="toggleUserSubmenu" class="flex justify-between items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                <div class="flex gap-2">
                    <Icon icon="material-symbols:frame-person-outline" class="w-5 h-5" />
                    <div> User </div>
                </div>
                <Icon :icon="showUserSubmenu ? 'material-symbols-light:keyboard-arrow-down' : 'material-symbols-light:keyboard-arrow-right'"
        class="w-5 h-5" />
            </div>
            <!-- manager submenu -->
            <div ref="userSubmenu" class="flex flex-col gap-2 pl-2 overflow-hidden">
                <div class="flex items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                    <Icon icon="fluent:slide-text-person-16-filled" class="w-5 h-5" />
                    <div> Profile </div>
                </div>
                <div class="flex items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                    <Icon icon="material-symbols:sick-outline" class="w-5 h-5" />
                    <div> Leave </div>
                </div>
                <div class="flex items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                    <Icon icon="material-symbols:payments" class="w-5 h-5" />
                    <div> Salary </div>
                </div>
                <div class="flex items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                    <Icon icon="material-symbols:money-outline" class="w-5 h-5" />
                    <div> Claims </div>
                </div>
            </div>
            <hr class="border-b border-gray-300" />
            <div class="flex items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                <Icon icon="material-symbols:settings" class="w-5 h-5" />
                <div>Settings</div>
            </div>
            <div class="flex items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                <Icon icon="material-symbols:logout" class="w-5 h-5" />
                <div>Sign Out</div>
            </div>
        </menu>
    </aside>
  </template>
  
  <script setup>
  import { ref, watch, onMounted, nextTick } from 'vue';
  import { gsap } from 'gsap';
  import { Icon } from '@iconify/vue';
  import { defineProps } from 'vue';
  
  const props = defineProps({
    showSidebar: Boolean,
    isMobile: Boolean
  });
  
  // ------------------ submenu animation ------------------
  
  const showCompanySubmenu = ref(false);
  const companySubmenu = ref(null);
  const showUserSubmenu = ref(false);
  const userSubmenu = ref(null);
  const showManagerSubmenu = ref(false);
  const managerSubmenu = ref(null);
  
  const toggleCompanySubmenu = () => {
    showCompanySubmenu.value = !showCompanySubmenu.value;
  };

    const toggleManagerSubmenu = () => {
        showManagerSubmenu.value = !showManagerSubmenu.value;
    };

    const toggleUserSubmenu = () => {
    showUserSubmenu.value = !showUserSubmenu.value;
  };
  
  watch(showCompanySubmenu, (newVal) => {
    gsap.to(companySubmenu.value, { height: newVal ? 'auto' : 0, duration: 0.5 });
  });

    watch(showManagerSubmenu, (newVal) => {
        gsap.to(managerSubmenu.value, { height: newVal ? 'auto' : 0, duration: 0.5 });
    });
  
  watch(showUserSubmenu, (newVal) => {
    gsap.to(userSubmenu.value, { height: newVal ? 'auto' : 0, duration: 0.5 });
  });
  
  // ------------------ sidebar animation ------------------
  
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
    gsap.set(companySubmenu.value, { height: 0 });
    gsap.set(managerSubmenu.value, { height: 0 });
    gsap.set(userSubmenu.value, { height: 0 });
  });
  </script>
  