<script>
import Navbar from '../components/Navbar.vue'
import Modal from '../components/Modal.vue'
import Dropbox from '../components/Dropbox.vue';
import Mystats from '../components/Mystats.vue';
import Grid from '../components/Grid.vue';
import ErrorMsg from '../components/ErrorMsg.vue';
export default {
    name: 'Profile',
	
	components: {
		Navbar,
		Modal,
		Dropbox,
		Mystats,
		Grid,
		ErrorMsg
	},
	data (){
		return{
			errormsg: null,
			username: localStorage.getItem('username'),
			searchUsername: null,
			isbanned: false,
			searchbanned: false,
			isfollowing: false,
			photos: null,
			userId: localStorage.getItem('userId'),
			numfollowers: null, 
			numfollowings: null, 
			numposts: null,
			createPost: null,
			loading: null
		}
	},
	watch: {
		async '$route.params.username'(newUsername) {
			this.username=localStorage.getItem('username')
			this.isbanned=false
			if(newUsername){
			this.searchUsername=newUsername		
			
			await this.getProfile()
			if (this.errormsg){
				return
			}
			
			await this.checkBanned()
			if (this.isbanned===true){
				return
			}
			
			if (this.searchUsername!==this.username){
			await this.getBans()
			}		
		
			await this.getPhotos();
			}
		},
  	},
	methods: {
		async handleMessageFromDropbox () {
			await this.getProfile();
			await this.getPhotos();
		},
		handleMessageFromMystats(followers) {
      		for (const item of followers){
				if(this.username===item.Username){

					this.isfollowing=true
				}
			}	
    	},
		async getProfile() {
			this.loading = true;
			this.errormsg = null;
			try {
				
				let response = await this.$axios.get(this.$route.path,{
        			headers: {Authorization: `Bearer ${this.userId}`}
      			});
				let profile  = response.data;

				if (response.status!=200){
					throw(response.status)
				}
				this.numfollowers=profile["followers"]
				this.numfollowings=profile["followings"]
				this.numposts=profile["posts"]
				
			} catch (e) {
				this.errormsg = "User does not exist";
				
			}
			this.loading = false;
		},
		async getPhotos() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get(`/users/${this.searchUsername}/photos/`,{
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
    	},
		async checkBanned(){
			this.loading = true;
			this.errormsg = null;
			try {
				
				let response = await this.$axios.get(`/users/${this.searchUsername}/bans/`,{
        			headers: {Authorization: `Bearer ${this.userId}`}
      			});
				
				if (response.status!=200){
					throw(response.status)
				}

				let searchbans = JSON.parse(JSON.stringify(response.data));
				
				if (searchbans){				
					for(const item of searchbans){
						if(this.username===item.Username){
							this.isbanned=true
						}
					}
				}
				
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		async getBans(){
			this.loading = true;
			this.errormsg = null;
			try {
				
				let response = await this.$axios.get(`/users/${this.username}/bans/`,{
        			headers: {Authorization: `Bearer ${this.userId}`}
      			});
				
				if (response.status!=200){
					throw(response.status)
				}

				let bans = JSON.parse(JSON.stringify(response.data));
				if (bans){
					for(const item of bans){
						if(this.searchUsername===item.Username){
							this.searchbanned = true
						}
					}
				}		

			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		async handlebtn(){

			if(this.isfollowing===true){
	
				await this.unfollow()
				return
			}
			if(this.isfollowing===false){
			
				await this.follow()
				return 
			}
		},
		async follow() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.put(`/users/${this.username}/followings/${this.searchUsername}`,
				null,
				{
        			headers: {Authorization: `Bearer ${this.userId}`}
      			});
				if (response.status!=200){
					throw(response.status)
				}

				this.isfollowing = true
				await this.getProfile()

			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		async unfollow() {
			this.loading = true;
			this.errormsg = null;
			try {
				
				let response = await this.$axios.delete(`/users/${this.username}/followings/${this.searchUsername}`,{
        			headers: {Authorization: `Bearer ${this.userId}`}
      			});
				
				if (response.status!=204){
					throw(response.status)
				}

				this.isfollowing = false
				await this.getProfile()

			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		async handlebtn2(){
			
			if(this.searchbanned===true){
	
				await this.unban()
				return
			}
			if(this.searchbanned===false){
			
				await this.ban()
				return 
			}
		},
		async ban() {
			this.loading = true;
			this.errormsg = null;
			try {
				
				let response = await this.$axios.put(`/users/${this.username}/bans/${this.searchUsername}`,
				null,
				{
        			headers: {Authorization: `Bearer ${this.userId}`}
      			});
				
				if (response.status!=200){
					throw(response.status)
				}

				this.searchbanned = true
				

			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},

		async unban() {
			this.loading = true;
			this.errormsg = null;
			try {
				
				let response = await this.$axios.delete(`/users/${this.username}/bans/${this.searchUsername}`,{
        			headers: {Authorization: `Bearer ${this.userId}`}
      			});
				
				if (response.status!=204){
					throw(response.status)
				}

				this.searchbanned = false
				

			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
	},
	
	async created (){
		this.userId = localStorage.getItem('userId')
		if (!this.userId){
			this.$router.push({name: 'Login'})
		}
		else{
			this.searchUsername=this.$route.params.username
		
			if (this.searchUsername!==this.username){
				await this.checkBanned()
				if (this.isbanned===true){
					return
				}
				await this.getBans()
			}		
			
			await this.getProfile()
			await this.getPhotos();
		}
		
	}

}

</script>

<template>
	<div id="profile" class="container">
		<Navbar :username="username"></Navbar>
		<Mystats v-if="!isbanned && errormsg===null" :username="searchUsername" :numfollowers="numfollowers" :numfollowings="numfollowings" :numposts="numposts" :userId="userId" @messageToParent=handleMessageFromMystats></Mystats>
		<Grid v-if="photos !== null && !isbanned && errormsg===null" :username="searchUsername" :userId="userId" :photos="photos"></Grid>
		<ErrorMsg id="errorProf" v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<button v-if="searchUsername!==username && !isbanned && errormsg===null" id="followbut" class="btn btn-primary" @click="handlebtn">{{isfollowing ? "unfollow" : "follow"}} </button>
		<button v-if="searchUsername!==username && !isbanned && errormsg===null" id="banbut" class="btn btn-primary" @click="handlebtn2">{{searchbanned ? "remove ban" : "ban"}} </button>
		<div id="banallert" class="alert alert-primary" role="alert" v-if="isbanned">
		You have been banned!
		</div>
	</div>
	<Modal v-if="searchUsername===username && !isbanned && errormsg===null" :username="username" :userId="userId"></Modal>
	<Dropbox v-if="searchUsername===username && !isbanned && errormsg===null" @changesDropbox="handleMessageFromDropbox" :username="username" :userId="userId"></Dropbox>
	
</template>
<style>
#banallert{
left: 500px;
top: 250px;
}
#errorProf{
margin-left: 0px;
left: 500px;
top: 250px;
}

#profile{
	position: relative;
	display: flex;
}

.btn-primary{
	 background: purple;
	 border: none;
	 border-radius: 8px;
 }

.btn-primary:hover {
      background-color: rgb(243, 187, 5);
      transition: 0.7s;
}

#followbut {
	position: absolute;
	left: 250px;
	top: 70px;
 }
#banbut {
	position: absolute;
	left: 350px;
	top: 70px;
}
</style>