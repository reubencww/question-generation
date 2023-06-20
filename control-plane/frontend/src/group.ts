import { RouteRecordRaw } from "vue-router";

export function group(prefix: string, routes: RouteRecordRaw[]) {
    return routes.map((route) => {
        if (route.path == "/") {
            route.path = "";
        }

        route.path = `/${prefix}${route.path}`;
        route.name = `${prefix}.${String(route.name)}`;

        return route;
    });
}
