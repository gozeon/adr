<template>
    <RouterLink :to="{ name: 'about', params: { id: project.ID } }"
        class="block px-4 underline decoration-gray-400 hover:(bg-stone-100)" v-for="project in projects"
        :key="project.id">{{ project.title }}
    </RouterLink>
</template>

<script setup lang="ts">
import { reactive, computed } from 'vue'
import { RouterLink } from 'vue-router'

import { useProjectStore } from "../stores/project"

// declare store variable
const store = useProjectStore()

// https://github.com/vuejs/core/issues/4960#issuecomment-980443607
await store.loadProject()

const projects: any = computed(() => {
    return store.projects
})

</script>
