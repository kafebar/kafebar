import { createApp } from 'vue'
import "./main.css"
import App from './App.vue'
import HomePage from "./pages/home.vue"
import {createRouter, createWebHistory} from 'vue-router'
import 'vue-toast-notification/dist/theme-bootstrap.css';
import ToastPlugin from 'vue-toast-notification'

const routes = [
  { path: '/', component: HomePage },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

const app = createApp(App)
app.use(router);
app.use(ToastPlugin);

app.mount('#app')
