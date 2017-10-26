class Edge extends React.Component {
    pixels() {
        var pixels = [];
        var pixel;
        for (pixel = 0; pixel < this.props.pixels; pixel++) {
            pixels.push(<Pixel />)
        }

        return (
            pixels
        )
    }

    render() {
        return (
            <div class="edge" id={ `frame-${this.props.side}` }>
                { this.pixels() }
            </div>
        )
    }
}