class Editor extends React.Component {
    save() {
        // Get our App's Editor component
        let editor = document.getElementById('editor')

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

        // Pixels are generated left to right, but physically rendered right-to-left.
        bottomPixels.reverse();

        // Pixels are generated top to bottom, but physically rendered bottom-to-top.
        leftPixels.reverse();

        // Actually represents the side render-ordering of the strip
        let sides = [bottomPixels, leftPixels, topPixels, rightPixels];

        // Will represent the pixel render-ordering of the strip
        let strip = [];

        for (let side of sides) {
            for (let pixel of side) {
                strip.push(pixel.style.backgroundColor);
            }
        }

        console.log(strip);
    }

    render() {
        return (
            <div id="editor">
                <Edge side="top" pixels="39" />
                <Edge side="left" pixels="33" />
                <Edge side="right" pixels="33" />
                <Edge side="bottom" pixels="39" />

                <input type="button" value="Save" onClick={this.save} />
            </div>
        );
    }
}
