class Color extends React.Component { // eslint-disable-line no-unused-vars
    style() {
        return {
            backgroundColor: this.props.color,
        };
    }

    render() {
        return (
            <div className="color"
                style={ this.style() }
                onMouseDown={ () => this.props.colorUpdate(this.props.color) }
                onMouseUp={ () => this.props.colorUpdate(null) }>
            </div>
        );
    }
}
