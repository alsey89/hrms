<!-- Sidebar.vue -->
<template>
  <aside ref="sidebar"
      :class="{ 'w-3/4 sm:w-[200px]': screenWidth < 768, 'w-[200px]': screenWidth >= 768, 'hidden': !showSidebar && screenWidth < 768 }"
      class="fixed top-0 left-0 h-full select-none bg-background border-r-2 border-gray-300 z-50 md:z-0">
      <menu class="flex flex-col gap-2 p-4">
          <div>
              <div class="text-xl font-bold"> People Matter </div>
              <div class="text-sm text-gray-500">Navigation</div>
          </div>
          <hr class="border-b border-gray-300" />
          <div class="flex items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
              <Icon icon="mdi-light:home" class="w-5 h-5" />
              <div>Dashboard</div>
          </div>
          <!-- company -->
          <div @click="toggleCompanySubmenu"  class="flex justify-between items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
              <div class="flex gap-2">
                  <Icon icon="mdi-light:home" class="w-5 h-5" />
                  <div>Company</div>
              </div>
              <Icon :icon="showCompanySubmenu ? 'material-symbols-light:keyboard-arrow-down' : 'material-symbols-light:keyboard-arrow-right'"
      class="w-5 h-5" />
          </div>
          <!-- company submenu -->
          <div ref="companySubmenu" class="flex flex-col gap-2 pl-4 overflow-hidden">
              <div class="flex items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                  <Icon icon="mdi-light:home" class="w-5 h-5" />
                  <div>Company</div>
              </div>
              <div class="flex items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                  <Icon icon="mdi-light:home" class="w-5 h-5" />
                  <div>Locations</div>
              </div>
              <div class="flex items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                  <Icon icon="mdi-light:home" class="w-5 h-5" />
                  <div>Departments</div>
              </div>
              <div class="flex items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                  <Icon icon="mdi-light:home" class="w-5 h-5" />
                  <div>Positions</div>
              </div>
          </div>
          <!-- user -->
          <div @click="toggleUserSubmenu" class="flex justify-between items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
              <div class="flex gap-2">
                  <Icon icon="mdi-light:home" class="w-5 h-5" />
                  <div>Users</div>
              </div>
              <Icon :icon="showCompanySubmenu ? 'material-symbols-light:keyboard-arrow-down' : 'material-symbols-light:keyboard-arrow-right'"
      class="w-5 h-5" />
          </div>
          <!-- user submenu -->
          <div ref="userSubmenu" class="flex flex-col gap-2 pl-4 overflow-hidden">
              <div class="flex items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                  <Icon icon="mdi-light:home" class="w-5 h-5" />
                  <div>Users</div>
              </div>
              <div class="flex items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                  <Icon icon="mdi-light:home" class="w-5 h-5" />
                  <div>Leave</div>
              </div>
              <div class="flex items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                  <Icon icon="mdi-light:home" class="w-5 h-5" />
                  <div>Salary</div>
              </div>
              <div class="flex items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
                  <Icon icon="mdi-light:home" class="w-5 h-5" />
                  <div>Claims</div>
              </div>
          </div>
          <div class="flex items-center gap-2 p-2 border-2 border-background rounded-md hover:border-accent hover:cursor-pointer">
              <Icon icon="mdi-light:home" class="w-5 h-5" />
              <div>Settings</div>
          </div>
      </menu>
  </aside>
</template>

<script setup>
import { ref, watch, onMounted, onBeforeUnmount } from 'vue';
import { gsap } from 'gsap';
import { Icon } from '@iconify/vue';
import { defineProps, defineEmits } from 'vue';

const props = defineProps({
  showSidebar: Boolean,
  screenWidth: Number
});

// ------------------ submenu animation ------------------

const showCompanySubmenu = ref(false);
const companySubmenu = ref(null);
const showUserSubmenu = ref(false);
const userSubmenu = ref(null);

const toggleCompanySubmenu = () => {
  showCompanySubmenu.value = !showCompanySubmenu.value;
};

const toggleUserSubmenu = () => {
  showUserSubmenu.value = !showUserSubmenu.value;
};

watch(showCompanySubmenu, (newVal) => {
  if (newVal) {
      gsap.to(companySubmenu.value, { height: 'auto', duration: 0.5 });
  } else {
      gsap.to(companySubmenu.value, { height: 0, duration: 0.5 });
  }
});

watch(showUserSubmenu, (newVal) => {
  if (newVal) {
      gsap.to(userSubmenu.value, { height: 'auto', duration: 0.5 });
  } else {
      gsap.to(userSubmenu.value, { height: 0, duration: 0.5 });
  }
});

// ------------------ sidebar animation ------------------

const sidebar = ref(null);

watch(() => props.showSidebar, (newVal) => {
  if (newVal) {
      gsap.to(sidebar.value, { x: 0, duration: 0.5 });
  } else {
      gsap.to(sidebar.value, { x: -sidebar.value.offsetWidth, duration: 0.5 });
  }
});

onMounted(() => {
  if (props.showSidebar) {
      gsap.set(sidebar.value, { x: 0 });
  } else {
      gsap.set(sidebar.value, { x: -sidebar.value.offsetWidth });
  }
  gsap.set(companySubmenu.value, { height: 0 });
  gsap.set(userSubmenu.value, { height: 0 });
});
</script>


