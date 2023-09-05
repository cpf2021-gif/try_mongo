export const PUT = async ({ request, redirect }) => {
    // get form data
    const formdata = await request.formData();
    const title = formdata.get("title");
    const content = formdata.get("content");

    // Check if the title or content is missing
    if (!title || !content) {
        return new Response("Missing title or content", {
            status: 400,
        });
    }

    // Send the data to the API
    const response = await fetch("http://localhost:1323/data/update", {
        method: "PUT",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            title: title,
            content: content,
        }),
    })

    // Check the response
    const responseJson = await response.json();

    if (responseJson.error) {
        return new Response("title already exists", {
            status: 400,
        });
    } else {
        return redirect("/blog");
    }
}