function getBaseUrl()
{
    let currentUrl = window.location;
    return currentUrl.protocol + "//" + currentUrl.host;
}

async function GET(url)
{
    console.log("GET request: " + url);
    let response = await fetch(url);

    console.log(response.statusText);
    if (response.ok)
    {
        let responseText = await response.text();
        let out = document.getElementById("out");
        out.innerText += responseText + "\n";
        console.log("response body: " + responseText);
    }
    else
    {
        alert("Failed http request. Response status: " + response.status);
    }
}

function onPut()
{
    let key = document.getElementById("put_key_field").value;
    let value = document.getElementById("put_value_field").value;
    let request = getBaseUrl() + "/put?" + key + "=" + value;

    if (key.length !== 0 && value.length !== 0) {
        GET(request).then(() => {
            document.getElementById("put_key_field").value = "";
            document.getElementById("put_value_field").value = "";
        });
    }
}

function onRemove()
{
    let key = document.getElementById("remove_key_field").value;
    let request = getBaseUrl() + "/remove?key=" + key;

    if (key.length !== 0) {
        GET(request).then(() => {
            document.getElementById("remove_key_field").value = "";
        });
    }
}

function onGet()
{
    let key = document.getElementById("get_key_field").value;
    let request = getBaseUrl() + "/get?key=" + key;

    if (key.length !== 0) {
        GET(request).then(() => {
            document.getElementById("get_key_field").value = "";
        });
    }
}