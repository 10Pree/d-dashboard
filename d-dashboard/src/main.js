import 'bootstrap/dist/css/bootstrap.css'
import { createApp } from 'vue'
import App from './App.vue'

// สำหรับ Font Awesome
import '@fortawesome/fontawesome-free/css/all.css'

// สำหรับ Material Icons
import '@mdi/font/css/materialdesignicons.css'

// สร้างแอปพลิเคชัน Vue
const app = createApp(App)

// นำเข้า CSS และใช้งาน
app.use('bootstrap')

// นำเข้า Font Awesome
app.use('@fortawesome/fontawesome-free/css/all.css')

// นำเข้า Material Icons
app.use('@mdi/font/css/materialdesignicons.css')

// เปิดการใช้งานแอปพลิเคชัน
app.mount('#app')
