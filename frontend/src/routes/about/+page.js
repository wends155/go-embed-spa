import { error } from "@sveltejs/kit"


export const ssr = false;

/** @type {import('./$types').PageLoad} */  
export async function load({ fetch, url }){
    
    
    
    const res = await fetch("/api");
    const message = await res.json();
    
    if (!res){
        throw error(404);
    }
   
    console.log(url.href);
    return message;

}