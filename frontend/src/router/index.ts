import { createRouter, createWebHistory } from 'vue-router';
import Layout from "../views/Layout.vue"
import Home from "../views/Home.vue"
import About from "../views/About.vue"
import Profile from "../views/Profile.vue"
import Items from "../views/Items.vue"

const routes = [
  {
    path: '/',
    component: Layout, // Usa el layout como componente base
    children: [
      {
        path: 'home',
        name: 'Home',
        component: Home, // Vista para la página de inicio
      },
      {
        path: 'about',
        name: 'About',
        component: About, // Vista para la página Acerca
      },
      {
        path: 'profile',
        name: 'Profile',
        component: Profile, // Vista para la página de Contacto
      },
      {
        path: 'items',
        name: 'Items',
        component: Items, // Vista para la página de Items
      },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;