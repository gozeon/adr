<template>
    <teleport to="body">
        <transition enter-from-class="translate-y-0">
            <div class="bg-dark-500 bg-opacity-60 h-[100vh] w-[100vw] z-50 fixed inset-0" @click.self="closeModal(type)">
                <slot></slot>
            </div>
        </transition>
    </teleport>

</template>

<script setup lang="ts">
import { defineProps } from 'vue'
import useToggleModal from './toggleModal'
const { closeModal } = useToggleModal()

defineProps(['type'])
</script>
