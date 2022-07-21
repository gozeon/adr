<template>
    <div class="underline decoration-gray-400 hover:(bg-stone-100) px-4" v-for="project in projects" :key="project.id"
        @click.stop="handleGotoRecord(project.ID)" @dblclick.stop="handleLoadDetail(project.ID)">
        {{ project.title }}</div>

    <Modal v-if="hasRole('update')" type="update">
        <form @submit.prevent="handleSubmit" class="bg-white w-3/4 xl:w-1/3  m-auto px-8 py-5 mt-3 opacity-100">
            <div class="grid grid-cols-4 col-span-2 mb-5">
                <label class="col-span-1 self-center">ID</label>
                <input v-model="formData.ID" type="text" readonly
                    class="col-span-3 p-3 focus:(border-dark-500 ring-gray-500)" placeholder="Project Name">
            </div>
            <div class="grid grid-cols-4 col-span-2 mb-5">
                <label class="col-span-1 self-center">Name</label>
                <input v-model="formData.title" type="text" required minlength="3" maxlength="50"
                    class="col-span-3 p-3 focus:(border-dark-500 ring-gray-500) placeholder-gray-500"
                    placeholder="Project Name">
            </div>

            <div class="grid grid-cols-4 col-span-2 mb-5">
                <label class="col-span-1 self-center">Description</label>
                <textarea v-model="formData.description" maxlength="120"
                    class="col-span-3 resize-none p-3 focus:(border-dark-500 ring-gray-500) placeholder-gray-500"
                    placeholder="Project Description"></textarea>
            </div>

            <div class="grid grid-cols-4 col-span-2">
                <button type="submit"
                    class="col-start-2 py-2 bg-light-100 rounded-sm border-1 border-dark-300 hover:(bg-dark-400 text-white) focus:(ring-2 ring-gray-300)">Save</button>
            </div>

        </form>
    </Modal>
</template>

<script setup lang="ts">
import { reactive, computed } from 'vue'
import { useRouter } from 'vue-router'

import Modal from './Modal/index.vue'
import { useProjectStore } from "../stores/project"
import useToggleModal from './Modal/toggleModal';

const router = useRouter()
// declare store variable
const store = useProjectStore()


// declare store variable
const { openModal, hasRole, closeModal } = useToggleModal()

// https://github.com/vuejs/core/issues/4960#issuecomment-980443607
await store.loadProject()

const projects: any = computed(() => {
    return store.projects
})


const formData = reactive<any>({})

const handleSubmit = async () => {
    const res = await store.createOrUpdateProject({ ...formData })
    if (res.data?.errNo === 0) {
        store.loadProject()
        closeModal('update')
        Object.assign(formData, {})
    } else {
        alert(res.data?.errMsg)
    }

}

const handleLoadDetail = async (id: any) => {
    const { data } = await store.loadDetail(id)
    if (data?.errNo === 0) {
        Object.assign(formData, data.data)
        openModal('update')
    } else {
        alert(data?.errMsg)
    }
}

const handleGotoRecord = async (id: any) => {
    router.push({
        name: 'about',
        query: { id }
    })

}
</script>