<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			username : '',
			userId: ''
		}
	},
	methods: {
		async refresh(){
			this.loading=true 
			this.errormsg=null
			try {
				let response = await this.$axios.get("/context");
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading=false
		},
		async login() {
			this.loading = true;
			this.errormsg = null;
			try {
				
				let response = await this.$axios.post("/session", JSON.stringify(this.username));
				this.userId  = response.data;
				
				if (response.status==201 || response.status==200){
					localStorage.setItem("userId",response.data)
					localStorage.setItem("username",this.username)
					this.$router.push({name: 'Stream'})
				} else {
					throw(response.status)
				}
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},

		handleSubmit() {
			// validate username
			this.errormsg = this.username.length > 4 ? '' : 'Username must have at least 5 chars'
			if (this.errormsg!=''){
				
				return}
			this.errormsg =  this.username.length < 26 ? '' : 'Username must have at most 25 chars'
			if (this.errormsg!=''){
				
				return}
			this.errormsg = new RegExp("^[a-zA-Z]").test(this.username) ? '' : 
			'Username must start with alphabetic character'
			if (this.errormsg!=''){
			
				return}
			this.errormsg = new RegExp("^[a-zA-Z][a-zA-Z0-9_.]{4,24}$").test(this.username) ? '' : 
			'Username can only contain . and _'
			if (this.errormsg!=''){
			
				return}
			this.login()
		},	

	},
	mounted() {
		let user = localStorage.getItem('userId')
		if (user){
			this.$router.push({name: 'Stream'})
		}
	}
}
</script>

<template>

<div class="container">
<div class="row justify-content-center align-items-center" style="height:100vh">
<div class="col-4">
<div class="card">
<article class="card-body">
	<h4 class="card-title text-center mb-4 mt-1">WasaPhoto Sign in</h4>
	
	<p class="text-success text">The username must begin with a an alphabetic character. It may also contain numbers, underscore and fullstop. It must have between 5 and 25 characters.</p>
	<form @submit.prevent = "handleSubmit">
	<div class="form-group">
	<div class="input-group">
		<div class="input-group-prepend">
		    <span class="input-group-text"> <font-awesome-icon :icon="['fas', 'user']" /> </span>	
		</div> 
		<input v-model="username" aria-describedby="format" type="text" id="username" name="username" class="form-control"  placeholder="username">
	</div> <!-- input-group.// -->
	</div> <!-- form-group// -->
	<div class="form-group">
	<button type="submit" class="btn btn-primary btn-block mt-4"> Login  </button>
	</div> <!-- form-group// -->
	</form>
</article>
</div> 
</div>
</div>
</div>
	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

</template>

<style>
.card {
	border: 1px solid #ffffff;
}
.card-login {
	margin-top: 130px;
	padding: 18px;
	max-width: 30rem;
}

.input-group-prepend span{
	width: 50px;
	height: 40px;
	background-color: purple;
	border: none;
	border-radius: 8px 0px 0px 8px;
	color: #fff;

}

input:focus{
	outline: 0 0 0 0  !important;
	box-shadow: 0 0 0 0 !important;
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
	border-radius: 0px 8px 8px 0px;
	transition: border-color 0.15s ease-in-out, box-shadow 0.15s ease-in-out;
}

.input-group-text {
	display: -ms-flexbox;
	display: flex;
	-ms-flex-align: center;
	align-items: center;
	font-size: 1.5rem;


}

</style>
