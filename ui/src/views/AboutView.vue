<template>
  <div class="text-red-400 dark:text-green-400">

    <pre>{{ project }}</pre>
  </div>
</template>

<script setup lang="ts">
import { watch, reactive, } from 'vue';
import { useRoute } from 'vue-router';
import { useProjectStore } from "../stores/project"

const route = useRoute()
const store = useProjectStore()
const project = reactive<any>({})

watch(
  () => route.query.id,
  async newId => {
    // TOOD: go back home with undefined
    if (newId === undefined) {
      return
    }
    const { data } = await store.loadDetail(newId)
    if (data?.errNo === 0) {
      Object.assign(project, { ...data.data })
    } else {
      alert('detail ' + data?.errMsg)
    }
  },
  { immediate: true }
)

</script>
