import {createRouter, createWebHashHistory} from 'vue-router'
import Login from '../views/Login.vue'
import Stream from '../views/Stream.vue'
import Profile from '../views/Profile.vue'
import Photo from '../views/Photo.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/session', name: 'Login', component: Login},
		{path: '/', name: 'Stream', component: Stream},
		{path: '/users/:username/profile', name: 'Profile', component: Profile},
		{path: '/users/:username/photos/:photoId', name: 'Photo', component: Photo}
	],
	linkActiveClass: 'active-router'
})

export default router
