<script>
export default {
	name: "Navbar",
	data: function(){
		return{
			login: false,
			searchUsername: null,
		}
	},
	computed: {
		profilePath(){ return "/users/"+this.username+"/profile"}
	},
	methods: {
		logout() {
			localStorage.clear()
			this.$router.push({name:"Login"})
		},
		async getProfile() {
			this.loading = true;
			this.errormsg = null;
			this.$router.push(`/users/${this.searchUsername}/profile`)			
		},
		
	},
	props: ['username']
}

</script>

<template>
  <div class="sidebar" id="sidenav">
	<header>
		<div class="text header-text">
		<span class="appname">WasaPhoto</span>
	</div>
	</header>
    <div class="sidebar-body" id="sidebarbody">
		<ul class="nav flex-column">
			<li class="">
				<RouterLink to="/" >
					<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
					<span class="text">Home</span>
				</RouterLink>
			</li>
			<li class="">
				<RouterLink :to="profilePath">
					<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#user"/></svg>
					<span class="text">Profile</span>
				</RouterLink>
			</li>
			<li class="search-box">
				<a class="">
					<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#search"/></svg>
					<input @keyup.enter="getProfile" v-model="searchUsername" placeholder="Search user..." type="text" required>
				</a>
			</li>
			<li v-on:click="logout" class="">
				<a class="">
					<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#log-out"/></svg>
					<span class="text">Log out</span>
				</a>
			</li>
		</ul>
    </div>
  </div>
</template>
<style >

  @import url('https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;500;600;700&display=swap');

*{
	font-family: 'Poppins', sans-serif;
}

body{
	height: 100vh;
	background: var(--body-color);
}

:root {
  --primary-color: #695CFE;
  --primary-color-light: #F6F5FF;
  --toggle-color: #DDDD;
  --text-color: #707070;
  --body-color: #E4E9F7;
  --sidebar-color: #fff;

  --tran-02: all 0.2s ease;
  --tran-03: all 0.3s ease;
  --tran-04: all 0.4s ease;
  --tran-05: all 0.5s ease;

}

.sidebar{
	background: var(--sidebar-color);	
	position: fixed;
	top: 0;
	left: 0;
	height: 100%;
	width: 250px;
	padding: 10px 14px;
}

.sidebar .text{
	font-size: 16px;
	font-weight: 500;
	color: var(--text-color);
}

.header-text{
	display: flex;
	flex-direction: column;
	color: var(--text-color);
	margin-left: 16px;
}

.header-text .name{
	font-weight: 600;
}

.sidebar li{
	height: 50px;
	margin-top: 10px;
	list-style: none;
	display: flex;
	align-items: center;
}

.sidebar li svg{
	justify-content: center;
	min-width: 60px;
	font-size: 20px;
	display: flex;
	align-items: center;
}

.sidebar li a {
	height: 100%;
	width: 100%;
	text-decoration: none;
	display: flex;
	align-items: center;
	border-radius: 6px;
	transition: var(--tran-04);
}
.sidebar li svg, .sidebar li .text {
	color: var(--text-color);
	transition: var(--tran-02);
}
.sidebar li a:hover, .sidebar li a:hover input{
	background: var(--primary-color);
}
.sidebar li a:hover svg, .sidebar li a:hover .text{
	color: var(--sidebar-color);
}
.sidebar .search-box{
	background: var(--primary-color-light);
}
.search-box input{
	height: 100%;
	width: 100%;
	outline: none;
	border: none;
	border-radius: 6px;
	background: var(--primary-color-light);
	transition: var(--tran-04);
}
.active-link {
  background: var(--primary-color);
}
.sidebar li .active-link .text, .sidebar li .active-link svg{
  color: var(--sidebar-color);
}
</style>