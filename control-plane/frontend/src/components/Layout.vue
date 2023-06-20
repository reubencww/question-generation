<script setup lang="ts">
import NavLink from "@/components/NavLink.vue";
</script>
<template>
    <div>
        <div class="font-sans antialiased" id="main-content">
            <nav
                class="border-t-4 border-b border-gray-100 border-t-red-500 bg-gray-800 text-white"
            >
                <!-- Primary Navigation Menu -->
                <div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
                    <div class="flex h-16 justify-between">
                        <div class="flex">
                            <!-- Navigation Links -->
                            <div class="hidden space-x-8 sm:flex">
                                <NavLink to="challenge.index">
                                    <span class="text-white">All</span>
                                </NavLink>
                                <NavLink to="challenge.create">
                                    <span class="text-white">Create</span>
                                </NavLink>
                            </div>
                        </div>
                    </div>
                </div>
            </nav>

            <div class="py-12">
                <div class="mx-auto max-w-7xl sm:px-6 lg:px-8">
                    <div
                        class="overflow-hidden bg-white shadow-sm sm:rounded-lg p-8"
                    >
                        <router-view />
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
