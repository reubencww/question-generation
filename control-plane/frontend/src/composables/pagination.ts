import { Ref, computed } from "vue";

import { paginate } from "@/paginate";

/**
 * Computes the required tabs that should be displayed for pagination.
 *
 * @param totalItems
 * @param currentPage
 * @param pageSize
 * @param maxPagesToDisplay
 */
export function usePaginateAvailablePages(
    totalItems: Ref<number>,
    currentPage: Ref<number>,
    pageSize: number = 20,
    maxPagesToDisplay: number = 5
) {
    const pagesAvailable = computed(() => {
        let pagination = paginate(
            totalItems.value,
            currentPage.value,
            pageSize,
            maxPagesToDisplay
        );
        let toDisplay: (number | string)[] = pagination.pages;

        if (!toDisplay.includes(1)) {
            toDisplay.unshift(1, "...");
        }

        if (!toDisplay.includes(pagination.totalPages)) {
            toDisplay.push("...", pagination.totalPages);
        }

        return toDisplay.map((x) => x.toString());
    });

    return { pagesAvailable };
}
