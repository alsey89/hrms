<template>
    <div class="w-full h-full flex flex-col gap-2">
        <div class="h-8 flex gap-2 text-lg font-semibold px-4 py-2">
            <div> Departments </div>
        </div>
        <hr class="border-t mx-2">
        <div v-for="department in companyStore.departments" :key="department.id">
            <div class="flex justify-between items-center px-4">
                <div class="h-8 flex items-center gap-2 text-lg font-semibold py-2">
                    {{ department.name }}
                </div>
                <button @click="onClickEditDepartmentButton(department)"
                    class="border rounded-md hover:cursor-pointer hover:bg-primary hover:text-primary-foreground">
                    <Icon v-if="!showEditDepartmentCard" icon="mdi:pencil" class="text-xl" />
                    <Icon v-else icon="mdi:close" class="text-xl" />
                </button>
            </div>
            <div v-auto-animate class="flex flex-col gap-2 px-4">
                <Card v-if="showEditDepartmentCard" ref="editDepartmentCard" class="w-full p-4">
                    <form @submit.prevent="onSubmitForm" class="flex flex-col gap-2">
                        <div class="flex flex-col">
                            <label for="departmentName"> Name </label>
                            <input type="text" id="departmentName" v-model.trim="editDepartmentForm.name"
                                class="bg-background border px-2" />
                        </div>
                        <div class="flex flex-col">
                            <label for="departmentDescription"> Description </label>
                            <textarea id="departmentDescription" v-model.trim="editDepartmentForm.description"
                                class="bg-background border px-2"></textarea>
                        </div>
                        <button type="submit"
                            class="border rounded-md hover:cursor-pointer hover:bg-primary hover:text-primary-foreground mt-2">
                            Submit
                        </button>
                    </form>
                </Card>
                <Card class="w-full flex flex-col gap-4 p-4">
                    <div>
                        <p class="font-bold">Description</p>
                        <p>{{ department.description }}</p>
                    </div>
                </Card>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue';
import { useCompanyStore } from '@/stores/Company';
import { Icon } from '@iconify/vue';
import { Card } from '@/components/ui/card';

const companyStore = useCompanyStore();

const showEditDepartmentCard = ref(false);
const editDepartmentForm = reactive({
    id: '',
    name: '',
    description: ''
});
const onClickEditDepartmentButton = (department) => {
    if (showEditDepartmentCard.value) {
        showEditDepartmentCard.value = false;
        return;
    }
    showEditDepartmentCard.value = true;
    editDepartmentForm.id = department.id;
    editDepartmentForm.name = department.name;
    editDepartmentForm.description = department.description;
};
const onSubmitForm = async () => {
    await companyStore.updateDepartment(editDepartmentForm);
    showEditDepartmentCard.value = false;
}

onMounted(() => {
    companyStore.getCompany();
})
</script>
