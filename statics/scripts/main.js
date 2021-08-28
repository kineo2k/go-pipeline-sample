window.onload = () => {
    document.getElementById("option-input-url").onkeyup = (ev) => {
        if (ev.key !== "Enter")
            return;

        ev.target.blur();

        run();
    };

    document.getElementById("btn-run").onclick = (ev) => {
        ev.target.blur();

        run();
    };
}

const run = () => {
    const spec = {};

    const url = document.getElementById("option-input-url").value;
    if (!url) {
        return;
    }
    spec.input = { url };

    const width = parseInt(document.getElementById("option-resize-width").value, 10);
    const height = parseInt(document.getElementById("option-resize-height").value, 10);
    const keepAspectRatio = document.querySelector("#option-resize-keep:checked") !== null;
    if (!width && !height) {
        return;
    }
    spec.resize = { width, height, keepAspectRatio };

    const anchor = document.querySelector("input[name='option-crop']:checked").value;
    spec.crop = { anchor };

    const type = document.querySelector("input[name='option-effect']:checked").value;
    spec.effect = { type };

    const id = renderCartItem(spec);
    fetchImage(spec, id);
};

const renderCartItem = spec => {
    const url = spec.input.url;
    const id = randomID();

    const cartItem = document.querySelector("#card-item");
    const clone = document.importNode(cartItem.content, true);
    clone.querySelector("div").setAttribute("id", id);

    const cardText = clone.querySelector("p.card-text:nth-child(1) > code");
    cardText.innerHTML = `<pre>${JSON.stringify(spec, undefined, 2)}</pre>`;

    // Input
    clone.querySelector("button:nth-child(1)").onclick = (ev) => {
        const newWin = window.open(url, "_blank");
        newWin.focus();
    };

    // Delete
    clone.querySelector("button:nth-child(3)").onclick = (ev) => {
        ev.target.closest("div[class=col]").remove();
    };

    const cardTable = document.querySelector("#card-table");
    cardTable.appendChild(clone);

    return id;
};

const fetchImage = (spec, id) => {
    axios.post("/image-processing", spec)
        .then(res => {
            const thumbnail = document.querySelector(`#${id} img:nth-child(1)`);
            thumbnail.setAttribute("src", `data:image/jpeg;base64,${res.data}`);

            // Output
            document.querySelector(`#${id} button:nth-child(2)`).onclick = (ev) => {
                const image = new Image();
                image.src = thumbnail.getAttribute("src");

                const newWin = window.open("", "_blank");
                newWin.document.write(image.outerHTML);
                newWin.document.close();
            };
        })
        .catch(function (error) {
            const cardText = document.querySelector(`#${id} p.card-text:nth-child(1) > code`);
            cardText.innerHTML = `<span>Image Processing Failed.</span><pre>${JSON.stringify(spec, undefined, 2)}</pre>`;
        });
};

const randomID = () => "img-" + Math.random().toString().substr(2,11);
