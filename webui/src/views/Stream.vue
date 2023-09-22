<script>
import Navbar from '../components/Navbar.vue'
import Grid from '../components/Grid.vue';
export default {
	async mounted() {
		this.userId = localStorage.getItem('userId')
		if (!this.userId){
			this.$router.push({name: 'Login'})
		}
		await this.getPhotos();
	},
	components: {
		Navbar,
		Grid
	},
	name: 'Stream',
	data (){
		return{
			username: localStorage.getItem('username'),
			photos: null,
			userId: null
		}
	},
	methods: {
		async getPhotos() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get(`/users/${this.username}/stream`,{
                headers: {Authorization: `Bearer ${this.userId}`}
                });
				if (response.status !== 200) {
				throw response.status;
				}
				this.photos = JSON.parse(JSON.stringify(response.data));
				

			} catch (e) {
				this.errormsg = e.toString();	
			}
			this.loading = false;
   	 	}
	}
}
</script>

<template>
	<div id="profile" class="container">
		<p>Home</p>
		<Navbar :username="username"></Navbar>
		<Grid v-if="photos !== null" :username="username" :photos="photos" :userId="userId"></Grid>
	</div>
</template>
<style scoped>
p {
    font-size: 25px;
	position: absolute;
	margin-left: 200px;
	margin-top: 49px;

}
</style>