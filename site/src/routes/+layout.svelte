<script lang="ts">
  import '../app.css';

  import Navbar from '$lib/navbar/Navbar.svelte';
  import Sidebar from '$lib/navbar/Sidebar.svelte';
  import Footer from '$lib/navbar/Footer.svelte';

  let open = false;
</script>

<Sidebar bind:open />
<Navbar bind:sidebar={open} />

<div class="main">

<slot />

<Footer />

</div>

<style lang="postcss">
  :global(body){
    padding: 0;
    background-color: rgb(223, 224, 184);
  }
</style>
