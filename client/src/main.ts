import { createApp } from "vue";
import "./assets/index.css";
import App from "./App.vue";

import { createPinia } from "pinia";
import initRouter from "@/router";
import { autoAnimatePlugin } from "@formkit/auto-animate/vue";
import gsap from "gsap";
import { ScrollTrigger } from "gsap/ScrollTrigger";
import { ScrollToPlugin } from "gsap/ScrollToPlugin";

// import posthogPlugin from "./plugins/posthog";

gsap.registerPlugin(ScrollTrigger, ScrollToPlugin);

(async () => {
  const app = createApp(App);

  const router = initRouter();
  app.use(router);

  const pinia = createPinia();
  app.use(pinia);

  app.use(autoAnimatePlugin);

  // app.use(posthogPlugin);

  app.mount("#app");
})();
