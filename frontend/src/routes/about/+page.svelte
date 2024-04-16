<script>
    export let data;
    import {page} from '$app/stores';
    
        
</script>
  
<svelte:head>
    <title>About</title>
</svelte:head>
  
  <h1>About</h1>
  <p>host: {$page.url.host}</p>
  <p>protocol: { $page.url.protocol}</p>
  <p>pathname: { $page.url.pathname }</p>
  <p>href: { $page.url.href }</p>
  <p>
    <strong>server respond</strong>:
    {#await data}
      loading
    {:then data}
      {data.Message}
    {:catch err}
      failed to load data
    {/await}
  </p>
  
  <p>This is about page</p>
  <a href="/">Back to home</a>
  
  <style>
    h1 {
      color: green;
    }
  </style>