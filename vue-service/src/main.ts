import { createApp } from 'vue';
import { createPinia } from 'pinia';
import { createVuetify } from 'vuetify';
import * as components from 'vuetify/components';
import * as directives from 'vuetify/directives';
import { aliases, mdi } from 'vuetify/iconsets/mdi';
import '@mdi/font/css/materialdesignicons.css';
import 'vuetify/styles';
import './styles/main.css';
import App from './App.vue';
import { router } from './router';

const vuetify = createVuetify({
  components,
  directives,
  icons: { defaultSet: 'mdi', aliases, sets: { mdi } },
  theme: {
    defaultTheme: 'matrixTheme',
    themes: {
      matrixTheme: {
        dark: false,
        colors: {
          background: '#f4f7f5',
          surface: '#ffffff',
          primary: '#126782',
          secondary: '#df7f45',
          accent: '#d7ed69',
          'on-background': '#12232d',
          'on-surface': '#12232d'
        }
      }
    }
  }
});

createApp(App).use(createPinia()).use(router).use(vuetify).mount('#app');
