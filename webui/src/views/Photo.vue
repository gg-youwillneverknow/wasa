<script>
import Likes from '../components/Likes.vue'
import Comments from '../components/Comments.vue'
  export default {
    name: 'Photo',
    components: {
    Likes,
    Comments
	  },
    data () {
        return{
            userId: localStorage.getItem('userId'),
            username: localStorage.getItem('username'),
            owner: null,
            errormsg: null,
            photo: null,
            imageUrl: null,
            formattedDate: null,
            liked: false,
        }
    },
    methods: {
        navigateToAnotherView(event) {
          if (event.target === event.currentTarget){
            
            this.$router.go(-1)
          }
        },
        async deletePhoto(){
          this.errormsg = null;
          try {
              let response = await this.$axios.delete(this.$route.path,{
        			headers: {Authorization: `Bearer ${this.userId}`}
      			});
              if (response.status !== 204) {
                throw response.status;
              }  
              this.$router.push({name: 'Profile'})         
          } catch (e) {
              this.errormsg = e.toString();
          }

        },
        handleMessageFromLikes(likes) {
          for (const item of likes){
            if(this.username===item.Username){
              this.liked=true
            }
          }
          
        },
        async handleMessageFromComments(){
          await this.getPhoto()
        },
        async toggleProps () {
          
          if(this.liked===false){
	
            await this.likePhoto()
            return
          }
          if(this.liked===true){

            await this.unlikePhoto()
            return 
          }
          
        },
        async likePhoto() {
          this.errormsg = null;
            try {
             
                let response = await this.$axios.put(this.$route.path+`/likes/${this.userId}`,
                null,
                {
        			  headers: {Authorization: `Bearer ${this.userId}`}
      			    });
                if (response.status != 200) {
                throw response.status;
                }
                this.liked=true
                await this.getPhoto()
            } catch (e) {
                this.errormsg = e.toString();
            }
        },
        async unlikePhoto() {
          this.errormsg = null;
          try {

            let response = await this.$axios.delete(this.$route.path+`/likes/${this.userId}`,{
        			headers: {Authorization: `Bearer ${this.userId}`}
      			});
            if (response.status!=204){
              throw(response.status)
            }
            this.liked = false
            await this.getPhoto()
          } catch (e) {
            this.errormsg = e.toString();
          }
        },
        async getPhoto() {

            this.errormsg = null;
            try {
                let response = await this.$axios.get(this.$route.path,{
                headers: {Authorization: `Bearer ${this.userId}`}
                });
                if (response.status !== 200) {
                throw response.status;
                }
                this.photo = JSON.parse(JSON.stringify(response.data));
                
            } catch (e) {
                this.errormsg = e.toString();
            }

        },
        async getImage(photoId) {
            try {
                let response = await this.$axios.get(`/images/${photoId}`, {
                responseType: "blob",
                headers: { Accept: "image/jpeg", Authorization: `Bearer ${this.userId}`}
                });
                if (response.status !== 200) {
                throw response.status;
                }
                const imageUrl = URL.createObjectURL(response.data);

                return imageUrl;
            } catch (e) {
                this.errormsg = e.toString();
                
            }
        },
        async resolveImageUrl() {
            if (!this.photo) return;    
    
            this.imageUrl = await this.getImage(this.photo.ID);
                
        }, 

    },

    watch: {
          photo(newValue) {
            const date = new Date(newValue.dateTime);

            const day = date.getDate(); 
            const month = date.toLocaleString('default', { month: 'long' });
            const year = date.getFullYear();
            const minutes = date.getMinutes();
            const hours = date.getHours();
            this.formattedDate = `${hours}:${minutes} ${day} ${month} ${year}`

            this.owner = newValue.owner
          },
        
    },

    async mounted() {
        this.userId = localStorage.getItem('userId')
        if (!this.userId){
          this.$router.push({name: 'Login'})
        }
        else{
          await this.getPhoto();
          await this.resolveImageUrl();
          this.username = localStorage.getItem('username')
        }
    },
};
</script>

<template>
    
  <div class="lightbox" @click="navigateToAnotherView"> 

    <img :src="imageUrl">

    <div class="lightbox-info">
      <div class="lightbox-info-inner">
        <div class="row">
          <div class="col-sm-3">
            <button class="likebtn">
              <span> <font-awesome-icon :icon="['fas', 'user']" size='lg' style="background:transparent"/> 
              </span>	        
            </button>   
          </div>
          <div class="col-sm-9">
            <p v-if="owner">{{ owner }}</p>
          </div>   
        </div>
        <div class="row">
          <div class="col-sm-3">
            <button class="likebtn">
              <span> <font-awesome-icon :icon="['fas', 'calendar-days']" size='lg' style="background:transparent"/> 
              </span>	        
            </button>   
          </div>
          <div class="col-sm-9">
            <p v-if="formattedDate">{{ formattedDate }}</p>
          </div>   
        </div>
        <div class="row">
          <div class="col-sm-3">
            <button @click="toggleProps" class="likebtn"> 
              <span v-if="!liked"> <font-awesome-icon :icon="['far', 'heart']" size='lg'/>
              </span>
              <span v-if="liked"> <font-awesome-icon :icon="['fas', 'heart']" size='lg' style="color: #de0d0d;"/>
              </span>
            </button>
          </div>
          <div class="col-sm-9">
            <button v-if="photo" id="likesbut" data-bs-toggle="modal"  data-bs-target="#likes" class="likebtn"> {{ photo.likesnum }} likes </button>
          </div>   
        </div>
        <div class="row">
          <div class="col-sm-3">
            <span> <font-awesome-icon :icon="['far', 'comment']" size='lg' style="background:transparent"/> </span>	
          </div>
          <div class="col-sm-9">
            <button v-if="photo" id="commentsbut" data-bs-toggle="modal"  data-bs-target="#comments" class="likebtn"> {{ photo.commentsnum }} comments </button>
          </div>   
        </div>
        <div class="row" v-if="username===owner">
          <div class="col-sm-12">
            <button type="button" class="btn btn-danger" @click="deletePhoto">delete</button>
          </div>
        </div>
      </div>
    </div>
  </div>
  <Likes @messageToParent="handleMessageFromLikes" :userId="userId" :liked="liked"></Likes>
  <Comments @changesComments="handleMessageFromComments" :username = "username" :userId="userId"></Comments>
</template>

<style>
.row {
  display: flex;
  align-items: center;
}
.col-sm-3 {
    display: flex;
    align-items: center;
}
.col-sm-9 {
    display: flex;
    align-items: center;
}
.lightbox-info-inner .row .col-sm-9 p {
 
    margin:0;
}
.likebtn {
  background: none;
  border: none;
  color: black;
  cursor: pointer;
  font: inherit;
  outline: none;
  padding: 0;
  height: 37px;
}

.lightbox {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.8);
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  grid-gap: 2rem;
  z-index: 1;
}

.lightbox img {
  margin: auto;
  width: 100%;
  grid-column-start: 2;
}

.lightbox-info {
  margin: auto 2rem auto 0;
}

.lightbox-info-inner {
  background-color: #FFFFFF;
  display: inline-block;
  padding: 2rem;
}

.btn-danger{
  border-radius: 8px;
  }
</style>