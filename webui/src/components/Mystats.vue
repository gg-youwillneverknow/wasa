<template>
<div class="container" id="table">
    <div class="row"> 
        <div class="col-sm-12">
            <p>{{ username }}</p>
        </div>
    </div>
    <div class="row">
        <div class="col-sm-4">
            <div class="dropdown">
                <button @click="show=!show" class="dropbtn" :class="{'custom-style': show}">Followers</button>
                <div v-show="show" id="myDropdown" class="dropdown-content">
                    <div v-for="follower in followers" :key="follower.Username">
                        <a href="#">{{follower.Username}}</a>
                    </div>
                </div>
            </div>
        </div>
        <div class="col-sm-4">
            <div class="dropdown">
                    <button @click="show2=!show2" class="dropbtn" :class="{'custom-style': show2}">Following</button>
                    <div v-show="show2" id="myDropdown" class="dropdown-content">
                        <div v-for="following in followings" :key="following.Username">
                            <router-link :to="`/users/${following.Username}/profile`">
                                {{following.Username}}
                            </router-link>  
                        </div>
                    </div>
            </div>
        </div>   
        <div class="col-sm-4">Posts</div>
    </div>
    <div class="row">
        <div class="col-sm-4">{{ numfollowers }}</div>
        <div class="col-sm-4">{{ numfollowings }}</div>
        <div class="col-sm-4">{{ numposts }}</div>
    </div>
</div>
</template>
<script>
export default {
    name: "Mystats",
    props: ["username","numfollowers","numfollowings","numposts", "userId"],
    data () {
        return{
            followers: null, 
            followings: null,
            show: false,
            show2: false,
        }
    },
    watch: {
        async username (newUsername) {
            await this.getFollowings()
            await this.getFollowers();
            if (this.followers!=null){
                this.$emit("messageToParent",this.followers);
            }
        },
        numfollowers(newPropValue, oldPropValue) {
            
            this.getFollowers(); // Call your method here
        },  
        numfollowings(newPropValue, oldPropValue) {
            
            this.getFollowings(); // Call your method here
        }  
    },
    methods: {
        async getFollowings(){

            this.loading = true;
			this.errormsg = null;
			try {
				
				let response = await this.$axios.get(`/users/${this.username}/followings/`,{
                headers: {Authorization: `Bearer ${this.userId}`}
                });
				
				if (response.status!=200){
					throw(response.status)
				}

				this.followings = JSON.parse(JSON.stringify(response.data));
				

			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
        },
        async getFollowers(){
            this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get(`/users/${this.username}/followers/`,{
                headers: {Authorization: `Bearer ${this.userId}`}
                });

                if (response.status!=200){
					throw(response.status)
				}

				this.followers = JSON.parse(JSON.stringify(response.data));
                
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
        }
    },
    async created () {
        await this.getFollowings()
        await this.getFollowers();
        if (this.followers!=null){
            this.$emit("messageToParent",this.followers);
        }
    }
}

</script>
<style scoped>
p {
    font-size: 25px;

}
#table{
	position: absolute;
	top: 20px; 
    max-width: 250px;
    left: 790px
}

.dropbtn {
  background: none;
  border: none;
  color: black;
  cursor: pointer;
  font: inherit;
  outline: none;
  padding: 0;
}

.dropbtn:hover, .custom-style {
  color: purple;
  font-weight: bold;
}

/* The container <div> - needed to position the dropdown content */
.dropdown {
  position: relative;
  display: inline-block;
}

/* Dropdown Content (Hidden by Default) */
.dropdown-content {
  position: absolute;
  background-color: rgb(200, 131, 211);
  min-width: 160px;
  border-radius: 8px;
  z-index: 1;
}

/* Links inside the dropdown */
.dropdown-content a {
  color: purple;
  padding: 12px 16px;
  text-decoration: none;
  display: block;
  font-size: 15px;
}

/* Change color of dropdown links on hover */
.dropdown-content a:hover {
    background-color: yellow;
    color: purple;
    border-radius: 8px;
}

</style>