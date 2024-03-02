import { createApp } from 'vue'
import "./main.css"
import App from './App.vue'
// import HomePage from "./pages/home.vue"
import CounterPage from "./pages/counter.vue"
import OrdersPage from "./pages/orders.vue"
import ProductsPage from "./pages/products.vue"
import {createRouter, createWebHistory} from 'vue-router'
import 'vue-toast-notification/dist/theme-bootstrap.css';
import ToastPlugin from 'vue-toast-notification'

const routes = [
  { path: '/', component: CounterPage },
  { path: '/orders', component: OrdersPage },
  { path: '/products', component: ProductsPage },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

const app = createApp(App)
app.use(router);
app.use(ToastPlugin);

app.mount('#app')
