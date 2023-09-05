export const POST = async ({ request, redirect }) => {
    // get form data
    const formdata = await request.formData();
    const title = formdata.get("title");

    // Check if the title is missing
    if (!title) {
        return new Response("Missing title", {
            status: 400,
        });
    }

    // Send the data to the API
    const response = await fetch("http://localhost:1323/data/delete", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            title: title,
        }),
    })

    // Check the response
    const responseJson = await response.json();
    if (responseJson.error) {
        return new Response("title not exists", {
            status: 400,
        });
    }
    return redirect("/blog");
}