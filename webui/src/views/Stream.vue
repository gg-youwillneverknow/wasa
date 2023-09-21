<script>
import Navbar from '../components/Navbar.vue'
import Grid from '../components/Grid.vue';
export default {
	async mounted() {
		let userId = localStorage.getItem('userId')
		if (!userId){
			this.$router.push({name: 'Login'})
		}
	},
	components: {
		Navbar,
		Grid
	},
	name: 'Stream',
	data (){
		return{
			username: localStorage.getItem('username'),
			photos: null
		}
	},
	methods: {
		async getPhotos() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get(`/users/${this.username}/stream`,{
                headers: {Authorization: "Bearer ${token}",token: localStorage.getItem("userId")}
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
	},
	created() {
    	this.getPhotos();
  	}

}
</script>

<template>
	<div id="profile" class="container">
		<p>Home</p>
		<Navbar :username="username"></Navbar>
		<Grid v-if="photos !== null" :username="username" :photos="photos"></Grid>
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