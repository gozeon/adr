<script setup lang="ts">
import { computed } from '@vue/reactivity';
import { watch } from 'vue';
import { RouterView, RouterLink, useRouter, useRoute } from 'vue-router'

import useToggleModal from './components/Modal/toggleModal';
import { useProjectStore } from "./stores/project";

const route = useRoute()
const store = useProjectStore()
const { closeAll } = useToggleModal()
const activeProject = computed<any>(() => {
  if (useRouter().currentRoute.value.name == 'home') {
    return {}
  }
  return store.activeProject
})

watch(() => route.name, () => {
  // TODO: 应该有更好的方法
  closeAll()
})
</script>

<template>
  <div class="font-mono tracking-wide">
    <header class="sticky top-0 bg-white py-3 px-8">
      <RouterLink to="/">Adr</RouterLink>
      <RouterLink v-if="activeProject.ID" :to="{ name: 'about', params: { id: activeProject.ID } }"> / {{
          activeProject.title
      }}</RouterLink>
    </header>

    <div class="w-full md:w-3/4 xl:w-2/4 m-auto break-all">
      <RouterView />
    </div>
  </div>
</template>
