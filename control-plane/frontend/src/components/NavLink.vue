<script setup lang="ts">
import { useRoute } from "vue-router";

const { to } = defineProps<{
    to: string;
}>();
</script>

<template>
    <router-link
        :to="{ name: to }"
        class="inline-flex items-center px-1 pt-1 border-b-2 border-transparent text-sm font-medium leading-5 text-gray-500 hover:text-gray-300 hover:border-red-400 focus:outline-none focus:text-gray-700 focus:border-red-400 transition duration-150 ease-in-out"
    >
        <slot />
    </router-link>
</template>

<style scoped>
.router-link-active {
    @apply border-red-400;
}
</style>
