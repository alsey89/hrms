<template>
    <div v-auto-animate class="w-full">
        <Card v-if="showCard" class="w-[400px] md:w-[480px] flex flex-col gap-2 p-4">
            <div>
                <h1 class="text-2xl font-bold">Select Company</h1>
                <p class="text-sm">Select a company to sign in</p>
            </div>
            <div v-for="company in userStore.companies">
                <Button @click="onSelect(company.id)" class="w-full">
                    <div>{{ company.name }} - {{ company.id }}</div>
                </Button>
            </div>
        </Card>
    </div>
</template>

<script setup>
import { onMounted, reactive, ref } from 'vue';
import { useRouter } from 'vue-router';

import { Icon } from '@iconify/vue';
import { Card } from '@/components/ui/card';
import { Button } from '@/components/ui/button';

import { useUserStore } from '@/stores/User';

const router = useRouter();
const userStore = useUserStore();

const showCard = ref(false);

//get email from query params
const email = router.currentRoute.value.query.email;

const onSelect = async (companyId) => {
    if (!email) {
        console.error('Email not found');
        return;
    }
    const success = await userStore.getJwt(companyId, email, router);
};

onMounted(() => {
    userStore.getCsrfToken();
    showCard.value = true;
});
</script>