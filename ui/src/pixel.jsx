class Pixel extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            bgcolor: '#000000',
        };
    }

    updateColor(event) {
        console.log(event);
        this.setState({
            bgcolor: event.target.value,
        });
    }

    render() {
        return (
            <label className={`pixel pixel-${this.props.side}`}
                 id={ `pixel-${this.props.side}-${this.props.id + 1}` }
                 style={ {backgroundColor: this.state.bgcolor} }>
                 <input type="color" onChange={ this.updateColor.bind(this) } />
            </label>
        );
    }
}
