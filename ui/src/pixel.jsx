class Pixel extends React.Component { // eslint-disable-line no-unused-vars
    constructor(props) {
        super(props);

        this.state = {
            bgcolor: this.props.bgcolor,
            bordercolor: 'white',
        };
    }

    updateColor(event) {
        this.setState({
            bgcolor: event.target.value,
        });
    }

    componentWillReceiveProps(nextProps) {
        this.setState({
            bgcolor: nextProps.bgcolor,
        });
    }

    style() {
        return {
            backgroundColor: this.state.bgcolor,
            borderColor: this.state.bordercolor,
        };
    }

    render() {
        return (
            <label className={`pixel pixel-${this.props.side}`}
                 id={ `pixel-${this.props.side}-${this.props.id + 1}` }
                 style={ this.style() }>
                 <input type="color" onChange={ this.updateColor.bind(this) } value={this.state.bgcolor} />
            </label>
        );
    }
}
