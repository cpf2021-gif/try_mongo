export const getPost = async ( name ) => {
    const response = await fetch(`http://localhost:1323/data/${name}`)
    const responseData = await response.json()
    
    return new Response(JSON.stringify(responseData)) 
}

export const getPosts = async () => {
    const response = await fetch("http://localhost:1323/data/list")
    const allPosts = await response.json()
    return new Response(JSON.stringify(allPosts)) 
}