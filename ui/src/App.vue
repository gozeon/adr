<script setup lang="ts">
import { computed } from '@vue/reactivity';
import { RouterView, RouterLink, useRouter } from 'vue-router'

import { useProjectStore } from "./stores/project";

const store = useProjectStore()
const activeProject = computed<any>(() => {
  if (useRouter().currentRoute.value.name == 'home') {
    return {}
  }
  return store.activeProject
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

    <div class="w-full md:w-3/4 xl:w-2/4 m-auto">
      <RouterView />
    </div>
  </div>
</template>
