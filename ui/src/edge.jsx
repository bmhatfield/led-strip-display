class Edge extends React.Component { // eslint-disable-line no-unused-vars
    pixels() {
        let pixels = [];
        let pixel;
        for (pixel = 0; pixel < this.props.pixels; pixel++) {
            let bgcolor = this.props.bgcolor;

            if (this.props.pixeldata && this.props.pixeldata[pixel]) {
                bgcolor = this.props.pixeldata[pixel];
            }

            pixels.push(<Pixel id={ pixel } bgcolor={bgcolor} side={ this.props.side } key={ `pixel-${this.props.side}-${pixel}` } />);
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
