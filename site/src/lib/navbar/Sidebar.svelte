<script>
	import '../css/navbar.scss';
	export let open = false;
</script>

<aside class="absolute w-full h-full nav-color border-r-2 shadow-lg" class:open>
	<nav class="p-12 text-xl sidebar-box">
		<a href="/" class="block nav-text hover:active-nav">Home</a>
		<a href="/about" class="block nav-text hover:active-nav">About</a>
		<a href="/blog" class="block nav-text hover:active-nav">Blog</a>
    <a href="/my-top-picks" class="block nav-text hover:active-nav">My Top Picks</a>
		<a href="/contact" class="block nav-text hover:active-nav">Contact</a>
	</nav>
</aside>

<style lang="scss">
	aside {
		left: -100%;
		transition: left 0.3s ease-in-out;
		z-index: 2;
	}

	.open {
		left: 0;
	}
</style>
