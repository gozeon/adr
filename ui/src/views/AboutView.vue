<template>
  <h3 class="mt-10 text-3xl cursor-pointer divide-light-600 px-3 md:px-4" @click="openModal('update')">{{ project.title
  }}</h3>

  <p class="mt-4 px-3 md:px-4 cursor-pointer" @click="openModal('update')">
    {{ project.description }}
  </p>

  <ul class="list-decimal mt-12 px-20">
    <li class="mb-3 hover:(bg-stone-100) pl-4 cursor-pointer" v-for="record in records">

      <div class="flex justify-between <sm:(flex-col gap-2)" @click.self="handleDetail(record.ID)">
        <div class="underline" :class="{
          'text-notion-yellow': record.status == 1,
          'text-notion-teal': record.status == 2,
          '': record.status == 3,
          'text-notion-red line-through': record.status == 4,
          'text-notion-gray line-through': record.status == 5,
        }" @click.self="handleDetail(record.ID)">{{ record.title }}</div>
        <select class="px-10" :class="{
          'bg-notion-yellow_background': record.status == 1,
          'bg-notion-teal_background': record.status == 2,
          '': record.status == 3,
          'bg-notion-red_background': record.status == 4,
          'bg-notion-gray_background': record.status == 5,
        }" :value="record.status" @change="handleUpdateStatus($event, record.ID)">
          <option v-for="item in statusMap" :value="item.val"> {{ item.label }}</option>
        </select>
      </div>
    </li>
  </ul>

  <div
    class="decoration-gray-400 px-4 text-center border-1 border-dark-500 border-dashed cursor-pointer hover:(bg-stone-100) mx-20 mt-5 text-center w-2/4 xl:w-1/3 <sm:(w-9/12 mt-10 mx-auto)"
    @click="openModal('create')">
    +
  </div>


  <div class="py-20"></div>

  <Modal v-if="hasRole('create')" type="create">
    <form @submit.prevent="handleCreate"
      class="bg-white w-3/4 xl:w-2/3 <sm:w-11/12 m-auto px-8 py-5 mt-3 opacity-100 break-all">
      <div class="grid grid-cols-4 <sm:grid-cols-1 col-span-2 mb-5">
        <label class="col-span-1 self-center md:justify-self-center">Title</label>
        <input type="text" v-model="formData.title" required minlength="3" maxlength="50"
          class="col-span-3 p-3 focus:(border-dark-500 ring-gray-500) placeholder-gray-500" placeholder="Record Title">
      </div>

      <div class="grid grid-cols-4 <sm:grid-cols-1 col-span-2 mb-5">
        <label class="col-span-1 self-start md:justify-self-center">Description</label>
        <div class="col-span-3 h-[50vh]">
          <QuillEditor placeholder="Record Description" content-type="html" v-model:content="formData.description">
          </QuillEditor>
        </div>
      </div>

      <div class="grid grid-cols-4 col-span-2 mt-20 <sm:mt-30">
        <button type="submit"
          class="col-start-2 <sm:col-start-1 py-2 bg-light-100 rounded-sm border-1 border-dark-300 hover:(bg-dark-400 text-white) focus:(ring-2 ring-gray-300)">Save</button>
      </div>
    </form>
  </Modal>

  <Modal v-if="hasRole('update')" type="update">
    <form @submit.prevent="handleUpdate" class="bg-white w-3/4 xl:w-1/3 <sm:w-11/12 m-auto px-8 py-5 mt-3 opacity-100">
      <div class="grid grid-cols-4 <sm:grid-cols-1 col-span-2 mb-5">
        <label class="col-span-1 self-center">ID</label>
        <input v-model="project.ID" type="text" readonly class="col-span-3 p-3 focus:(border-dark-500 ring-gray-500)"
          placeholder="Project Name">
      </div>
      <div class="grid grid-cols-4 <sm:grid-cols-1 col-span-2 mb-5">
        <label class="col-span-1 self-center">Name</label>
        <input v-model="project.title" type="text" required minlength="3" maxlength="50"
          class="col-span-3 p-3 focus:(border-dark-500 ring-gray-500) placeholder-gray-500" placeholder="Project Name">
      </div>

      <div class="grid grid-cols-4 <sm:grid-cols-1 col-span-2 mb-5">
        <label class="col-span-1 self-center">Description</label>
        <textarea v-model="project.description" maxlength="120"
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
import { reactive, computed, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import { QuillEditor } from '@vueup/vue-quill'

import '@vueup/vue-quill/dist/vue-quill.snow.css';
import { useProjectStore } from "../stores/project";
import { useRecordStore } from "../stores/record";
import useToggleModal from '../components/Modal/toggleModal';
import Modal from '../components/Modal/index.vue'


const route = useRoute();
const router = useRouter();
const projectStore = useProjectStore();
const recordStore = useRecordStore();
const project = computed<any>(() => projectStore.activeProject);
const records = computed<any>(() => recordStore.activeRecords);
const statusMap = computed(() => recordStore.statusMap)
const pid = route.params.id
const { openModal, hasRole, closeModal } = useToggleModal()


// load data
projectStore.loadDetail(pid)
recordStore.loadRecordByID(pid)

const handleUpdateStatus = async (event: any, id: any) => {
  await recordStore.updateStatus(id, +event.target.value)
  recordStore.loadRecordByID(pid)
}

const defaultFormData = {
  title: '',
  description: '',
  status: 1
}

const formData = reactive<any>({ ...defaultFormData })

const handleCreate = async () => {
  console.log(formData)
  const { data } = await recordStore.createOrUpdateRecord({ ...formData, project_id: +pid })
  if (data?.errNo == 0) {
    Object.assign(formData, { ...defaultFormData })
    recordStore.loadRecordByID(pid)
    closeModal('create')
  } else {
    alert(data.errMsg)
  }
}

const handleDetail = async (id: any) => {
  router.push({
    name: 'record',
    params: { id, pid }
  })
  return
  const { data } = await recordStore.loadDetail(id)
  if (data?.errNo == 0) {
    Object.assign(formData, { ...data.data })

    openModal('create')
  } else {
    alert(data.errMsg)
  }
}

const handleUpdate = async () => {
  const res = await projectStore.createOrUpdateProject({ ...project.value })
  if (res.data?.errNo === 0) {
    projectStore.loadDetail(pid)
    closeModal('update')
    Object.assign(formData, {})
  } else {
    alert(res.data?.errMsg)
  }

}

// const loadData = () => Promise.all([
//   projectStore.loadDetail(pid),
//   recordStore.loadRecordByID(pid),
// ]).then(res => {
//   const [projectRes, recordRes] = res
//   if (projectRes.data?.errNo == 0) {
//     Object.assign(project, projectRes.data.data)
//   } else {
//     alert(projectRes.data?.errMsg)
//   }

//   if (recordRes.data?.errNo == 0) {
//     Object.assign(records, recordRes.data.data)
//   } else {
//     alert(recordRes.data?.errMsg)
//   }
// })

// watch(
//   () => route.query.id,
//   newId => {
//     // https://stackblitz.com/edit/vitejs-vite-tksrjl?file=src%2Fabout.vue
//     // TOOD: go back home with undefined
//     if (newId === undefined) {
//       return;
//     }
//     loadData(newId)
//   },
//   { immediate: true }
// );

</script>
