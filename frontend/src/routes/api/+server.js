import { json } from "@sveltejs/kit"

export function GET(){
    const message = {
        Message: "hello from svelte"
    }
    

    return json(message)
}