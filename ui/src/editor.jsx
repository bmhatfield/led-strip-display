class Editor extends React.Component {
    save() {
        /*
            The actual physical LED strip is a 1px by 150px display. For the purposes
            of this editor, it will be wrapped around a window frame, creating a 39x33
            marquee that encircles the window.

            The service side accepts pixel values in strip order, which does not exactly
            match the ordering of the HTML elements on the page. This means we need to
            reorder slightly so when we send the values to the service, we are sending
            them in clockwise order, starting from the bottom right (as viewed from *outside*
            the window)

            To accomplish this, we'll find the elements matching each side separately, which
            will allow us to order the 'sides' in clockwise order. It will also allow us to
            order the pixels within a given side as appropriate so that the total flow is
            clockwise.
        */

        // Get our App's Editor component
        let editor = document.getElementById('editor');

        // This gets us the HTMLCollection objects representing our DOM nodes
        let bottomPixelsCollection = editor.getElementsByClassName('pixel-bottom');
        let leftPixelsCollection = editor.getElementsByClassName('pixel-left');
        let topPixelsCollection = editor.getElementsByClassName('pixel-top');
        let rightPixelsCollection = editor.getElementsByClassName('pixel-right');

        // But we need to mutate it with array syntax, so let's convert to arrays
        let bottomPixels = Array.from(bottomPixelsCollection);
        let leftPixels = Array.from(leftPixelsCollection);
        let topPixels = Array.from(topPixelsCollection);
        let rightPixels = Array.from(rightPixelsCollection);

        /*
            The HTML elements for each 'side' are generated left-to-right.
            This means that two of the sides will be 'wrong' when considered
            for the purposes of clockwise order, so we need to reverse their
            ordering so the final array of pixels represents the continuous
            clockwise ordering of pixels. The result will be that the display
            matches the UI.
        */
        bottomPixels.reverse();
        leftPixels.reverse();

        // We can now represent the total ordering of HTML elements in a nested array
        let sides = [bottomPixels, leftPixels, topPixels, rightPixels];

        // And this will be the actual array ordering of RGB values to send to the
        // backend service.
        let strip = [];

        // Iterate over the two-level array, extracting RGB values for each pixel.
        for (let side of sides) {
            for (let pixel of side) {
                strip.push(pixel.style.backgroundColor);
            }
        }

        let data = {
            brightness: 100,
            pixels: strip,
        };

        fetch('/frame/rgb', {
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json',
            },
            method: 'POST',
            body: JSON.stringify(data),
        });
    }

    render() {
        return (
            <div id="editor">
                <Edge side="top" pixels="37" bgcolor="#000000" />
                <Edge side="left" pixels="32" bgcolor="#000000" />
                <Edge side="right" pixels="32" bgcolor="#000000" />
                <Edge side="bottom" pixels="37" bgcolor="#000000" />

                <input type="button" value="Save" onClick={this.save} />
            </div>
        );
    }
}
