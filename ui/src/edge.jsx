class Edge extends React.Component {
    pixels() {
        let pixels = [];
        let pixel;
        for (pixel = 0; pixel < this.props.pixels; pixel++) {
            pixels.push(<Pixel id={ pixel } side={ this.props.side } />)
        }

        return (
            pixels
        );
    }

    render() {
        return (
            <div class="edge" id={ `frame-${this.props.side}` }>
                { this.pixels() }
            </div>
        );
    }
}
