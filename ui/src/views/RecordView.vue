<template>
  <h3 class="mt-10 text-3xl cursor-pointer divide-light-600 px-3 md:px-4" @click="openModal('create')">
    <span class="px-2 rounded-md transform scale-" :class="{
      'bg-notion-yellow_background': record.status == 1,
      'bg-notion-teal_background': record.status == 2,
      'bg-notion-brown_background': record.status == 3,
      'bg-notion-red_background': record.status == 4,
      'bg-notion-gray_background': record.status == 5,
    }">{{ recordStore.getStatusLabelById(record.status) }}</span>
    {{ record.title }}
  </h3>

  <small class="px-3 md:px-4">修改时间: {{ formatTime(record.UpdatedAt) }}</small>
  <p class="mt-5 px-3 md:px-4 cursor-pointer" v-html="record.description">
  </p>

  <div class="px-3 md:px-5 mt-6 m-3 border-1 " v-for="comment in comments">
    <div class="py-1 border-b-1"><small>{{ formatTime(comment.CreatedAt) }}</small></div>
    <div v-html="comment.description"></div>
  </div>

  <div class="m-3 mt-12">
    <div class="h-[150px] box-border">
      <QuillEditor placeholder="Record Comment" content-type="html" v-model:content="description">
      </QuillEditor>
    </div>
    <button
      class="mt-[50px] p-2 bg-light-100 rounded-sm border-1 border-dark-300 hover:(bg-dark-400 text-white) focus:(ring-2 ring-gray-300)"
      @click="handleAddComment">
      Add Comment</button>
  </div>

  <Modal v-if="hasRole('create')" type="create">
    <form @submit.prevent="handleCreate" class="bg-white w-3/4 xl:w-2/3  m-auto px-8 py-5 mt-3 opacity-100">
      <div class="grid grid-cols-4 <sm:grid-cols-1 col-span-2 mb-5">
        <label class="col-span-1 self-center md:justify-self-center">Title</label>
        <input type="text" v-model="record.title" required minlength="3" maxlength="50"
          class="col-span-3 p-3 focus:(border-dark-500 ring-gray-500) placeholder-gray-500" placeholder="Record Title">
      </div>

      <div class="grid grid-cols-4 <sm:grid-cols-1 col-span-2 mb-5">
        <label class="col-span-1 self-start md:justify-self-center">Description</label>
        <div class="col-span-3 h-[50vh]">
          <QuillEditor placeholder="Record Description" content-type="html" v-model:content="record.description">
          </QuillEditor>
        </div>
      </div>

      <div class="grid grid-cols-4 col-span-2 mt-20">
        <button type="submit"
          class="col-start-2 <sm:col-start-1 py-2 bg-light-100 rounded-sm border-1 border-dark-300 hover:(bg-dark-400 text-white) focus:(ring-2 ring-gray-300)">Save</button>
      </div>
    </form>
  </Modal>

</template>

<script setup lang="ts">
import mediumZoom from 'medium-zoom'
import { computed, ref, watchEffect } from "vue";
import { useRoute } from "vue-router";
import { QuillEditor } from '@vueup/vue-quill'

import '@vueup/vue-quill/dist/vue-quill.snow.css';
import { useProjectStore } from "../stores/project";
import { useRecordStore } from "../stores/record";
import { useCommentStore } from "../stores/comment";
import useToggleModal from '../components/Modal/toggleModal';
import Modal from '../components/Modal/index.vue'

import { DateTime } from 'luxon'


const route = useRoute();
const projectStore = useProjectStore();
const recordStore = useRecordStore();
const commentStore = useCommentStore()
const record = computed<any>(() => recordStore.activeRecord);
const comments = computed<any>(() => commentStore.activeComments);
const pid = route.params.pid
const recordId = route.params.id
const { openModal, hasRole, closeModal } = useToggleModal()
const description = ref('')

// load data
projectStore.loadDetail(pid)
recordStore.loadDetail(recordId)
commentStore.loadCommentById(recordId)

const handleCreate = async () => {
  const { data } = await recordStore.createOrUpdateRecord({ ...record.value })
  if (data?.errNo == 0) {
    recordStore.loadDetail(recordId)
    closeModal('create')
  } else {
    alert(data.errMsg)
  }
}

const handleAddComment = async () => {
  if (description.value) {
    const { data } = await commentStore.createOrUpdateComment({
      description: description.value,
      record_id: +recordId,
    })
    if (data?.errNo == 0) {
      // TODO: not work
      description.value = ""
      commentStore.loadCommentById(recordId)
    } else {
      alert(data.errMsg)
    }
  }

}

const formatTime = (time: any) => DateTime.fromISO(time).setLocale('zh-cn').toLocaleString(DateTime.DATETIME_MED_WITH_SECONDS)

watchEffect(() => {
  // TODO not woring
  mediumZoom('img')
})
</script>
