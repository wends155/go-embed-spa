<script>
	import { page } from '$app/stores';
	
</script>

<svelte:head>
    <title>Page not Found</title>
</svelte:head>
  
<h1>{$page.status} {$page.error.message}</h1>
<p>Error </p>
  
 
 