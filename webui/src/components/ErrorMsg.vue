<script>
export default {
	name: "ErrorMsg",
	props: ['msg']
}
</script>

<template>
	<div class="alert alert-danger" role="alert">
		{{ this.msg }}
		
	</div>
</template>

<style scoped>
</style>
