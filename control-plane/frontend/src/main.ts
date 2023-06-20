import "@fontsource/nunito/300.css";
import "@fontsource/nunito/400-italic.css";
import "@fontsource/nunito/400.css";
import "@fontsource/nunito/500.css";
import "@fontsource/nunito/600.css";
import "@fontsource/nunito/700.css";
import Sockette from "sockette";
import "tippy.js/dist/tippy.css";
import { createApp } from "vue";

import App from "@/App.vue";
import "@/main.css";
import { registerDefaultTitle, router } from "@/router";

registerDefaultTitle(document.title);
const app = createApp(App);
app.use(router);
app.mount("#app");

new Sockette("ws://" + window.location.host + "/ws", {
    onmessage: (event) => {
        let data = JSON.parse(event.data);
        console.debug(data);
    },
    onopen: (event) => {
        console.debug("WebSocket connection established");
    },
});
