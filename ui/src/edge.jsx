class Edge extends React.Component {
    pixels() {
        let pixels = [];
        let pixel;
        for (pixel = 0; pixel < this.props.pixels; pixel++) {
            pixels.push(<Pixel id={ pixel } side={ this.props.side } key={ `pixel-${this.props.side}-${pixel}` } />)
        }

        return (
            pixels
        );
    }

    render() {
        return (
            <div className="edge" id={ `frame-${this.props.side}` }>
                { this.pixels() }
            </div>
        );
    }
}
