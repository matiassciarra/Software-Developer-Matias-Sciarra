import { createApp } from 'vue'
import './index.css'
import App from './App.vue'
import router from './router/index.ts';

const app = createApp(App);

app.use(router); // Agrega el router a la aplicación
app.mount('#app');

