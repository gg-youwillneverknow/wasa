<script>

  const STATUS_INITIAL = 0, STATUS_SAVING = 1, STATUS_SUCCESS = 2, STATUS_FAILED = 3;

  export default {
  emits: ['changesDropbox'],
	name: "Dropbox",
	props: ["userId","username"],
    data() {
      return {
        uploadedFile: null,
        uploadError: null,
        currentStatus: null,
        uploadFieldName: 'photos',
      }
    },
    computed: {
      path(){ return "/users/"+this.username+"/photos/"},
      isInitial() {
        return this.currentStatus === STATUS_INITIAL;
      },
      isSaving() {
        return this.currentStatus === STATUS_SAVING;
      },
      isSuccess() {
        return this.currentStatus === STATUS_SUCCESS;
      },
      isFailed() {
        return this.currentStatus === STATUS_FAILED;
      }
    },
    methods: {
      reset() {

        this.currentStatus = STATUS_INITIAL;
        this.uploadedFile = null;
        this.uploadError = null;
      },
      save(formData) {

        this.currentStatus = STATUS_SAVING;

        this.upload(formData)
          .then(response => {
            this.uploadedFile = JSON.parse(JSON.stringify(response.data));
            this.currentStatus = STATUS_SUCCESS;
            this.$emit("changesDropbox");
          })
          .catch(err => {
            this.uploadError = err.response;
            this.currentStatus = STATUS_FAILED;
          });
      },
      fileChange(event) {
        let fieldName=event.target.name
        let fileList=event.target.files
        const formData = new FormData();
        if (fileList.length!=1) return;
        formData.append(fieldName, fileList[0], fileList[0].name);
        // save it
        this.save(formData);
      },

      upload(formData) {
        const config = {headers: { 'content-type': 'multipart/form-data',
                                  Authorization: `Bearer ${this.userId}`
                                }
                      }
        return this.$axios.post(this.path, formData, config)
          // get data
      }
    },
    mounted() {
      this.reset();
    },
  }

</script>


<template> 
<button id="postbut" data-bs-toggle="modal" @click="reset" data-bs-target="#post" class="btn btn-primary"> create post </button>
<div class="modal fade" id="post" tabindex="-1">
  <div class="modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title">Create post</h5>
        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
      </div>
      <div class="modal-body">
		<div class="container">
		<form enctype="multipart/form-data" novalidate v-if="isInitial || isSaving">
			<h1>Upload image</h1>
			<div class="dropbox">
				<input type="file" :name="uploadFieldName" :disabled="isSaving" @input="fileChange"
				accept="image/*" class="input-file">
				<p v-if="isInitial">
				Drag your file here to begin<br> or click to browse
				</p>
				<p v-if="isSaving">
				Uploading  file...
				</p>
			</div>
    </form>
  	  	</div>
	  </div>
    </div>
  </div>
</div>
</template>
<style>
 #postbut {
     position: absolute;
     left: 320px;
     top: 70px;
 }
</style>

