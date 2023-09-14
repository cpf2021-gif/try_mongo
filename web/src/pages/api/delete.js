export const POST = async ({ request, redirect }) => {
    // get form data
    const formdata = await request.formData();
    const id = formdata.get("id");

    // Check if the title is missing
    if (!id) {
        return new Response("Missing title", {
            status: 400,
        });
    }

    // Send the data to the API
    const response = await fetch(`http://localhost:8000/blogs/delete/id/${id}`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
    })

    // Check the response
    const responseJson = await response.json();
    if (!responseJson.ok) {
        return new Response("title not exists", {
            status: 400,
        });
    }
    return redirect("/blog");
}