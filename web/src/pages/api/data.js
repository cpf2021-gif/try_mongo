export const getPost = async ( id ) => {
    const response = await fetch(`http://localhost:8000/blogs/get/id/${id}`)
    const responseData = await response.json()
    
    return new Response(JSON.stringify(responseData)) 
}

export const getPosts = async () => {
    const response = await fetch("http://localhost:8000/blogs/list")
    const allPosts = await response.json()
    return new Response(JSON.stringify(allPosts)) 
}