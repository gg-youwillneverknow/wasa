<script>
export default {
	name: "Comments",
	props: ["username"],
	data (){
		return{
			errormsg: null,
			comments: [],
		}
	},
	computed: {
		formData () { 
			return {Text: "", Commenter: this.username} 
		}
	},
	methods: {
		async getComments(){
			this.errormsg = null;
			try {
				let response = await this.$axios.get(this.$route.path+`/comments`,{
                headers: {Authorization: "Bearer ${token}",token: localStorage.getItem("userId")}
                });
				if (response.status!=200){
					throw(response.status)
				}
				this.comments=JSON.parse(JSON.stringify(response.data));
				
			} catch (e) {
				this.errormsg = e.toString();
			}
		},
		async postComment(){
			const config = {headers: { 'content-type': 'application/json',
									Authorization: "Bearer ${token}",
									token: localStorage.getItem("userId") }
							}
			this.errormsg = null;
			try {

				let response = await this.$axios.post(this.$route.path+`/comments`,
				JSON.stringify(this.formData),config);
				
				if (response.status!=201){
					throw(response.status)
				}
				await this.getComments()
				this.$emit("changesComments");
				
			} catch (e) {
				this.errormsg = e.toString();
			}
		},

		async deleteComment(commentId) {
			this.errormsg = null;
			try {
				let response = await this.$axios.delete(this.$route.path+`/comments/${commentId}`,{
                headers: {Authorization: "Bearer ${token}",token: localStorage.getItem("userId")}
                });
				if (response.status!=204){
					throw(response.status)
				}
				await this.getComments()
				this.$emit("changesComments");
				
			} catch (e) {
				this.errormsg = e.toString();
			}
		}
	},
	async mounted () {
		await this.getComments()
	}
}

</script>


<template> 
<div class="modal fade" id="comments" tabindex="-1">
  <div class="modal-dialog modal-dialog-scrollable">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title">Comments</h5>
        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
      </div>
      <div class="modal-body">
		<hr>
		<ul class="list-group overflow-auto">
			<li class="list-group-item"  style="max-height: 100px;" v-for="comment in comments" :key="comment.ID">
				<p>{{ comment.commenter }}</p>
				<p>{{ comment.text }}</p>
				<button type="button" v-if="username===comment.commenter" @click="deleteComment(comment.ID)" class="btn btn-outline-primary">Delete</button>
			</li>
		</ul>
		<hr>
		<form @submit.prevent = "postComment" class="form-group">
			<input type="text" name="Text" v-model="formData.Text" placeholder="Write a comment..." class="form-control">
			<button type="submit" class="btn btn-primary"> Publish </button>
    	</form>		
      </div>
    </div>
  </div>
</div>
</template>

<style scoped>
.form-group {
  display: grid;
  grid-template-columns: 1fr auto; /* Divide the container into two columns, the first one will take the remaining space, and the second will be auto-sized (fit the button's width) */
}
hr{
	border-color: gray;
	margin: 0px;
}
.modal-body {
	padding: 0px;
}

.list-group-item {
	border: white;
}

.list-group-item p:first-child {
	display: inline-block;
    margin: 5px;
	font-weight: bold;
}
.list-group-item p {
	display: inline-block;
	margin-right: 20px;
}

.list-group-item button {
	padding-bottom: 0px;
	padding-left: 4px;
	padding-right: 4px;
	padding-top: 0px;
}

.form-control {
	padding: 0.375rem 0.75rem;
	font-size: 1.2rem;
	line-height: 1.6;
	color: white;
	background-color: purple;
	background-clip: padding-box;
	border: 1px solid purple;
	border-radius: 0px;
	transition: border-color 0.15s ease-in-out, box-shadow 0.15s ease-in-out;
}
.form-control::placeholder {
    color: white;
}

.btn-primary{
	border-radius: 0px;
}

.btn-outline-primary{
	color: purple;
	border-color: purple;
	border-radius: 8px;
}

.btn-outline-primary:hover{
	background-color: purple;
	color: white;
}

.modal{
--bs-modal-header-border-color: --bs-white-rgb;
--bs-modal-header-border-width: 0px;
}
</style>