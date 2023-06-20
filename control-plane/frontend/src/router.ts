import NProgress from "nprogress";
import { nextTick } from "vue";
import { RouteRecordRaw, createRouter, createWebHistory } from "vue-router";

import ChallengeCreate from "@/components/Challenge/Create.vue";
import ChallengeIndex from "@/components/Challenge/Index.vue";
import ChallengeShow from "@/components/Challenge/Show.vue";
import HelloWorld from "@/components/HelloWorld.vue";
import Dashboard from "@/components/Layout.vue";
import PageNotFound from "@/components/PageNotFound.vue";
import { group } from "@/group";

const routes: RouteRecordRaw[] = [
    {
        path: "/",
        component: Dashboard,
        children: [
            { path: "/", name: "index", redirect: { name: "challenge.index" } },
            ...group("challenge", [
                {
                    path: "/",
                    name: "index",
                    component: ChallengeIndex,
                },
                {
                    path: "/:id(\\d+)",
                    name: "show",
                    component: ChallengeShow,
                },
                {
                    path: "/create",
                    name: "create",
                    component: ChallengeCreate,
                },
            ]),
        ],
    },
    {
        path: "/:pathMatch(.*)*",
        component: PageNotFound,
        meta: { title: "404" },
    },
];

let router = createRouter({
    history: createWebHistory(),
    routes,
});

router.beforeEach((to, from) => {
    // if this is a fresh page load
    if (from.name) {
        NProgress.start();
    }
});

router.afterEach(() => {
    NProgress.done();
});

export { router };

export const registerDefaultTitle = (defaultTitle: string) => {
    router.afterEach((to, from) => {
        nextTick(() => {
            if (typeof to.meta?.title === "undefined") {
                document.title = defaultTitle;
                return;
            }

            let title =
                typeof to.meta.title === "function"
                    ? to.meta.title(to)
                    : to.meta.title;

            document.title = `${defaultTitle} :: ${title}`;
        }).then();
    });
};
