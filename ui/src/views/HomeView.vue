<template>
  <h3 class="mt-10 text-3xl px-3 md:px-4">Projects</h3>

  <div class="mt-10 grid md:grid-cols-2 gap-2 cursor-pointer px-3 md:px-5 grid-flow-row">
    <Suspense>
      <template #default>
        <ProjectList></ProjectList>
      </template>

      <template #fallback>
        <p>Loading...</p>
      </template>
    </Suspense>
  </div>

  <div class="mt-6 grid md:grid-cols-2 gap-2 cursor-pointer px-3 md:px-5 grid-flow-row">
    <div class="decoration-gray-400 hover:(bg-stone-100) px-4 text-center border-1 border-dark-500 border-dashed"
      @click="openModal('create')">+
    </div>
  </div>

  <Modal v-if="hasRole('create')" type="create">
    <form @submit.prevent="handleSubmit" class="bg-white w-3/4 xl:w-1/3  m-auto px-8 py-5 mt-3 opacity-100">
      <div class="grid grid-cols-4 <sm:grid-cols-1 col-span-2 mb-5">
        <label class="col-span-1 self-center">Name</label>
        <input v-model="formData.title" type="text" required minlength="3" maxlength="50"
          class="col-span-3 p-3 focus:(border-dark-500 ring-gray-500) placeholder-gray-500" placeholder="Project Name">
      </div>

      <div class="grid grid-cols-4 <sm:grid-cols-1 col-span-2 mb-5">
        <label class="col-span-1 self-center">Description</label>
        <textarea v-model="formData.description" maxlength="120"
          class="col-span-3 resize-none p-3 focus:(border-dark-500 ring-gray-500) placeholder-gray-500"
          placeholder="Project Description"></textarea>
      </div>

      <div class="grid grid-cols-4 col-span-2">
        <button type="submit"
          class="col-start-2 <sm:col-start-1 py-2 bg-light-100 rounded-sm border-1 border-dark-300 hover:(bg-dark-400 text-white) focus:(ring-2 ring-gray-300)">Save</button>
      </div>

    </form>
  </Modal>
</template>

<script setup lang="ts">
import { defineAsyncComponent, reactive } from 'vue'

import useToggleModal from '../components/Modal/toggleModal';
import Modal from '../components/Modal/index.vue'
import { useProjectStore } from "../stores/project"

// declare store variable
const store = useProjectStore()
const ProjectList = defineAsyncComponent(() => import('../components/ProjectList.vue'))
const { openModal, hasRole, closeModal } = useToggleModal()

// https://github.com/vuejs/core/issues/1081
const defaultFormData = {
  title: '',
  description: ''
}
const formData = reactive({ ...defaultFormData })

const handleSubmit = async () => {
  const res = await store.createOrUpdateProject({ ...formData })
  if (res.data?.errNo === 0) {
    store.loadProject()
    closeModal('create')
    Object.assign(formData, { ...defaultFormData })
  } else {
    alert(res.data?.errMsg)
  }

}
</script>
