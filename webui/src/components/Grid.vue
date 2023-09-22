<template>
    <div class="gallery">
      <div class="gallery-panel" v-for="photo in photos" :key="photo.ID">
        <router-link :to="`/users/${photo.owner}/photos/${photo.ID}`">
          <img :src="resolvedImageUrls[photo.ID]" >
        </router-link>  
      </div>
    </div>
  </template>  
<script>
export default {
  name: "Grid",
  props: ["username","photos","userId"],
  data() {
    return {
      errormsg: null,
      loading: null,
      resolvedImageUrls: {}, // New data property to store resolved image URLs
      
    };
  },
  watch: {
    async photos (newPhotos) {
      await this.resolveImageUrls();
    }    
  },
  methods: {
    async getImage(photoId) {
      try {
        let response = await this.$axios.get(`images/${photoId}`, {
          responseType: "blob",
          headers: { Accept: "image/jpeg", 
                    Authorization: `Bearer ${this.userId}`
                  },
        });
        if (response.status !== 200) {
          throw response.status;
        }
        const imageUrl = URL.createObjectURL(response.data);

        return imageUrl;
      } catch (e) {
        console.error(`Error loading image for photo ID ${photoId}:`, e);
        return ""; // Return an empty string or any fallback URL/error handling
      }
    },
    async resolveImageUrls() {

        const promises = this.photos.map(async (photo) => {
            const imageUrl = await this.getImage(photo.ID);
            this.resolvedImageUrls[photo.ID] = imageUrl;
        });
        await Promise.all(promises);

    },
  },
  async created() {
    await this.resolveImageUrls();
  }
};
</script>

<style>
.gallery {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(20rem, 1fr));
    grid-gap: 1rem;
    max-width: 80rem;
    margin: 5rem auto;
    padding: 0 5rem;
    margin-left: 90px;
    margin-top: 130px;
  }

.gallery-panel img {

  width: 100%;
  height: 22vw;
  object-fit: cover;
  border-radius: 0.75rem;
}

</style>