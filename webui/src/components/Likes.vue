<script>
export default {
	name: "Likes",
	data (){
		return{
			errormsg: null,
			likes: null,
		}
	},
	methods: {
		async getLikes(){
			this.errormsg = null;
			try {
				let response = await this.$axios.get(this.$route.path+`/likes`,{
                headers: {Authorization: "Bearer ${token}",token: localStorage.getItem("userId")}
                });
				if (response.status!=200){
					throw(response.status)
				}
				this.likes=JSON.parse(JSON.stringify(response.data));
				
			} catch (e) {
				this.errormsg = e.toString();
			}
		}
	},
	async mounted () {
		await this.getLikes()
		if (this.likes!=null){
            this.$emit("messageToParent",this.likes);
        }
	}
}

</script>


<template> 
<div class="modal fade" id="likes" tabindex="-1">
  <div class="modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title">Likes</h5>
        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
      </div>
      <div v-if="likes" class="modal-body">
		<ul class="list-group overflow-auto shadow" style="max-height: 100px;">
			<li class="list-group-item" v-for="like in likes" :key="like.Username">{{ like.Username }}</li>
		</ul>
		
      </div>
    </div>
  </div>
</div>
</template>

<style>
</style>