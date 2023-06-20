<template>
    <div>
        <div class="flex items-center mb-2 w-full">
            <div class="flex items-end">
                <h2 class="mr-2 text-3xl">Challenges</h2>
            </div>

            <div class="ml-auto"></div>

            <div class="flex justify-between items-center text-sm">
                <nav aria-label="Page Navigation">
                    <ul class="inline-flex -space-x-px">
                        <li>
                            <button
                                @click="prev"
                                :disabled="isFirstPage"
                                class="h-full py-2 px-3 ml-0 leading-tight text-gray-500 bg-white rounded-l-lg border border-gray-300 hover:bg-gray-100 hover:text-gray-700"
                            >
                                Previous
                            </button>
                        </li>
                        <li v-for="item in pagesAvailable" :key="item">
                            <!-- If it's the current page -->
                            <button
                                v-if="currentPage === parseInt(item)"
                                disabled
                                class="h-full py-2 px-3 text-gray-600 bg-gray-300 border border-gray-300 hover:bg-gray-400 hover:text-gray-900"
                            >
                                {{ item }}
                            </button>
                            <!-- If it's a ... -->
                            <button
                                v-else-if="item === '...'"
                                disabled
                                class="h-full py-2 px-3 text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700"
                            >
                                {{ item }}
                            </button>
                            <!-- If it's a navigable page -->
                            <button
                                v-else
                                @click="currentPage = parseInt(item)"
                                class="h-full py-2 px-3 text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700"
                            >
                                {{ item }}
                            </button>
                        </li>
                        <li>
                            <button
                                @click="next"
                                :disabled="isLastPage"
                                class="h-full py-2 px-3 leading-tight text-gray-500 bg-white rounded-r-lg border border-gray-300 hover:bg-gray-100 hover:text-gray-700"
                            >
                                Next
                            </button>
                        </li>
                    </ul>
                </nav>
            </div>
        </div>

        <div
            class="overflow-hidden w-full rounded-lg shadow-xs"
            v-if="trades.length > 0"
        >
            <div class="overflow-x-auto w-full">
                <table class="w-full whitespace-no-wrap">
                    <thead>
                        <tr
                            class="text-xs font-semibold tracking-wide text-left text-gray-500 uppercase bg-gray-50 border-b"
                        >
                            <th class="px-4 py-3">ID</th>
                            <th class="px-4 py-3">Name</th>
                            <th class="px-4 py-3">Questions</th>
                            <th class="px-4 py-3">Created</th>
                            <th class="px-4 py-3">Processed</th>
                            <th class="px-4 py-3 text-center">Actions</th>
                        </tr>
                    </thead>
                    <tbody class="bg-white divide-y">
                        <tr
                            class="text-gray-700 hover:bg-gray-200"
                            v-for="trade in trades"
                        >
                            <td class="px-4 py-3 text-sm">
                                {{ trade.id }}
                            </td>
                            <td class="px-4 py-3">
                                <div class="flex items-center text-sm">
                                    <div>
                                        <p class="font-semibold">
                                            {{ trade.name }}
                                        </p>
                                        <p>{{ trade.caption }}</p>
                                    </div>
                                </div>
                            </td>
                            <td class="px-4 py-3 text-sm">
                                {{ trade.questions.length }}
                            </td>
                            <td class="px-4 py-3 text-sm">
                                <Tippy>
                                    {{ fuzzyDate(trade.created_at) }}
                                    <template #content>
                                        {{ trade.created_at }}
                                    </template>
                                </Tippy>
                            </td>
                            <td class="px-4 py-3 text-sm">
                                <template v-if="trade.completed_at">
                                    <Tippy>
                                        {{ fuzzyDate(trade.completed_at) }}
                                        <template #content>
                                            {{ trade.completed_at }}
                                        </template>
                                    </Tippy>
                                </template>
                                <template v-else> Processing </template>
                            </td>
                            <td class="px-4 py-3 text-sm">
                                <div class="flex justify-center items-center">
                                    <router-link
                                        title="View"
                                        class="text-purple-500 bg-transparent border border-purple-500 hover:bg-purple-500 hover:text-white active:bg-purple-600 font-bold uppercase text-sm px-2 py-1.5 rounded outline-none focus:outline-none mb-1 ease-linear transition-all duration-150"
                                        :to="{
                                            name: 'challenge.show',
                                            params: { id: trade.id },
                                        }"
                                    >
                                        <EyeIcon class="h-5" />
                                    </router-link>
                                </div>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
            <div
                class="grid px-4 py-3 text-xs font-semibold tracking-wide text-gray-500 uppercase bg-gray-50 border-t sm:grid-cols-9"
            >
                <span class="flex col-span-3 items-center">
                    {{ meta.total }} trade(s) available
                </span>
            </div>
        </div>
        <div v-else>
            <p class="text-2xl text-gray-800">No records found</p>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { EyeIcon } from "@heroicons/vue/24/outline";
import { useOffsetPagination } from "@vueuse/core";
import { Ref, onBeforeMount, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import { Tippy } from "vue-tippy";

import { Challenge } from "@/challenges";
import { usePaginateAvailablePages } from "@/composables/pagination";
import { fuzzyDate } from "@/dates";
import http from "@/http";
import { resolvePage } from "@/paginate";

interface KycMeta {
    max: number;
    total: number;
}
interface KycResponse {
    data: Challenge[];
    meta: KycMeta;
}

const route = useRoute();
const router = useRouter();
const initialPage = resolvePage(route.query.page);
const trades: Ref<Challenge[]> = ref([]);
const meta: Ref<KycMeta> = ref({} as KycMeta);
const total: Ref<number> = ref(1);

onBeforeMount(async () => {
    await fetchKyc(initialPage, false);
});

async function fetchKyc(
    currentPage: number,
    shouldUpdateQueryParams: boolean = true
) {
    let { data } = await http.get<KycResponse>(
        `/challenge?page=${currentPage}`
    );
    trades.value = data.data;
    meta.value = data.meta;
    total.value = meta.value.total;
    // a really simple but effective way to ensure the internal page tracking play nicely with
    // fresh page loads. This means on hard page loads:
    //     - visiting /kyc stays on the 1st page without altering the existing query params
    //     - visiting /kyc?page=3 works
    if (shouldUpdateQueryParams) {
        await router.push({ query: { page: currentPage } });
    }
}

async function fetchData({ currentPage }: { currentPage: number }) {
    await fetchKyc(currentPage);
}

const { currentPage, pageCount, isFirstPage, isLastPage, prev, next } =
    useOffsetPagination({
        // TODO: explain the difference between meta.value.total/total ref
        total,
        page: initialPage,
        pageSize: 15,
        onPageChange: fetchData,
        onPageSizeChange: fetchData,
    });
const { pagesAvailable } = usePaginateAvailablePages(total, currentPage);
</script>
