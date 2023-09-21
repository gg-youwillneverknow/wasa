<script>
import ErrorMsg from '../components/ErrorMsg.vue';
export default {
	components: {
		ErrorMsg
	},
	props: ["userId","username"],
	name: "Modal",
	data (){
		return{
			errormsg: null,
			newUsername: null,
			path: "/accounts/"+this.userId+"/edit",
			profilePath: "/users/"+this.username+"/profile",
		}
	},
	watch: {
		newUsername(newValue){		
			if (newValue===""){
				this.errormsg=null
			}
	
		}
	},
	methods: {
		async changeUsername(){
			const config = {headers: 
					{ 	'content-type': 'application/json', 
						Authorization: `Bearer ${this.userId}`, 
						token: this.userId }
							}
			this.errormsg = null;
			try {
				let response = await this.$axios.put(this.path,
				JSON.stringify({Username: this.newUsername, ID: this.userId}),config);
				if (response.status!=200){
					throw(response.status)
				}else{

					localStorage.setItem("username",this.newUsername)
					this.profilePath="/users/"+this.newUsername+"/profile"
					this.$router.push({path: this.profilePath})
				}

			} catch (e) {
				this.errormsg = e.toString();
			}
		},
		handleSubmit() {
			// validate username
			this.errormsg = this.newUsername.length > 4 ? '' : 'Username must have at least 5 chars'
			if (this.errormsg!=''){
				
				return}
			this.errormsg =  this.newUsername.length < 26 ? '' : 'Username must have at most 25 chars'
			if (this.errormsg!=''){
				
				return}
			this.errormsg = new RegExp("^[a-zA-Z]").test(this.newUsername) ? '' : 
			'Username must start with alphabetic character'
			if (this.errormsg!=''){
			
				return}
			this.errormsg = new RegExp("^[a-zA-Z][a-zA-Z0-9_.]{4,24}$").test(this.newUsername) ? '' : 
			'Username can only contain alphabetic characters, numbers, underscore and fullstop.'
			if (this.errormsg!=''){
			
				return}
			this.changeUsername()
		}
	}



}

</script>


<template> 
<button id="editbut" data-bs-toggle="modal"  data-bs-target="#edit" class="btn btn-primary"> edit profile </button>
<div class="modal fade" id="edit" tabindex="-1">
  <div class="modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title">Edit profile</h5>
        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
      </div>
      <div class="modal-body">
        <input v-model="newUsername" type="text" id="username" name="username" class="form-control" placeholder="insert new username">
      </div>
      <div class="modal-footer">
        <button v-if="!errormsg" @click="handleSubmit" type="button" class="btn btn-primary">Save changes</button>
        <button v-if="!errormsg" type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
		<ErrorMsg id="error" v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
    </div>
  </div>
</div>

</template>

<style>
#editbut {
     position: absolute;
	 top: 20px;
	 right: 140px;
 }
.form-control {
display: block;
width: 100%;
height: 40px;
padding: 0.375rem 0.75rem;
font-size: 1.2rem;
line-height: 1.6;
background-color: transparent;
background-clip: padding-box;
border: 1px solid purple;
border-radius: 8px;
transition: border-color 0.15s ease-in-out, box-shadow 0.15s ease-in-out;
}
</style>