<template>
    <div>
        {{ companyStore.company }}
    </div>
</template>


<script setup>
import { onBeforeMount, onMounted, ref } from 'vue';
import { useCompanyStore } from '@/stores/Company';

const companyStore = useCompanyStore();

onMounted(() => {
    companyStore.getCompany();
}),
</script>